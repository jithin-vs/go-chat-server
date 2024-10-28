import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";

const readexRegular = localFont({
  src: "./fonts/ReadexPro-Regular.ttf",
  variable: "--font-readex-regular",
  weight: "100 900",
});

const readexMedium = localFont({
  src: "./fonts/ReadexPro-Medium.ttf",
  variable: "--font-readex-medium",
  weight: "100 900",
});

const readexSemiBold = localFont({
  src: "./fonts/ReadexPro-SemiBold.ttf",
  variable: "--font-readex-semibold",
  weight: "100 900",
});

const readexBold = localFont({
  src: "./fonts/ReadexPro-Bold.ttf",
  variable: "--font-readex-bold",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${readexRegular.variable} ${readexMedium.variable} ${readexSemiBold.variable} ${readexBold.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
