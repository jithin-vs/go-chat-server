package services

import (
	"chatserver/db"
	"chatserver/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var chatCollection *mongo.Collection

func init() {
    chatCollection = db.GetCollection("chats")
}
func FindChatByParticipants(ctx context.Context, senderID, recipientID string) (*models.Chat, error) {
    filter := bson.M{
        "$or": []bson.M{
            {
                "participants": bson.M{
                    "$all": []string{senderID, recipientID},
                },
            },
        },
    }
    
    var chat models.Chat
    err := chatCollection.FindOne(ctx, filter).Decode(&chat)
    if err == mongo.ErrNoDocuments {
        return nil,nil  // Custom error indicating no chat exists
    }
    if err != nil {
        return nil, err  // Other database errors
    }
    
    return &chat, nil
}

func CreateChat(ctx context.Context, senderID, recipientID string) (*models.Chat, error) {
	newChat := &models.Chat{
        Participants: []string{senderID, recipientID},
        Type: "single",
    }
    
    _, err := chatCollection.InsertOne(ctx, newChat) 
    if err != nil {
        return nil, err
    }
    
    return newChat, nil
}