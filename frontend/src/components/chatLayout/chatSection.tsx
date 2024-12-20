"use client"
import { Chat, ChatMessageInterface } from "@/interfaces/ChatMessage";
import MessageInput from "./TextField";
import { useEffect, useRef, useState } from "react";
import WebSocketService from "@/utils/webSocketService";
import { PORT, SOCKET_URL } from "@/config";
import axiosInstance from "@/config/axiosInstance";
import { useAuth } from "@/context/AuthContext";
import axios from "axios";
import { User } from "@/interfaces/userInterface";
import MessageContainer from "./IncomingMessage";

export default function ChatSection() {

  const [messages, setMessages] = useState<ChatMessageInterface[]>([]);
  const [message, setMessage] = useState(""); 
  const [chats, setChats] = useState([]);
  const [chat, setChat] = useState<Chat | null>();
  const [isTyping, setIsTyping] = useState(false); // To track if someone is currently typing
  const [selfTyping, setSelfTyping] = useState(false); // To track if the current user is typing
  const typingTimeoutRef = useRef<NodeJS.Timeout | null>(null);
  const { user } = useAuth();
  let currentUserId = user?.id;
  const [receiver,setReceiver] = useState<User|null>();
  const [socket,setSocket] = useState<WebSocketService | null>(null);
  
  // useEffect(() => {
  //   if (user) {      
  //     console.log("user", user);
  //     const service = new WebSocketService(`${SOCKET_URL}/chat?userId=${currentUserId}`);
  //     setSocket(service);
  //     // handleAddChat()
  //     return () => {
  //       service.close();  
  //     };
  //   }
  // }, [user]);

  // useEffect(() => {
  //   console.log('Current user in this component:', user)
  // }, [user])

  const getChats = () => { 
    axios.post(`${PORT}/chat/users?userId=${currentUserId}`, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    .then((response) => {
      if (response.status === 200) {
        console.log("Chats loaded successfully:", response.data.data);
        // Handle the response, such as updating state or navigating
        setChats(response.data.data);
        // router.push('/chat');
        return;
      }
      console.log("Chat created but unexpected status:", response);
    })
    .catch((error) => {
      console.error("Error creating chat:", error);
      // setErrors([error.response?.data?.error || "Unknown error occurred"]);
    });
  }
  
  useEffect(() => {
    if (user) {
      handleAddChat()
      getChats()
    }
  }, [user])

  // Function to send a chat message
  const sendChatMessage = async () => {
    // If no current chat ID exists or there's no socket connection, exit the function
    const apiData = {
      senderId: currentUserId,
      content: message,
      chatId: chat?.id,
      recipientId: receiver?.id
    }
    if (!chat?.id || !socket) return;
    console.log("chat", chat);
    // Optionally send via WebSocket as well
    if (socket?.isOpen()) {
      socket.emit('message', apiData); // Ensure correct payload
  }
    // axios.post(`${PORT}/chat/messages/${chat?.id}`,apiData,{
    //   headers: {
    //     'Content-Type': 'application/json',
    //   },
    // })
    // .then((response) => {
    //   if (response.status === 200) {
    //     console.log("message sended successfully:", response.data.data);
    //     // Handle the response, such as updating state or navigating
    //     if (response.status === 200) {
    //       console.log("Message sent successfully:", response.data.data);
    //       setChats(response.data.data);
    //    }
    //     setChats(response.data.data);
    //     // router.push('/chat');
    //     return;
    //   }
    //   console.log("Chat created but unexpected status:", response);
    // })
    // .catch((error) => {
    //   console.error("Error creating chat:", error);
    //   // setErrors([error.response?.data?.error || "Unknown error occurred"]);
    // });

    
  };

  const handleOnMessageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    // Update the message state with the current input value
    setMessage(e.target.value);

    // If socket doesn't exist or isn't connected, exit the function
    if (!socket) return;

    // Check if the user isn't already set as typing
    // if (!selfTyping) {
    //   // Set the user as typing
    //   setSelfTyping(true);

    //   // Emit a typing event to the server for the current chat
    //   socket.emit(TYPING_EVENT, currentChat.current?._id)
    // }

    // Clear the previous timeout (if exists) to avoid multiple setTimeouts from running
    if (typingTimeoutRef.current) {
      clearTimeout(typingTimeoutRef.current);
    }

    // Define a length of time (in milliseconds) for the typing timeout
    const timerLength = 3000;

    // Set a timeout to stop the typing indication after the timerLength has passed
    typingTimeoutRef.current = setTimeout(() => {
      // Emit a stop typing event to the server for the current chat
      // socket.emit(STOP_TYPING_EVENT, currentChat.current?._id);

      // Reset the user's typing state
      // setSelfTyping(false);
    }, timerLength);
  };

  const handleAddChat = async () => {
    let selectedUser ="67513c7e7b461271f5fa0d05" ;
    // let selectedUser = "674daaf70c9707db1aa341ae";
    if (!selectedUser) return;
   if(currentUserId)
    axios.post(`${PORT}/chat/create?userId=${currentUserId}`, { senderId:currentUserId, recipientId:selectedUser}, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    .then((response) => {
      if (response.status === 201) {
        console.log("Chat created successfully:", response.data.data);
        // Handle the response, such as updating state or navigating
        setChat(response.data.data);
        const participantDetails = response.data.participantDetails;
        const receiver = participantDetails.find(
          (participant: { id: string }) => participant.id !== currentUserId
        );
        
        if (receiver) {
          setReceiver(receiver);
        } else {
          console.error("Receiver not found");
        }
        // router.push('/chat');
        return;
      }
      if (response.status === 200) {
        console.log("Chat loaded successfully:", response.data);
        // Handle the response, such as updating state or navigating         
        setChat(response.data);

        const participantDetails = response.data.participantDetails;
        const receiver = participantDetails.find(
          (participant: { id: string }) => participant.id !== currentUserId
        );
        
        if (receiver) {
          setReceiver(receiver);
        } else {
          console.error("Receiver not found");
        }
        // setReciever(response.data.participantDetails[1]);
        // router.push('/chat');
        return;
      }
      console.log("Chat created but unexpected status:", response);
    })
    .catch((error) => {
      console.error("Error creating chat:", error);
      // setErrors([error.response?.data?.error || "Unknown error occurred"]);
    });
  
    // setChats((prevChats) => [...prevChats, response.data]);
    // setIsModalOpen(false);
  };
  return (
    <div className="flex-1">
      {/* <!-- Chat Header --> */}
      <header className="bg-white p-4 text-gray-700">
        <h1 className="text-2xl font-semibold">{receiver?.username}</h1>
      </header>

      {/* <!-- Chat Messages --> */}
      <div className="h-screen overflow-y-auto p-4 pb-36">  
        {/* <!-- Incoming Message --> */}
        <MessageContainer />
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
        <MessageInput
          placeholder="Message"
          value={message}
          onChange={handleOnMessageChange}
          onKeyDown={(e:React.KeyboardEvent<HTMLInputElement>) => {
            if (e.key === "Enter") {
              sendChatMessage();
            }
          }}
          />
      </footer>
    </div>
  );
}
