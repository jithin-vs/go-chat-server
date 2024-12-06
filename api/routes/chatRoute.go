package routers

import (
	"chatserver/chatsocket"
	"net/http"
)

func ChatRoutes(){
	http.HandleFunc("/chat",chatsocket.NewChatServer().HandleConnection)
}