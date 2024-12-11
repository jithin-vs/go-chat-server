package main

import (
	"chatserver/db"
	"chatserver/routes"
    "chatserver/v2"
	"log"
	"net/http"

	"github.com/rs/cors"
)
 


func main()  {
	// handleChat()
	_, dberr := db.ConnectMongoDB()
	if dberr != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", dberr)
	}
	hub := chatws.NewHub()
	wsHandler := chatws.NewHandler(hub)
	routers.SetupRoutes(wsHandler)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{	"http://localhost:3000","http://localhost:3001",}, // Allow only your React app
		AllowCredentials: true,                             // Allow cookies/auth headers
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Accept"},
	})
	log.Println("server running on port 8080")
	err := http.ListenAndServe(":8080",corsHandler.Handler(http.DefaultServeMux))
	if err!= nil {
        log.Println(err)
    }

}