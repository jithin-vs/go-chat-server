package handlers

import (
	"chatserver/services"
	"chatserver/utils"
	"chatserver/websocket"
	"fmt"
	"log"
	"net/http"
)

func GlobalChatHandler(w http.ResponseWriter, r *http.Request){
    pool := websocket.NewPool()
	go pool.Start()
	websocket.ServeWS(pool, w, r)
}

func GetChats(w http.ResponseWriter, r *http.Request){
	userID := r.URL.Query().Get("userId")
	fmt.Println("user",userID)
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
	chats, err := services.GetAllChats(r.Context(), userID)
	if err!= nil {
		log.Println("error getting data",err)
        utils.SendErrorResponse(w, http.StatusInternalServerError, "server error ")
        return
    }
	response := map[string]interface{}{
		"message": "fetched chats successfully",
		"data": chats,
	}
    // fmt.Println("fetched chats successfully",response);
    // Notify recipient about new chat
	// Respond with created chat
	utils.SendResponse(w, http.StatusOK, response);
}
