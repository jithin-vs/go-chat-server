package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/xid"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
 
func Upgrade(w http.ResponseWriter,r *http.Request)(*websocket.Conn, error){
      
	upgrader.CheckOrigin = func(r * http.Request) bool { return true }
	conn , err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("connection established");
	return conn, nil
}

func ServeWS(pool *Pool,w http.ResponseWriter, r *http.Request){

	// log.Println("Attempting to upgrade connection")
	conn, err := Upgrade(w, r)
	if err != nil {
		log.Printf("Error upgrading connection: %+v\n", err)
		http.Error(w, "Could not upgrade connection", http.StatusBadRequest)
		return
	}
id := xid.New()
client := &Client{
	ID : id.String(),
	Conn: conn,
	Pool: pool,
}
pool.Register <- client
client.Read()
}