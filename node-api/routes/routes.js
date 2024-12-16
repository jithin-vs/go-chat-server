import express from "express";
// import chatController from "./controllers/chatController";
// import wsController from "./controllers/wsController";
import { loginHandler, registerHandler } from "../controllers/authController.js";

const router = express.Router();

// Auth routes
router.post("/login", loginHandler);
router.post("/signup",registerHandler);

// Chat routes
// router.get("/chat", chatController.handleConnection);
// router.post("/chat/create", chatController.createChat);
// router.post("/chat/messages/:id", chatController.sendMessages);

// WebSocket chat routes
// router.post("/ws/chat/create", wsController.createRoom);

export default router;