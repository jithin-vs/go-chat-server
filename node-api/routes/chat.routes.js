import express from "express";
import chatController from "../controllers/chat.controller.js";

const router = express.Router();

router.route('/')
    .get(chatController.fetchChats)
    .post(chatController.createChat)


export default router;