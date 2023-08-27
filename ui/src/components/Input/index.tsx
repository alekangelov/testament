import { JSX } from "solid-js";
import s from "./index.module.css";
import { clsx } from "@data/clsx";

export default function Input({
  label,
  ...props
}: {
  label: string;
} & JSX.HTMLElementTags["input"]) {
  return (
    <div class={s.wrapper}>
      <div class={s.container}>
        <input
          type="text"
          {...props}
          placeholder=""
          class={clsx(props.class, s.input)}
        />
        <label class={s.label} for={props.id}>
          {label}
        </label>
      </div>
    </div>
  );
}
