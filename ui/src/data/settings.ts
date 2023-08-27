import { get } from ".";

interface Settings {
  title?: string;
  description?: string;
  logo?: string;
  socials?: string[];
  colors?: {
    primary?: string;
    success?: string;
    warning?: string;
    info?: string;
    error?: string;
  };
}

export const getSettings = async () => {
  return get<Settings>("/v1/settings");
};
