package routers

import (
	"chatserver/chatsocket"
	"chatserver/handlers"
	"net/http"
)

func AuthRoutes() {
	// Chat routes
	// http.HandleFunc("/chat", handlers.HandleChat)
	// Auth routes
	http.HandleFunc("POST /login", handlers.LoginHandler)
	http.HandleFunc("/", handlers.GlobalChatHandler)
	http.HandleFunc("POST /signup", handlers.RegisterHandler)
	http.HandleFunc("/chat",chatsocket.NewChatServer().HandleConnection)
}
