import "./globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { getSettings } from "../data";
import ToastContainer from "@components/Toast";

const inter = Inter({ subsets: ["latin"] });

export async function generateMetadata(): Promise<Metadata> {
  const data = await getSettings();
  return {
    title: data?.title,
    description: data?.description,
  };
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ToastContainer />
        {children}
      </body>
    </html>
  );
}
