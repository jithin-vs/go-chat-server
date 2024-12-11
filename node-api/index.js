import express from 'express';
import http from 'http';
import { WebSocketServer } from 'ws';

const app = express();
const server = http.createServer(app);

// WebSocket server setup
const wss = new WebSocketServer({ server });

wss.on('connection', (ws) => {
  console.log('WebSocket connection established');
  ws.on('message', (message) => {
    console.log('Received:', message.toString());
    ws.send('Message received');
  });
});


// Start server
server.listen(8080, () => {
  console.log('Server running on http://localhost:8080');
});
