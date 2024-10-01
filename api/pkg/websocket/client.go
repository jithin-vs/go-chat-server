package websocket

import (
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
		if err != nil{
			log.Println(err)
		}
		message := Message{Type:MessageType,Body: string(b)}
		c.Pool.Broadcast <- message
	}
}

