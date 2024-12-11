package chatsocket

import (
	"chatserver/models"
	"chatserver/services"
	"chatserver/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// // type Message struct {
// // 	SenderID    string `json:"senderId"`
// // 	RecipientID string `json:"recipientId"`
// // 	Content     string `json:"content"`
// // }

type Client struct {
	Conn *websocket.Conn
	Send chan models.Messages
}

type ChatServer struct {
	clients     map[string]*Client
	mu          sync.RWMutex
	upgrader    websocket.Upgrader
}

type CreateChatRequest struct {
	SenderID    string `json:"senderId"`
	RecipientID string `json:"recipientId"`
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
	if userID =="undefined" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "User ID is required")
		return
	}
    isExist,Usererr := services.IsUserExists(r.Context(),userID)
	if Usererr!= nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "server error ")
		return
    }
	if !isExist {
		utils.SendErrorResponse(w, http.StatusNotFound, "User does not exist")
		return
    }
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("here")
		log.Printf("Upgrade error: %v", err)
		return
	}
    fmt.Println("connection established")
	client := &Client{
		Conn: conn,
		Send: make(chan models.Messages, 256),
	}

	s.mu.Lock()
	s.clients[userID] = client
	s.mu.Unlock()

}

func (s *ChatServer) CreateChat(w http.ResponseWriter, r *http.Request) {
	// Validate request method
	if r.Method != http.MethodPost {
		utils.SendErrorResponse(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	// Parse request body
	var req CreateChatRequest
	err := utils.ParseRequest(r, &req)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
    fmt.Println("chat users",req);
	// Validate sender and recipient IDs
	if req.SenderID == "" || req.RecipientID == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Sender and Recipient IDs are required")
		return
	}
	// Check if sender exists
	senderExists, err := services.IsUserExists(r.Context(), req.SenderID)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error checking sender")
		return
	}
	if !senderExists {
		utils.SendErrorResponse(w, http.StatusNotFound, "Sender does not exist")
		return
	}

	// Check if recipient exists
	recipientExists, err := services.IsUserExists(r.Context(), req.RecipientID)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error checking recipient")
		return
	}
	if !recipientExists {
		utils.SendErrorResponse(w, http.StatusNotFound, "Recipient does not exist")
		return
	}

	// Check if chat already exists
	existingChat, err := services.FindChatByParticipants(r.Context(), req.SenderID, req.RecipientID)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error checking existing chat")
		return
	}
	// If chat exists, return existing chat
	if existingChat != nil {
		utils.SendResponse(w, http.StatusOK, existingChat)
		return
	}
	
	// Create new chat
	newChat, err := services.CreateChat(r.Context(), req.SenderID, req.RecipientID)
	if err != nil {
		fmt.Println("chat error:", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create chat")
		return
	}
	fmt.Println("chats exists? ",newChat);

    // recipientDetails,err := services.FindUserById(r.Context(),req.RecipientID)
	// if err != nil {
	// 	utils.SendErrorResponse(w, http.StatusNoContent, "Recipient not found")
	// 	return
	// }

	response := map[string]interface{}{
		"message": "Chat created successfully",
		"data": newChat,
	}
    fmt.Println("chat created successfully",response);
    // Notify recipient about new chat
	// Respond with created chat
	utils.SendResponse(w, http.StatusCreated, response)
}

func (s *ChatServer) readMessages(userID string, client *Client) {
	defer func() {
		s.mu.Lock()
		delete(s.clients, userID)
		s.mu.Unlock()
		client.Conn.Close()
	}()

	for {
		var msg models.Messages
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		// Validate sender and recipient
		userObjectID, _ := primitive.ObjectIDFromHex(userID)
		if msg.SenderID != userObjectID {
			log.Printf("Unauthorized sender: expected %s, got %s", userID, msg.SenderID)
			continue
		}

		// Route message to recipient
		s.routeMessage(msg)
	}
}

func (s *ChatServer) routeMessage(msg models.Messages) {
	s.mu.RLock()
	recipientClient, exists := s.clients[msg.RecipientID.Hex()]
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
     // Loop through messages to send and write to WebSocket connection
    log.Println("Starting write goroutine for", client.Conn.RemoteAddr().String())
	for msg := range client.Send {
		err := client.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Write error: %v", err)
			return
		}
	}
}

func (s *ChatServer) SendMessages(w http.ResponseWriter, r *http.Request) {
    var msg models.Messages
    if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
        utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid message format")
        return
    }
    fmt.Println("here..")
    // Validate chat existence (as in your previous implementation)
    isExist, chatErr := services.IsChatExists(r.Context(), msg.ChatID)
    if chatErr != nil {
        utils.SendErrorResponse(w, http.StatusInternalServerError, "Server error")
        return
    }

    if !isExist {
        utils.SendErrorResponse(w, http.StatusNotFound, "Chat does not exist")
        return
    }

    // Store message
    storedMessage, err := services.StoreMessage(r.Context(), msg)
    if err != nil {
        utils.SendErrorResponse(w, http.StatusInternalServerError, "Error storing message")
        return
    }

    recipientIDStr := msg.RecipientID.Hex()
    // Attempt to send via WebSocket if recipient is online
    s.mu.RLock()
    recipientClient, exists := s.clients[recipientIDStr]
    s.mu.RUnlock()
    println("exists?",exists)
    if exists {
        select {
        case recipientClient.Send <- *storedMessage:
            log.Println("Message sent to recipient via WebSocket")
        default:
            log.Println("Recipient channel full")
        }
    }

    // Respond to sender
    response := map[string]interface{}{
        "message": "Message sent",
        "data":    storedMessage,
    }
    utils.SendResponse(w, http.StatusOK, response)
}
