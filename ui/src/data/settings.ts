import { get } from ".";

interface Settings {
  title: string;
  description: string;
}

export const getSettings = async () => {
  return get<Settings>("/v1/settings");
};
