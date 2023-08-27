import { useToaster } from "@hooks/useToaster";
import s from "./index.module.css";
import { clsx } from "@data/clsx";
import { Close, Info, Warning, Error, Check } from "@components/Icons";
import { Accessor, For, createEffect, createMemo, onMount } from "solid-js";

type Type = "success" | "error" | "warning" | "info";

const ICONS = {
  success: Check,
  error: Error,
  warning: Warning,
  info: Info,
};

const COLORS: Record<Type, string> = {
  success: "rgb(var(--success-rgb))",
  error: "rgb(var(--error-rgb))",
  warning: "rgb(var(--warning-rgb))",
  info: "rgb(var(--info-rgb))",
};

const TIMEOUTS: Record<number, NodeJS.Timeout> = {};

const Toast = ({
  notification,
}: {
  notification: Accessor<{
    id: number;
    message: string;
    type: Type;
    state: "entering" | "in" | "exiting";
    duration?: number | "infinite";
    title?: string;
    height?: number;
  }>;
}) => {
  const { setHeight, remove } = useToaster();

  const Icon = createMemo(() => ICONS[notification().type]);

  onMount(() => {
    const duration = notification().duration;
    if (duration === "infinite") return;
    const timeout = setTimeout(() => {
      remove(notification().id);
    }, duration || 5000);
    TIMEOUTS[notification().id] = timeout;
    return () => {
      clearTimeout(timeout);
    };
  });

  return (
    <div
      data-toast-id={notification().id}
      class={clsx(s.toast, s[notification().state])}
      style={
        {
          "--maxheight": `calc(${notification()?.height}px + 1rem)`,
        } as any
      }
    >
      <div
        class={s.toastInner}
        ref={(x) => {
          setTimeout(() => {
            const { height } = x.getBoundingClientRect();
            setHeight(notification().id, height);
          }, 50);
        }}
      >
        <div
          role="presentation"
          class={s.infobox}
          style={
            {
              color: COLORS[notification().type],
              background: "white",
            } as any
          }
        >
          {Icon()({ size: 24 })}
        </div>
        <div class={s.message}>
          {notification().title && (
            <div class={s.title}>{notification().title}</div>
          )}
          <div class={s.messageInner}>{notification().message}</div>
        </div>
        <button
          class={s.close}
          onClick={(e) => {
            e.stopPropagation();
            remove(notification().id);
          }}
        >
          <Close />
        </button>
      </div>
    </div>
  );
};

export default function ToastContainer() {
  const { notifications } = useToaster();
  return (
    <div class={s.container}>
      <For each={[...notifications].reverse()}>
        {(notification) => {
          return <Toast notification={() => notification} />;
        }}
      </For>
    </div>
  );
}
