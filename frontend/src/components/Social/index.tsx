import Button from "@components/Button";
import s from "./index.module.css";
import { Facebook, Google, Apple } from "@components/Icons";
import { useToaster } from "@hooks/useToaster";

const SIZE = 24;

export default function SocialLogin({
  onClick,
}: {
  onClick?: (type: "google" | "facebook" | "apple") => void;
}) {
  return (
    <div className={s.c}>
      <div className={s.line}>
        <span className={s.span}>or continue with</span>
      </div>
      <div className={s.wrap}>
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
