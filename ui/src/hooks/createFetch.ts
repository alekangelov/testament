import delay from "@utils/delay";
import { Accessor, createSignal, onMount } from "solid-js";

export default function createRequest<T, X>(
  url: string,
  options?: RequestInit
): [
  (variables: X) => Promise<T>,
  {
    data: Accessor<T | null>;
    loading: Accessor<boolean>;
    error: Accessor<Error | null>;
  }
] {
  const [data, setData] = createSignal<T | null>(null);
  const [loading, setLoading] = createSignal(false);
  const [error, setError] = createSignal<Error | null>(null);
  let abort = new AbortController();

  onMount(() => {
    return () => {
      abort.abort();
    };
  });

  return [
    async (variables: X) => {
      const path = import.meta.env.VITE_API_URL + url;
      setLoading(true);
      await delay(500);
      try {
        abort.abort();
        abort = new AbortController();
        const res = await fetch(path, {
          signal: abort.signal,
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          ...options,
          body: JSON.stringify(variables),
        });
        if (!res.ok) throw await res.json();
        const json = await res.json();
        setData(json);
        setLoading(false);
        return json;
      } catch (err: any) {
        const error = new Error(err?.message ?? "An error occurred!");
        setLoading(false);
        setError(error);
        throw error;
      }
    },
    {
      data,
      loading,
      error,
    },
  ];
}
