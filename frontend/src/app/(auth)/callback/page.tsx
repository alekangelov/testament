const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

export default async function Index(props: {
  searchParams?: Record<string, string>;
  params?: Record<string, string>;
}) {
  const { searchParams = {} } = props;
  const { token, refreshToken } = searchParams;
  await sleep(1000);
  return (
    <main>
      <h1>Callback</h1>

      <code>{token}</code>
      <br />
      <code>{refreshToken}</code>
    </main>
  );
}
