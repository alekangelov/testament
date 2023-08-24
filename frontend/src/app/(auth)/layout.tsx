import { getSettings } from "../../data";
import s from "./layout.module.css";

export default async function Layout({
  children,
}: {
  children: React.ReactNode;
}) {
  const data = await getSettings();
  return (
    <>
      <div className="background" />
      <div className={s.container}>
        <div className={s.wrapper}>
          {data?.title && <h1>{data.title}</h1>}
          {children}
        </div>
      </div>
    </>
  );
}
