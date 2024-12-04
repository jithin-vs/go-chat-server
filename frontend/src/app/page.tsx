"use client"
import ChatMessage from "@/components/ChatComponents/ChatMessage";
import ChatInput from "@/components/ChatComponents/ChatInput";
import ChatHistory from "@/components/ChatComponents/ChatHistory";
import WebSocketService from "@/utils/webSocketService";
import { useEffect, useState } from "react";

export default function Home() {
  const [wsService, setWsService] = useState<WebSocketService | null>(null);
  
  useEffect(() => {
    const service = new WebSocketService("ws://localhost:8080/home");
    setWsService(service);
    return () => {
      service.close();  
    };
  }, []);
  
  return (
    <div>
      <ChatHistory socket = {wsService} /> 
      {/* <ChatMessage socket = {wsService} /> */}
      <ChatInput   socket = {wsService}/>
    </div>
  );
}
