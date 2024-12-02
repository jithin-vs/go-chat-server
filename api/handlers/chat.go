package handlers

import (
	"chatserver/websocket"
	"net/http"
)
func ChatHandler(w http.ResponseWriter, r *http.Request){
    pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWS(pool, w, r)
	})
}
