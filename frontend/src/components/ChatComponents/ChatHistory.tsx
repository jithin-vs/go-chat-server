import WebSocketService from "@/utils/webSocketService";
import { useEffect, useState } from "react";
import ChatMessage from "./ChatMessage";
import { format } from "date-fns";

interface WebSocketProp {
  socket: WebSocketService | null;
}
interface Message {
  type: number;
  text: string;
}
function ChatHistory({ socket }: WebSocketProp) {
  const [messages, setMessages] = useState<any[]>([]);
  useEffect(() => {
    if (socket) {
      const handleMessage = (message: string) => {
        console.log("New message received:", message);
        const msg = JSON.parse(message);
        const timestamp = format(msg.time, "HH:mm");
        msg.time = timestamp;
        setMessages((prevMessages) => [...prevMessages, msg]);
      };
      socket.onMessage(handleMessage);
      return () => {};
    }
  }, [socket]);

  return (
    <>
      {messages.map((message) => {
        if (message?.type === 1) {
          return <ChatMessage message={message.body} time={message.time} />;
        } else if (message?.type === 2) {
          return (
            // <div
            //   className="flex flex-col items-center gap-2.5 m-8"
            //   key={message.body}
            // >
            //   <div className="flex flex-col gap-1 w-full max-w-[320px]">
            //     <div className="flex justify-center mb-4">
            //       <p className="text-sm font-normal text-gray-900 dark:text-white text-center">
            //         {message.body}
            //       </p>
            //     </div>
            //   </div>
            // </div>
            <div
              className="flex flex-col items-center gap-2.5 m-8"
              key={message.Id}
            >
              <p className="text-sm font-normal text-gray-900 dark:text-white text-center">
              <span >{message.time}</span>
              <span className="ml-4">{message.body}</span>
              </p>
            </div>
          );
        }
        return null;
      })}
    </>
  );
}
export default ChatHistory;
