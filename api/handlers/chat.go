package handlers

import (
	"chatserver/websocket"
	"net/http"
)
func ChatHandler(w http.ResponseWriter, r *http.Request){
    pool := websocket.NewPool()
	go pool.Start()
	websocket.ServeWS(pool, w, r)
}
