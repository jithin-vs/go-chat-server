import { User } from "./userInterface";

export interface ChatMessageInterface {
    _id: string;
    sender: Pick<User, "id" | "email" | "name">;
    content: string;
    chat: string;
    attachments: {
      url: string;
      localPath: string;
      _id: string;
    }[];
    createdAt: string;
    updatedAt: string;
}
  
export interface Chat {
  id: string;
  participants: string[]; // Array of user IDs
  participantsDetails?: User[]; // Populated user details, optional
  createdAt: Date;
  type: "single" | "group"; // Specify the type of chat
}