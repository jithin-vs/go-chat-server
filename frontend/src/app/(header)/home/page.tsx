"use client"
import ChatSection from "@/components/chatLayout/chatSection";
import SideBar from "@/components/SideBar";
import WebSocketService from "@/utils/webSocketService";
import { useEffect, useState } from "react";

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