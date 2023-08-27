import { createEffect } from "solid-js";
import { redirect } from "solid-start";

export default function Home() {
  const response = redirect("/login");
  return;
}
