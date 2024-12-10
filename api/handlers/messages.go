package handlers

import (
	"chatserver/models"
	"chatserver/services"
	"chatserver/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRequest struct {
	SenderID    primitive.ObjectID `json:"senderId"`
	ChatID      primitive.ObjectID `json:"chatId"`
	Content     string             `json:"content"`
}

func SendMessages(w http.ResponseWriter, r *http.Request){
	chatId := r.PathValue("id")
	fmt.Println("user",chatId)
    
	var req models.Messages
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return 
    }
	if chatId =="undefined" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Chat ID is required")
		return
	}
	
	chatObjectId,_ := primitive.ObjectIDFromHex(chatId)
    isExist,chatErr := services.IsChatExists(r.Context(),chatObjectId)
	if chatErr!= nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "server error ")
		return
    }
	if !isExist {
		utils.SendErrorResponse(w, http.StatusNotFound, "chat does not exist")
		return
    }
    
	//saving message to db
	messages, err := services.StoreMessage(r.Context(),req)
	if err != nil {
	  log.Println("error storing data",err)
	  utils.SendErrorResponse(w, http.StatusInternalServerError, "server error: error storing message")
	}
	response := map[string]interface{}{
		"message": "message stored",
		"data": messages,
	}
    // fmt.Println("fetched chats successfully",response);
    // Notify recipient about new chat
	// Respond with created chat
	utils.SendResponse(w, http.StatusOK, response);
}
