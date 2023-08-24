"use client";
import Input from "@components/Input";
import styles from "./page.module.css";
import Button from "@components/Button";
import SocialLogin from "@components/Social";
import { usePost } from "@hooks/useFetch";

export default function Index() {
  const [login, { loading }] = usePost<
    {
      token: string;
      refreshToken: string;
    },
    {
      email: string;
      password: string;
    }
  >("/api/auth/login");
  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    "use client";
    e.preventDefault();
    const data = new FormData(e.target as HTMLFormElement);
    const email = `${data.get("email")}`;
    const password = `${data.get("password")}`;
    if (!email || !password) return console.error("Missing email or password");
    const { token, refreshToken } = await login({
      email,
      password,
    });
  }
  console.log({ loading });
  return (
    <main className={styles.main}>
      <h3>Login</h3>
      <form className={styles.form} onSubmit={handleSubmit}>
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

      <div className={styles.footer}>
        <a href="/register">
          Don&apos;t have an account? <b>Register</b>
        </a>
        <a href="/forgot-password">
          Forgot your password? <b>Reset it</b>
        </a>
      </div>
    </main>
  );
}
