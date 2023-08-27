import { Title } from "solid-start";
import { HttpStatusCode } from "solid-start/server";

export default function NotFound() {
  return (
    <main
      style={{
        display: "flex",
        "align-items": "center",
        "justify-content": "center",
        "min-height": "100vh",
        width: "100%",
      }}
    >
      <Title>Not Found</Title>
      <HttpStatusCode code={404} />
      <h1
        style={{
          "font-size": "3rem",
          "font-weight": "bold",
          margin: 0,
          display: "inline",
        }}
      >
        Page Not Found
      </h1>
    </main>
  );
}
