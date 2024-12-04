import Header from "@/components/Header";

export default function HeaderLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
  return (
    <>
      {/* <Header /> */}
      {children}
    </>
  );
}
