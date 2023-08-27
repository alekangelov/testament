import { clsx } from "@data/clsx";
import s from "./index.module.css";
import Spin from "@components/Spin";
import { Accessor, JSX } from "solid-js";

const SIZE = {
  small: s.sm,
  medium: s.md,
  large: s.lg,
};

type SpinProps = Parameters<typeof Spin>[0];

export default function Button({
  bordered,
  loading,
  loadingColor,
  ...props
}: JSX.HTMLElementTags["button"] & {
  bordered?: boolean;
  size?: keyof typeof SIZE;
  loading?: Accessor<boolean>;
  loadingColor?: SpinProps["color"];
}) {
  const size = SIZE[props.size || "medium"];
  return (
    <button
      {...props}
      disabled={props.disabled || loading?.()}
      class={clsx(s.wrapper, bordered && s.bordered, size, props.class)}
    >
      {loading?.() && (
        <div class={clsx(s.loader, loading?.() && s.open)}>
          <Spin color={loadingColor} size="small" />
        </div>
      )}
      <span class={clsx(s.children, !loading?.() && s.open)}>
        {props.children}
      </span>
    </button>
  );
}
