package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct{
	ID		string 	`json:"id"`
	Conn 	*websocket.Conn
	Pool    *Pool
	mut  	sync.Mutex
}

type Message struct{
	Type  	int `json:"type"`
	Body    string `json:"body"`
}

func (c *Client) Read(){

	defer func(){
		c.Pool.Unregister <- c
		c.Conn.Close()
	 }()

	for{
		MessageType,b,err := c.Conn.ReadMessage()
		fmt.Println("message :",MessageType)
		if err != nil{
			log.Println(err)
			return
		}
		message := Message{Type:MessageType,Body: string(b)}
		c.Pool.Broadcast <- message
	}
 
}

