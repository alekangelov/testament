import s from "./index.module.css";

const SIZE = {
  small: "18px",
  medium: "24px",
  large: "32px",
};

const COLOR = {
  foreground: "var(--foreground-rgb)",
  background: "var(--background-rgb)",
  surface: "var(--surface-rgb)",
};

export default function Spin({
  size = "small",
  color = "foreground",
}: {
  size?: keyof typeof SIZE;
  color?: keyof typeof COLOR;
}) {
  return (
    <div className={s.wrapper}>
      <div
        role="progressbar"
        style={
          {
            "--size": SIZE[size],
            "--color": COLOR[color],
          } as React.CSSProperties
        }
        className={s.loader}
      />
    </div>
  );
}
