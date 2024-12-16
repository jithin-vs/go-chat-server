import express from "express";
import http from "http";
import { WebSocketServer } from "ws";
import bodyParser from "body-parser";
import connectDB from "./db/mongo.js";
import Routes from "./routes/routes.js";
import { configDotenv } from "dotenv";
import cors from "cors";

configDotenv();

const app = express();
const server = http.createServer(app);

app.use(cors());

app.use(express.json()); // To parse JSON bodies
app.use("/api", Routes);
// Middleware


// MongoDB Connection
connectDB();

// API Routes

// WebSocket server setup
const wss = new WebSocketServer({ server });

wss.on("connection", (ws) => {
  console.log("WebSocket connection established");
  ws.on("message", (message) => {
    console.log("Received:", message.toString());
    ws.send("Message received");
  });
});

// Start server
const PORT = 8080;
server.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`);
});
