import Button from "@components/Button";
import s from "./index.module.css";

export default function SocialLogin() {
  return (
    <div className={s.c}>
      <div className={s.line}>
        <span className={s.span}>or continue with</span>
      </div>
      <div className={s.wrap}>
        <Button size="small" type="button" bordered>
          Google
        </Button>
        <Button size="small" type="button" bordered>
          Facebook
        </Button>
        <Button size="small" type="button" bordered>
          Twitter
        </Button>
        <Button size="small" type="button" bordered>
          Apple
        </Button>
      </div>
    </div>
  );
}
