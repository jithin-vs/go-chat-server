package services

import (
	"chatserver/db"
	"chatserver/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var messageCollection *mongo.Collection

func init() {
    messageCollection = db.GetCollection("messages")
}

func StoreMessage(ctx context.Context, message models.Messages) (*models.Messages, error){
	insertResult, err := messageCollection.InsertOne(ctx, message)
	if err != nil {
		log.Println("Error inserting user:", err)
		return nil, err
	}
	// Populate the message with the inserted ID
	message.ID = insertResult.InsertedID.(primitive.ObjectID)

	// Return the populated message
	return &message, nil

}