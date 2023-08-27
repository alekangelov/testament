import { getSettings } from "@data/settings";
import s from "./modal.module.css";
import { ParentProps, createEffect, createResource } from "solid-js";

export default function ModalLayout({ children }: ParentProps) {
  const [data] = createResource(() => {
    return getSettings();
  });
  createEffect(() => {
    console.log(data());
  });
  return (
    <>
      <div class="background" />
      <div class={s.container}>
        <div class={s.wrapper}>
          {data()?.title && <h1>{data()?.title}</h1>}
          {children}
        </div>
      </div>
    </>
  );
}
