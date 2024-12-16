import { sendErrorResponse, sendResponse } from "../utils/responseUtils.js";
import chatsService from "../services/chat.services.js";

const fetchChats = async (req, res) => {
    try {
        const { userId} = req.body;
        console.log("User ID:", userId);
    
        if (!userId) {
          return sendErrorResponse(res, 400, "User ID is required");
        }
        const chats = await chatsService.findAllChats();

        const response = {
            message: "chat retrieved successfully",
            results: chats,
          };
        console.log("resp:",response.results)
        sendResponse(res, 200, response);

      } catch (error) {
        console.error("Error in findChats handler:", error.message);
        sendErrorResponse(res, 500, "Internal server error");
      }
};
  
const createChat = async (req, res) => {
    try {
        const { userId,senderId } = req.body;
        console.log("User ID:", userId);
    
        if (!userId) {
          return sendErrorResponse(res, 400, "User ID is required");
        }
        const chats = await chatsService.findOrCreateChat(userId,senderId);

        const response = {
            message: "chat created successfully",
            results: chats,
        };
        console.log("resp:",response.results)
      
        sendResponse(res, 200, response);

      } catch (error) {
        console.error("Error in findChats handler:", error.message);
        sendErrorResponse(res, 500, "Internal server error");
      }
};
  
export default {
    fetchChats,
    createChat,
}