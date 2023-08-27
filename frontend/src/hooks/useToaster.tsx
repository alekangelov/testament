"use client";
import { create } from "zustand";

interface Notification {
  id: number;
  message: string;
  title?: string;
  type: "success" | "error" | "warning" | "info";
  height: number;
  state: "entering" | "in" | "exiting";
  duration: number | "infinite";
}

interface Toaster {
  notifications: Record<number, Notification>;
  add: (notification: Partial<Notification>) => VoidFunction;
  remove: (id: number) => void;
  clear: () => void;
  setHeight: (id: number, height: number) => void;
  setState: (id: number, state: Notification["state"]) => void;
}

export const useToaster = create<Toaster>((set) => ({
  notifications: {},
  add: ({ message = "", type = "info", duration = 5000, title }) => {
    if (!message) return () => {};
    const id = Date.now();
    set((state) => ({
      notifications: {
        ...state.notifications,
        [id]: {
          id,
          message,
          type,
          height: 0,
          state: "entering",
          duration,
          title,
        },
      },
    }));
    return () => {
      set((state) => {
        const { [id]: _, ...notifications } = state.notifications;
        return { notifications };
      });
    };
  },
  setState: (id, state) => {
    set((s) => ({
      notifications: {
        ...s.notifications,
        [id]: {
          ...s.notifications[id],
          state,
        },
      },
    }));
  },
  remove: (id) => {
    set((state) => {
      const { [id]: _, ...notifications } = state.notifications;
      return { notifications };
    });
  },
  clear: () => {
    set({ notifications: {} });
  },
  setHeight: (id, height) => {
    set((state) => ({
      notifications: {
        ...state.notifications,
        [id]: {
          ...state.notifications[id],
          height,
        },
      },
    }));
  },
}));
