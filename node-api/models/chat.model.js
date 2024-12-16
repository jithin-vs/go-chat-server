import mongoose, { Schema,} from "mongoose";

const ChatSchema = new Schema(
    {
      participants: [{ type: Schema.Types.ObjectId, ref: "User", required: true }],
      isGroupChat: { type: Boolean, default: false },
      chatName: { type: String, default: "" },
      lastMessage: { type: Schema.Types.ObjectId, ref: "Message" },
    },
    { timestamps: true }
);
  
const Chat = mongoose.model("Chat", ChatSchema);

export default Chat ;