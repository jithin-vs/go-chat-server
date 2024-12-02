package websocket

import (
	"fmt"
	"log"
	"sync"
    "time"
	"github.com/gorilla/websocket"
)

type Client struct{
	ID		string 	`json:"id"`
	Conn 	*websocket.Conn
	Pool    *Pool
	mut  	sync.Mutex
}

type Message struct{
	Id      string `json:"id"`
	Type  	int `json:"type"`
	Body    string `json:"body"`
	TimeStamp time.Time `json:"time"`
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
		uid := fmt.Sprintf("%s-%d", c.ID, time.Now().Unix())
		message := Message{Id:uid,Type:MessageType,Body: string(b),TimeStamp: time.Now()}
		c.Pool.Broadcast <- message
	}
 
}

