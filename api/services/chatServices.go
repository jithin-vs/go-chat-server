package services

import (
	"chatserver/db"
	"chatserver/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var chatCollection *mongo.Collection

func init() {
    chatCollection = db.GetCollection("chats")
}
func FindChatByParticipants(ctx context.Context, senderID, recipientID string) (*models.Chat, error) {
	senderObjectID, err := primitive.ObjectIDFromHex(senderID)
	if err != nil {
		return nil, err
	}
	recipientObjectID, err := primitive.ObjectIDFromHex(recipientID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
        "$or": []bson.M{
            {
                "participants": bson.M{
                    "$all": []primitive.ObjectID{senderObjectID, recipientObjectID},
                },
            },
        },
    }
    
    var chat models.Chat
    err = chatCollection.FindOne(ctx, filter).Decode(&chat)
    if err == mongo.ErrNoDocuments {
        return nil,nil  // Custom error indicating no chat exists
    }
    if err != nil {
        return nil, err  // Other database errors
    }
    populatedChat,err := aggregateParticipants(ctx,chat.ID)
	if err != nil {
		return nil, err 
	}
    return &populatedChat[0], nil
}

func IsChatExists(ctx context.Context, chatId primitive.ObjectID) (bool, error) {

    var foundChat models.Chat
    objectID :=chatId
    err := chatCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&foundChat)
    // fmt.Printf("User query result - err: %v\n", err)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return false, nil
        }
        return false, err
    }
    return true, nil
}

func CreateChat(ctx context.Context, senderID, recipientID string) (*models.Chat, error) {
	senderObjectID, err := primitive.ObjectIDFromHex(senderID)
	if err != nil {
		return nil, err
	}
	recipientObjectID, err := primitive.ObjectIDFromHex(recipientID)
	if err != nil {
		return nil, err
	}
	
	newChat := &models.Chat{
		Participants: []primitive.ObjectID{senderObjectID, recipientObjectID},
		Type: "single",
	}
	
    
	insertResult, err := chatCollection.InsertOne(ctx, newChat)
	if err != nil {
		return nil, err
	}
	
	// Cast the inserted ID to ObjectID and assign it to chatId
	newChat.ID = insertResult.InsertedID.(primitive.ObjectID)
	

    chats,err := aggregateParticipants(ctx, newChat.ID)
	if err != nil{
		log.Println("error aggregating",err)
		return nil,err
	}
	fmt.Println("chats aggregated \n",chats)
    // return newChat, nil
	return &chats[0], nil
}

func GetAllChats(ctx context.Context, userID string) ([]models.Chat, error) {
	// Define the aggregation pipeline
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err!= nil {
        return nil, err
    }
	filter := bson.M{
		"participants": bson.M{
			"$in": []primitive.ObjectID{userObjectID},
		},
	}
	
	// Sort by the latest message or creation timestamp
	opts := options.Find().SetSort(bson.D{{Key:"updatedAt",Value:-1}})
	
	var chats []models.Chat
	cursor, err := chatCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	
	if err = cursor.All(ctx, &chats); err != nil {
		return nil, err
	}

    // Populate participants for each chat
	for i, chat := range chats {
		// Call the aggregateParticipants function to populate participants for each chat
		populatedChat, err := aggregateParticipants(ctx, chat.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to populate participants for chat ID %s: %v", chat.ID.Hex(), err)
		}
		// Replace the original chat with the populated one
		chats[i] = populatedChat[0]
	}
	return chats, nil

}

func aggregateParticipants(ctx context.Context,chatId primitive.ObjectID) ([]models.Chat, error) {
		// Perform an aggregation to populate the `participantsDetails` field
		matchStage := bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "_id", Value: chatId},
			}},
		}
		lookupStage := bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "participants"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "participantDetails"},
			}},
		}
	
		// fmt.Printf("Aggregation pipeline: %+v\n", mongo.Pipeline{matchStage, lookupStage})
		cursor, err := chatCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			lookupStage,
		})
		if err != nil {
			log.Printf("Error decoding cursor: %v\n", err)
			return nil, err
		}
		defer cursor.Close(ctx)
	
		// Decode the result into the newChat object
		var populatedChats []models.Chat
		if err := cursor.All(ctx, &populatedChats); err != nil {
			return nil, err
		}
		// fmt.Println("chat --->",populatedChats)
		if len(populatedChats) == 0 {
			return nil, fmt.Errorf("chat not found after insertion")
		}
		// Return the populated chat document
		return populatedChats, nil
}