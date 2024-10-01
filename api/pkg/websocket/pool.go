package websocket

import "fmt"

type Pool struct{
   Register		chan *Client
   Unregister 	chan *Client
   Clients		map[*Client]bool
   Broadcast    chan Message
}
func NewPool() *Pool {
   return &Pool{
	Register: make(chan *Client),
	Unregister: make(chan *Client),
	Clients: make(map[*Client]bool),
	Broadcast: make(chan Message),
   }
}

func (pool *Pool) Start(){
	for range pool.Clients{ 
		select{
		case client:= <-pool.Register:
			pool.Clients[client] = true ;
			fmt.Println("total number of users",len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type:1,Body: "New user joined!!"})
			}
		case client:= <-pool.Unregister:
			 delete(pool.Clients, client)
			fmt.Println("total number of users",len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type:1,Body: "User Disconnected!!"})
			}
		case message:= <-pool.Broadcast:
			 
			fmt.Println("sendinf message to users")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
			}
		}
	}
  }
}