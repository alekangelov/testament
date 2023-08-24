import { clsx } from "@data/clsx";
import s from "./index.module.css";
import Spin from "@components/Spin";

const SIZE = {
  small: s.sm,
  medium: s.md,
  large: s.lg,
};

type SpinProps = React.ComponentProps<typeof Spin>;

export default function Button({
  bordered,
  loading,
  loadingColor,
  ...props
}: React.ButtonHTMLAttributes<HTMLButtonElement> & {
  bordered?: boolean;
  size?: keyof typeof SIZE;
  loading?: boolean;
  loadingColor?: SpinProps["color"];
}) {
  const size = SIZE[props.size || "medium"];
  return (
    <button
      {...props}
      disabled={props.disabled || loading}
      className={clsx(s.wrapper, bordered && s.bordered, size, props.className)}
    >
      {loading && (
        <div className={clsx(s.loader, loading && s.open)}>
          <Spin color={loadingColor} size="small" />
        </div>
      )}
      <span className={clsx(s.children, !loading && s.open)}>
        {props.children}
      </span>
    </button>
  );
}
