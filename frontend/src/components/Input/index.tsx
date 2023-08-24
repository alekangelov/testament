import s from "./index.module.css";
import { clsx } from "@data/clsx";

export default function Input({
  label,
  ...props
}: {
  label: string;
} & React.InputHTMLAttributes<HTMLInputElement>) {
  return (
    <div className={s.wrapper}>
      <div className={s.container}>
        <input
          type="text"
          {...props}
          placeholder=""
          className={clsx(props.className, s.input)}
        />
        <label className={s.label} htmlFor={props.id}>
          {label}
        </label>
      </div>
    </div>
  );
}
