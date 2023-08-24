"use client";
import { useCallback, useEffect, useRef, useState } from "react";

const delay = (ms = 1000): Promise<void> => {
  console.log(process.env.NODE_ENV);
  if (ms <= 0) return Promise.resolve();
  if (process.env.NODE_ENV !== "development") return Promise.resolve();
  return new Promise((res) =>
    setTimeout(() => {
      res();
    }, ms)
  );
};

export const useGet = <T, X>(path: string, variables: X) => {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);
  const abort = useRef(new AbortController());
  const get = useCallback(async () => {
    try {
      const res = await fetch(path, {
        signal: abort.current.signal,
        method: "GET",
      });
      const json = await res.json();
      setData(json);
    } catch (err: any) {
      console.log(err);
      setError(new Error(err?.message ?? "An error occurred!"));
    } finally {
      setLoading(false);
    }
  }, [path]);

  useEffect(() => {
    const currentAbort = abort.current;
    get();
    return () => {
      currentAbort.abort();
    };
  }, [get]);

  return { data, loading, error };
};

export function usePost<T, X>(
  path: string
): [
  (variables: X) => Promise<T>,
  {
    data: T | null;
    loading: boolean;
    error: Error | null;
  }
] {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);
  const abort = useRef(new AbortController());

  useEffect(() => {
    const currentAbort = abort.current;
    return () => {
      currentAbort.abort();
    };
  }, []);

  return [
    useCallback(
      async (variables: X) => {
        setLoading(true);
        await delay(2000);
        try {
          abort.current.abort();
          abort.current = new AbortController();
          const res = await fetch(path, {
            signal: abort.current.signal,
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(variables),
          });
          const json = await res.json();
          setData(json);
          setLoading(false);
          return json;
        } catch (err: any) {
          const error = new Error(err?.message ?? "An error occurred!");
          await delay(2000);
          setLoading(false);
          setError(error);
          throw error;
        }
      },
      [path]
    ),
    {
      data,
      loading,
      error,
    },
  ];
}
