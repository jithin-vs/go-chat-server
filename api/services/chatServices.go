package services

import (
	"chatserver/db"
	"chatserver/models"
	"context"
	"fmt"

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
	
	// Perform an aggregation to populate the `participantsDetails` field
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "_id", Value: newChat.ID},
		}},
	}
	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "participants"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "participants"},
		}},
	}
	
	cursor, err := chatCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		lookupStage,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode the result into the newChat object
	var populatedChats []models.Chat
	if err := cursor.All(ctx, &populatedChats); err != nil {
		return nil, err
	}
	if len(populatedChats) == 0 {
		return nil, fmt.Errorf("chat not found after insertion")
	}
    fmt.Println("chat",populatedChats)
	// Return the populated chat document
	return &populatedChats[0], nil
    
    // return newChat, nil
}