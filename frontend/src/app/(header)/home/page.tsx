import ChatSection from "@/components/chatLayout/chatSection";
import SideBar from "@/components/SideBar";

export default function HomePage() {
  return (
    <div className="flex h-screen overflow-hidden">
        {/* <!-- Sidebar --> */}
        <SideBar />
        {/* <!-- Main Chat Area --> */}
        <ChatSection />
    </div>
  )
}