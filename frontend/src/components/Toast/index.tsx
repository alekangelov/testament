"use client";
import { useToaster } from "@hooks/useToaster";
import { useEffect, useRef } from "react";
import s from "./index.module.css";
import { clsx } from "@data/clsx";
import { Close, Info, Warning, Error, Check } from "@components/Icons";

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
  id,
  message,
  type,
  state,
  duration = 5000,
  title,
}: {
  id: number;
  message: string;
  type: Type;
  state: "entering" | "in" | "exiting";
  duration?: number | "infinite";
  title?: string;
}) => {
  const { remove, setHeight, height, setState } = useToaster((s) => ({
    remove: s.remove,
    setHeight: s.setHeight,
    height: s.notifications[id]?.height,
    setState: s.setState,
  }));
  const ref = useRef<HTMLDivElement>(null);

  // height measurement and animation
  useEffect(() => {
    if (!ref.current) return;
    const height = ref.current?.offsetHeight;
    const out = setTimeout(() => {
      setHeight(id, height);
      setState(id, "in");
    }, 50);
    return () => {
      clearTimeout(out);
    };
  }, [id, setHeight, setState]);

  // exit animation
  useEffect(() => {
    if (duration === "infinite") return;
    TIMEOUTS[id] = setTimeout(() => {
      setState(id, "exiting");
    }, duration);
    return () => {
      clearTimeout(TIMEOUTS[id]);
    };
  }, [id, setState, duration]);

  // remove after animation finishes
  useEffect(() => {
    if (state !== "exiting") return;
    const out = setTimeout(() => {
      remove(id);
    }, 500);
    return () => {
      clearTimeout(out);
    };
  }, [id, remove, state]);
  const Icon = ICONS[type];
  return (
    <div
      className={clsx(s.toast, s[state])}
      style={
        {
          "--maxheight": `calc(${height}px + 1rem)`,
        } as any
      }
    >
      <div className={s.toastInner} ref={ref}>
        <div
          role="presentation"
          className={s.infobox}
          style={
            {
              color: COLORS[type],
              background: "white",
            } as any
          }
        >
          <Icon size={16} />
        </div>
        <div className={s.message}>
          {title && <div className={s.title}>{title}</div>}
          <div className={s.messageInner}>{message}</div>
        </div>
        <button
          className={s.close}
          onClick={(e) => {
            e.stopPropagation();
            setState(id, "exiting");
            setState;
          }}
        >
          <Close />
        </button>
      </div>
    </div>
  );
};

export default function ToastContainer() {
  const notifications = useToaster((s) => s.notifications);
  const arrNotifications = Object.values(notifications);
  return (
    <div className={s.container}>
      {arrNotifications.reverse().map((notification) => (
        <Toast key={`${notification.id}-toast`} {...notification} />
      ))}
    </div>
  );
}
