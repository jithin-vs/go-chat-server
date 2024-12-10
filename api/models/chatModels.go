package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Participants []primitive.ObjectID  `json:"participants"`
	ParticipantsDetails []User `json:"participantDetails" bson:"participantDetails"`
	CreatedAt    time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updated_at"`
	Type         string    `json:"type" bson:"type"`
}

type Messages struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	ChatID       primitive.ObjectID   `json:"chatId"`
	Content      string    `json:"content" bson:"content"`
	SenderID     primitive.ObjectID   `json:"senderId" bson:"senderId"`
	RecipientID  primitive.ObjectID   `json:"recipientId" bson:"recipientId"`
	// ParticipantsDetails []User `json:"participantDetails" bson:"participantDetails"`
	CreatedAt    time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updated_at"`
	// Type         string    `json:"type" bson:"type"`
}