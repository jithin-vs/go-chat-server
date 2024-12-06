package chatsocket

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	SenderID    string `json:"senderId"`
	RecipientID string `json:"recipientId"`
	Content     string `json:"content"`
}

type Client struct {
	Conn *websocket.Conn
	Send chan Message
}

type ChatServer struct {
	clients     map[string]*Client
	mu          sync.RWMutex
	upgrader    websocket.Upgrader
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[string]*Client),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}

func (s *ChatServer) HandleConnection(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	fmt.Println("suer",userID)
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}
    fmt.Println("connection established")
	client := &Client{
		Conn: conn,
		Send: make(chan Message, 256),
	}

	s.mu.Lock()
	s.clients[userID] = client
	s.mu.Unlock()

	// Start goroutines for reading and sending messages
	go s.readMessages(userID, client)
	go s.writeMessages(client)
}

func (s *ChatServer) readMessages(userID string, client *Client) {
	defer func() {
		s.mu.Lock()
		delete(s.clients, userID)
		s.mu.Unlock()
		client.Conn.Close()
	}()

	for {
		var msg Message
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		// Validate sender and recipient
		if msg.SenderID != userID {
			log.Printf("Unauthorized sender: expected %s, got %s", userID, msg.SenderID)
			continue
		}

		// Route message to recipient
		s.routeMessage(msg)
	}
}

func (s *ChatServer) routeMessage(msg Message) {
	s.mu.RLock()
	recipientClient, exists := s.clients[msg.RecipientID]
	s.mu.RUnlock()

	if !exists {
		log.Printf("Recipient %s not connected", msg.RecipientID)
		return
	}

	select {
	case recipientClient.Send <- msg:
		log.Printf("Message sent to %s", msg.RecipientID)
	default:
		log.Printf("Failed to send message to %s: channel full", msg.RecipientID)
	}
}

func (s *ChatServer) writeMessages(client *Client) {
	defer client.Conn.Close()

	for msg := range client.Send {
		err := client.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Write error: %v", err)
			return
		}
	}
}