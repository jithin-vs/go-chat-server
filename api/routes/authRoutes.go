package routers

import (
	"chatserver/chatsocket"
	"chatserver/handlers"
     "chatserver/v2"
	"net/http"
)

func AuthRoutes(wsHandler *chatws.Handler) {
	// Chat routes
	// http.HandleFunc("/chat", handlers.HandleChat)
	// Auth routes
	http.HandleFunc("POST /login", handlers.LoginHandler)
	// http.HandleFunc("/", handlers.GlobalChatHandler)
	http.HandleFunc("POST /signup", handlers.RegisterHandler)
	http.HandleFunc("/chat",chatsocket.NewChatServer().HandleConnection)
	http.HandleFunc("/chat/create",chatsocket.NewChatServer().CreateChat)
	http.HandleFunc("/chat/messages/{id}",chatsocket.NewChatServer().SendMessages)
	http.HandleFunc("/ws/chat/create",wsHandler.CreateRoom)

}