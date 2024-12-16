import mongoose, { Schema } from "mongoose";

const MessageSchema = new Schema(
    {
      sender: { type: Schema.Types.ObjectId, ref: "User", required: true },
      content: { type: String, required: true },
      chatId: { type: Schema.Types.ObjectId, ref: "Chat", required: true },
      timestamp: { type: Date, default: Date.now },
      isRead: { type: Boolean, default: false },
    },
    { timestamps: true }
);
  
const Message = mongoose.model("Message", MessageSchema);
export default Message;
