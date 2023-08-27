export const get = async <T>(path: string) => {
  if (!import.meta.env.VITE_API_URL) {
    throw new Error("API_URL is not defined");
  }
  return fetch(import.meta.env.VITE_API_URL + path)
    .then((res) => res.json() as Promise<T>)
    .catch((e) => {
      console.log(import.meta.env);
      console.error(e);
      return null;
    });
};
