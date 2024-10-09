import WebSocketService from "@/utils/webSocketService";
import { useEffect, useState } from "react";

interface WebSocketProp {
  socket: WebSocketService | null;
}
function ChatHistory({ socket }: WebSocketProp) {
  const [messages, setMessages] = useState<string[]>([]);
  useEffect(() => {
    if (socket) {
      // Register the onMessage callback
      const handleMessage = (message: string) => {
        console.log("New message received:", message);
        setMessages((prevMessages) => [...prevMessages, message]); // Append new message to state
      };
      socket.onMessage(handleMessage);
      return () => {};
    }
  }, [socket]);
  return (
    <div className="flex items-start gap-2.5 m-8">
      {/* <img className="w-8 h-8 rounded-full" src="/docs/images/people/profile-picture-3.jpg" alt="Jese image" /> */}
      <div className="flex flex-col gap-1 w-full max-w-[320px]">
        <div className="flex items-center space-x-2 rtl:space-x-reverse">
          <span className="text-sm font-semibold text-gray-900 dark:text-white">
            Bonnie Green
          </span>
          <span className="text-sm font-normal text-gray-500 dark:text-gray-400">
            11:46
          </span>
        </div>
        {messages.map((message) => {
          return (
            <div className="flex flex-col leading-1.5 mb-4 p-4 border-gray-200 bg-gray-100 rounded-e-xl rounded-es-xl dark:bg-gray-700">
              <p className="text-sm font-normal text-gray-900 dark:text-white">
                {" "}
                <span key={message}>{message}</span>;
              </p>
            </div>
          );
        })}
      </div>
    </div>
  );
}
export default ChatHistory;
