package routers

import (
	"net/http"
	"chatserver/handlers"
)

func SetupRoutes() {
	// Chat routes
	// http.HandleFunc("/chat", handlers.HandleChat)
	// Auth routes
	http.HandleFunc("POST /login", handlers.LoginHandler)
	http.HandleFunc("/home", handlers.ChatHandler)
	http.HandleFunc("POST /signup", handlers.RegisterHandler)
}
