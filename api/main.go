package main

import (
	"chatserver/pkg/websocket"
	// "fmt"
	"log"
	"net/http"
	"github.com/rs/xid"
)
 
func serveWS(pool *websocket.Pool,w http.ResponseWriter, r *http.Request){

		// log.Println("Attempting to upgrade connection")
		conn, err := websocket.Upgrade(w, r)
		if err != nil {
			log.Printf("Error upgrading connection: %+v\n", err)
			http.Error(w, "Could not upgrade connection", http.StatusBadRequest)
			return
		}
	id := xid.New()
	client := &websocket.Client{
		ID : id.String(),
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}
func handleChat(){
    pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main()  {
	handleChat()
	log.Println("server running on port 8080")
	err := http.ListenAndServe(":8080",nil)
	if err!= nil {
        log.Println(err)
    }
}