package chatws

import (
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Participant struct {
	Conn *websocket.Conn
	Message chan Message
	ID string `json:"id"`
}

type Message struct {
	ID primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Content      string    `json:"content" bson:"content"`
	SenderID     primitive.ObjectID   `json:"senderId" bson:"senderId"`
}