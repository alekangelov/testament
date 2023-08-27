import { getSettings } from "@data/settings";
import s from "./modal.module.css";
import { ParentProps, createEffect, createResource } from "solid-js";
import { Head } from "solid-start";

const extractRgbFromColor = (color: string) => {
  const [r, g, b] = color
    .replace("rgb(", "")
    .replace(")", "")
    .split(",")
    .map((c) => parseInt(c));
  return [r, g, b].join(",");
};

const prepColor = (color?: string) => {
  if (!color) return ``;
  return `${extractRgbFromColor(color)}`;
};

export default function ModalLayout({ children }: ParentProps) {
  const [data] = createResource(() => {
    return getSettings();
  });
  let logoHolder: HTMLDivElement;
  createEffect(() => {
    if (!logoHolder || !data()?.logo) return;
  });
  return (
    <>
      <Head>
        <title>{data()?.title || "Solid Start"}</title>
        <meta name="description" content={data()?.description} />
        <style data-from-server>
          {`
            :root {
              --primary: ${prepColor(data()?.colors?.primary)};
              --success: ${prepColor(data()?.colors?.success)};
              --error: ${prepColor(data()?.colors?.error)};
              --warning: ${prepColor(data()?.colors?.warning)};
              --info: ${prepColor(data()?.colors?.info)};
              --warning: ${prepColor(data()?.colors?.warning)};
            }
          `}
        </style>
      </Head>
      <div class="background" />
      <div class={s.container}>
        <header class={s.header}>
          <div class={s.logo} innerHTML={data()?.logo ?? ""} />{" "}
          {data()?.title && <h1 class={s.title}>{data()?.title}</h1>}
        </header>
        <div class={s.wrapper}>{children}</div>
      </div>
    </>
  );
}
