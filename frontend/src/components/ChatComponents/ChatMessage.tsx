import WebSocketService from "@/utils/webSocketService";
import { useEffect, useState } from "react";

interface WebSocketProp {
  socket: WebSocketService | null;
}
function ChatMessage({ socket }: WebSocketProp) {
  const [message, setMessage] = useState('');
  useEffect(() => {
    if (socket) {
      const handleMessage = (message: string) => {
        console.log("New message received:", message);
        setMessage(message);
      };

      socket.onMessage(handleMessage);

 
      return () => {
      };
    }
  }, [socket]);
  return (
    <div className="flex items-start justify-end gap-2.5 m-8">
      <div className="flex flex-col gap-1 w-full max-w-[320px]">
        <div className="flex items-center space-x-2 rtl:space-x-reverse justify-end">
          <span className="text-sm font-semibold text-gray-900 dark:text-white">
            Bonnie Green
          </span>
          <span className="text-sm font-normal text-gray-500 dark:text-gray-400">
            11:46
          </span>
        </div>
        <div className="flex flex-col leading-1.5 p-4 border-gray-200 bg-gray-100 rounded-s-xl rounded-se-xl dark:bg-gray-700">
          <p className="text-sm font-normal text-gray-900 dark:text-white">
            {" "}
            {message}
          </p>
        </div>
      </div>
      {/* <img
        className="w-8 h-8 rounded-full"
        src="/docs/images/people/profile-picture-3.jpg"
        alt="Jese image"
      /> */}
    </div>
  );
}
export default ChatMessage;
