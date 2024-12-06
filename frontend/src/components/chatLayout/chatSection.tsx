"use client"
import { ChatMessageInterface } from "@/interfaces/ChatMessage";
import IncomingMessage from "./IncomingMessage";
import OutgoingMessage from "./OutgoingMessage";
import MessageInput from "./TextField";
import { useEffect, useRef, useState } from "react";
import WebSocketService from "@/utils/webSocketService";
import { PORT, SOCKET_URL } from "@/config";
import axiosInstance from "@/config/axiosInstance";
import { useAuth } from "@/context/AuthContext";
import axios from "axios";

export default function ChatSection() {

  const [messages, setMessages] = useState<ChatMessageInterface[]>([]);
  const [message, setMessage] = useState(""); 
  const [isTyping, setIsTyping] = useState(false); // To track if someone is currently typing
  const [selfTyping, setSelfTyping] = useState(false); // To track if the current user is typing
  const typingTimeoutRef = useRef<NodeJS.Timeout | null>(null);
  const { user } = useAuth();
 
  const [socket,setSocket] = useState<WebSocketService | null>(null);
  
  useEffect(() => {
    // const service = new WebSocketService(`${SOCKET_URL}/home`);
    // setSocket(service);
    handleAddChat()
    // return () => {
    //   service.close();  
    // };
  }, []);

  // const handleOnMessageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  //   // Update the message state with the current input value
  //   setMessage(e.target.value);

  //   // If socket doesn't exist or isn't connected, exit the function
  //   if (!socket || !isConnected) return;

  //   // Check if the user isn't already set as typing
  //   if (!selfTyping) {
  //     // Set the user as typing
  //     setSelfTyping(true);

  //     // Emit a typing event to the server for the current chat
  //     socket.emit(TYPING_EVENT, currentChat.current?._id)
  //   }

  //   // Clear the previous timeout (if exists) to avoid multiple setTimeouts from running
  //   if (typingTimeoutRef.current) {
  //     clearTimeout(typingTimeoutRef.current);
  //   }

  //   // Define a length of time (in milliseconds) for the typing timeout
  //   const timerLength = 3000;

  //   // Set a timeout to stop the typing indication after the timerLength has passed
  //   typingTimeoutRef.current = setTimeout(() => {
  //     // Emit a stop typing event to the server for the current chat
  //     socket.emit(STOP_TYPING_EVENT, currentChat.current?._id);

  //     // Reset the user's typing state
  //     setSelfTyping(false);
  //   }, timerLength);
  // };

  const handleAddChat = async () => {
    let selectedUser ;
    let currentUserId = user?.id;
    if (!selectedUser) return;
  
    const response = await axios.post(`${PORT}/chat?userId=${currentUserId}`, {
      participants: [currentUserId, selectedUser],
    });
    // setChats((prevChats) => [...prevChats, response.data]);
    // setIsModalOpen(false);
  };
  return (
    <div className="flex-1">
      {/* <!-- Chat Header --> */}
      <header className="bg-white p-4 text-gray-700">
        <h1 className="text-2xl font-semibold">Alice</h1>
      </header>

      {/* <!-- Chat Messages --> */}
      <div className="h-screen overflow-y-auto p-4 pb-36">
        {/* <!-- Outgoing Message --> */}
        <OutgoingMessage />
        {/* <!-- Incoming Message --> */}
        <IncomingMessage />
      </div>

      {/* <!-- Chat Input --> */}
      <footer className="bg-white border-t border-gray-300 p-4 absolute bottom-0 w-3/4">
        {/* <div className="flex items-center">
          <input
            type="text"
            placeholder="Type a message..."
            className="w-full p-2 rounded-md border border-gray-400 focus:outline-none focus:border-blue-500"
          />
          <button className="bg-indigo-500 text-white px-4 py-2 rounded-md ml-2">
            Send
          </button>
        </div> */}
        {/* <MessageInput
          placeholder="Message"
          value={message}
          onChange={handleOnMessageChange}
          // onKeyDown={(e) => {
          //   if (e.key === "Enter") {
          //     sendChatMessage();
          //   }
          // }}
          /> */}
      </footer>
    </div>
  );
}