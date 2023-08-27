import { Accessor, createSignal } from "solid-js";
import { createStore } from "solid-js/store";

export interface Notification {
  id: number;
  message: string;
  title?: string;
  type: "success" | "error" | "warning" | "info";
  height: number;
  state: "entering" | "in" | "exiting";
  duration: number | "infinite";
}

interface Toaster {
  notifications: Notification[];
  add: (notification: Partial<Notification>) => VoidFunction;
  remove: (id: number) => void;
  clear: () => void;
  setHeight: (id: number, height: number) => void;
  setState: (id: number, state: Notification["state"]) => void;
}

const [toasts, setToasts] = createStore([] as Notification[]);

const selector =
  <T extends { id: number }>(id: number) =>
  (x: T) =>
    x.id === id;

export const useToaster = (): Toaster => {
  const notifications: Toaster["notifications"] = toasts;
  const deleteToast = (id: number) => {
    setToasts((prev) => {
      return prev.filter((x) => x.id !== id);
    });
  };
  const setHeight: Toaster["setHeight"] = (id, height) => {
    setToasts(selector(id), "height", height);
  };
  const clear: Toaster["clear"] = () => {
    const ids = Object.keys(toasts).map((id) => parseInt(id));
    for (const id of ids) {
      remove(id);
    }
  };
  const setState: Toaster["setState"] = (id, state) => {
    setToasts(selector(id), "state", state);
  };
  const remove: Toaster["remove"] = (id) => {
    setState(id, "exiting");
    setTimeout(() => {
      deleteToast(id);
    }, 500);
  };
  const add: Toaster["add"] = (notification) => {
    const id = Date.now();
    const x = {
      id,
      message: notification.message ?? "",
      title: notification.title,
      type: notification.type ?? "info",
      height: 0,
      state: "entering",
      duration: notification.duration ?? 3000,
    } as const;
    setToasts((prev) => {
      return [...prev, x];
    });
    setTimeout(() => {
      setState(id, "in");
    }, 100);
    return () => {
      remove(id);
    };
  };
  return {
    notifications,
    add,
    remove,
    clear,
    setHeight,
    setState,
  };
};
