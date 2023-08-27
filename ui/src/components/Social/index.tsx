import Button from "@components/Button";
import s from "./index.module.css";
import { Facebook, Google, Apple } from "@components/Icons";
import { createResource } from "solid-js";
import { getSettings } from "@data/settings";

const SIZE = 24;

export default function SocialLogin({
  onClick,
}: {
  onClick?: (type: "google" | "facebook" | "apple") => void;
}) {
  const [settings] = createResource(getSettings);
  console.log(settings());
  return (settings()?.socials?.length ?? 0) === 0 ? (
    <div />
  ) : (
    <div class={s.c}>
      <div class={s.line}>
        <span class={s.span}>or continue with</span>
      </div>
      <div class={s.wrap}>
        <Button
          size="small"
          type="button"
          bordered
          onClick={() => onClick?.("google")}
        >
          <Google />
        </Button>
        <Button
          size="small"
          type="button"
          bordered
          onClick={() => onClick?.("facebook")}
        >
          <Facebook />
        </Button>
        <Button
          size="small"
          type="button"
          bordered
          onClick={() => onClick?.("apple")}
        >
          <Apple />
        </Button>
      </div>
    </div>
  );
}
