package services

import (
	"chatserver/db"
	"chatserver/models"
	"chatserver/utils"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func init() {
    userCollection = db.GetCollection("users")
}


func RegisterUser(ctx context.Context, user models.User) (interface{}, error) {
	hashedPassword,_ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	insertResult, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Error inserting user:", err)
		return nil, err
	}
	
	return insertResult, nil
}

func LoginUser(ctx context.Context, user models.User) (*models.User, error) {

    var foundUser models.User
    err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, fmt.Errorf("user not found")
        }
        return nil, err
    }
    
    match := utils.CheckPasswordHash(user.Password, foundUser.Password) 
    if !match {
        return nil, fmt.Errorf("incorrect password")
    }

    return &foundUser, nil
}

func FindUserById(ctx context.Context, userId string) (*models.User, error) {

    var foundUser models.User
    err := userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&foundUser)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, fmt.Errorf("user not found")
        }
        return nil, err
    }

    return &foundUser, nil
}

func IsUserExists(ctx context.Context, userId string) (bool, error) {

    var foundUser models.User
    objectID, err := primitive.ObjectIDFromHex(userId)
    if err!= nil {
        return false, err
    }
    err = userCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&foundUser)
    fmt.Printf("User query result - err: %v\n", err)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return false, nil
        }
        return false, err
    }
    return true, nil
}