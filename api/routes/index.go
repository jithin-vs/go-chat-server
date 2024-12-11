package routers

import "chatserver/v2"
func SetupRoutes(wsHandler *chatws.Handler){
    AuthRoutes(wsHandler)
	// ChatRoutes()
}