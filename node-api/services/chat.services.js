import Chat from "../models/chat.model.js"
import User from "../models/user.model.js";

const findAllChats = async (query) => {
    try {
  
    const chats = await Chat.find({
        users: { $elemMatch: { $eq: userId } },
      })
    .populate("users", "name email ") // Populate user details
    .sort({ updatedAt: -1 }); // Sort by last updated
      if (!chats) {
        throw new Error("chat not found");
      }
    return chats;
    } catch (error) {
      throw new Error("Internal error: " + error.message);
    }
};

const findOrCreateChat = async (userId,senderId) => {
    try {
        let chatExists = await Chat.find({
            isGroup: false,
            $and: [
                { participants: { $elemMatch: { $eq: userId } } },
                { participants: { $elemMatch: { $eq: senderId } } },
            ],
        })
        .populate('users', '-password')
        //   chatExists = await User.populate(chatExists, {
        //     path: 'latestMessage.sender',
        //     select: 'name email profilePic',
        //   });
        if (chatExists.length > 0) {
            return chatExists
        } else {
            let data = {
                chatName: 'sender',
                participants: [userId, senderId],
                isGroup: false,
            };
            const newChat = await Chat.create(data);
            const chat = await Chat.find({ _id: newChat._id }).populate(
                'participants',
                '-password'
            );
            return chat;
        }
    } catch (error) {
      throw new Error("Internal error: " + error.message);
    }
};
  

export default {
    findAllChats,
    findOrCreateChat,
};