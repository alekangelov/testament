export default function delay(ms: number) {
  if (import.meta.env.PROD === true) return;
  return new Promise((resolve) => setTimeout(resolve, ms));
}
