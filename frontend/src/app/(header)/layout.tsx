"use client"
import Header from "@/components/Header";
import { AuthProvider } from "@/context/AuthContext";

export default function HeaderLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
  return <AuthProvider>{children}</AuthProvider>;
}
