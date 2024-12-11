package chatws

import (
	"chatserver/services"
	"chatserver/utils"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
    return &Handler{hub: hub}
}

type CreateRoomReq struct {
	ID  primitive.ObjectID `json:"id"`
	SenderID string `json:"senderId"`
	RecipientID string `json:"recipientId"`
}


func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	// Validate request method
	if r.Method != http.MethodPost {
		utils.SendErrorResponse(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	// Parse request body
	var req CreateRoomReq
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

	req.ID = primitive.NewObjectID()
    newID := req.ID.Hex()
	h.hub.Rooms[newID] = &Room{
		ID: req.ID,
		Participants: make(map[string]*Participant),
	}

	// Check if chat already exists
	// existingChat, err := services.FindChatByParticipants(r.Context(), req.SenderID, req.RecipientID)
	// if err != nil {
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Error checking existing chat")
	// 	return
	// }
	// // If chat exists, return existing chat
	// if existingChat != nil {
	// 	utils.SendResponse(w, http.StatusOK, existingChat)
	// 	return
	// }
	
	// Create new chat
	// newChat, err := services.CreateChat(r.Context(), req.SenderID, req.RecipientID)
	// if err != nil {
	// 	fmt.Println("chat error:", err)
	// 	utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create chat")
	// 	return
	// }

	// response := map[string]interface{}{
	// 	"message": "Chat created successfully",
	// 	"data": newChat,
	// }
    // fmt.Println("chat created successfully",response);
    // // Notify recipient about new chat
	// // Respond with created chat
	// utils.SendResponse(w, http.StatusCreated, response)
}
