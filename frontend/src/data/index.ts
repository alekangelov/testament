export const get = async <T>(path: string) => {
  if (!process.env.API_URL) {
    throw new Error("API_URL is not defined");
  }
  return fetch(process.env.API_URL + path)
    .then((res) => res.json() as Promise<T>)
    .catch((e) => {
      console.error(e);
      return null;
    });
};

interface Settings {
  title: string;
  description: string;
}

export const getSettings = async () => {
  return get<Settings>("/v1/settings");
};

export const login = async ({
  email,
  password,
}: {
  email: string;
  password: string;
}): Promise<{
  token?: string;
  refreshToken?: string;
}> => {
  return fetch(process.env.API_URL + "/v1/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  })
    .then((res) => res.json())
    .catch((e) => {
      console.log(e);
      return {
        message: e.message,
      };
    });
};
