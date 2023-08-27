"use client";
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
  const [login, { loading }] = createFetch<
    {
      token: string;
      refreshToken: string;
    },
    {
      email: string;
      password: string;
    }
  >("/v1/auth/login");
  return (
    <ModalLayout>
      <main class={styles.main}>
        <h3 style={{ margin: 0 }}>Login</h3>
        <form
          class={styles.form}
          onSubmit={async (e) => {
            e.preventDefault();
            const data = new FormData(e.target as HTMLFormElement);
            const email = `${data.get("email")}`;
            const password = `${data.get("password")}`;

            if (!email || !password)
              return toast({
                title: "Error",
                message: "No email or password provided",
                type: "error",
              });
            try {
              const { token, refreshToken } = await login({
                email,
                password,
              });
              if (!token || !refreshToken) {
                throw new Error("There was an error logging you in");
              }
              toast({
                title: "You're in!",
                message: "We're logging you in now, sit tight",
                type: "success",
              });
              if (!callback) {
                return toast({
                  title: "We don't know what to do",
                  message:
                    "We don't know where to send you, so we're sending you nowhere",
                  type: "warning",
                });
              }
              const url = new URL(callback);
              url.searchParams.set("token", token);
              url.searchParams.set("refreshToken", refreshToken);
              window.location.href = url.href;
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
          <Button loadingColor="surface" loading={loading} type="submit">
            Login
          </Button>
        </form>
        <SocialLogin />

        <div class={styles.footer}>
          <a href="/register">
            Don&apos;t have an account? <b>Register</b>
          </a>
          <a href="/forgot-password">
            Forgot your password? <b>Reset it</b>
          </a>
        </div>
      </main>
    </ModalLayout>
  );
}
