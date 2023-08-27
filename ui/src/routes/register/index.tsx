import Input from "@components/Input";
import styles from "./index.module.css";
import Button from "@components/Button";
import SocialLogin from "@components/Social";
import createFetch from "@hooks/createFetch";
import { useToaster } from "@hooks/useToaster";
import { useSearchParams } from "solid-start";
import ModalLayout from "@components/Layout/ModalLayout";

export default function Index() {
  const [{ callback = "" }, setParams] = useSearchParams();
  const { add: toast } = useToaster();
  const [register, { loading }] = createFetch<
    {
      message: string;
      success: boolean;
    },
    {
      email: string;
      password: string;
      passwordConfirmation: string;
    }
  >("/v1/auth/register");
  return (
    <ModalLayout>
      <main class={styles.main}>
        <hgroup>
          <h3 style={{ margin: 0 }}>Register</h3>
        </hgroup>
        <form
          class={styles.form}
          onSubmit={async (e) => {
            e.preventDefault();
            const data = new FormData(e.target as HTMLFormElement);
            const email = `${data.get("email")}`;
            const password = `${data.get("password")}`;
            const passwordConfirmation = `${data.get("passwordConfirmation")}`;

            if (!email || !password || !passwordConfirmation)
              return toast({
                title: "Error",
                message: "No email or password provided",
                type: "error",
              });
            try {
              const { message } = await register({
                email,
                password,
                passwordConfirmation,
              });

              toast({
                title: "Great!",
                message: message,
                type: "success",
              });
            } catch (e: any) {
              toast({
                title: "Oops, there's something wrong",
                message: e?.message || "An error occurred",
                type: "error",
              });
            }
          }}
        >
          <Input label="Email" id="email" name="email" type="email" required />
          <Input
            label="Password"
            id="password"
            name="password"
            type="password"
            required
          />
          <Input
            label="Password (again)"
            id="passwordConfirmation"
            name="passwordConfirmation"
            type="password"
            required
          />
          <Button loadingColor="surface" loading={loading} type="submit">
            Register
          </Button>
        </form>
        <SocialLogin />

        <div class={styles.footer}>
          <a href="/login">
            Already have an account? <b>Login</b>
          </a>
          <a href="/forgot-password">
            Forgot your password? <b>Reset it</b>
          </a>
        </div>
      </main>
    </ModalLayout>
  );
}
