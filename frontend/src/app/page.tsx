import ChatMessage from "@/components/ChatComponents/ChatMessage";
import ChatInput from "@/components/ChatComponents/ChatInput";
import ChatHistory from "@/components/ChatComponents/ChatHistory";
import Image from "next/image";

export default function Home() {
  return (
    <div>
      <ChatHistory /> 
      <ChatMessage />
      <ChatInput />
    </div>
  );
}
