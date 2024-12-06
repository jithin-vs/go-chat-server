package services

import (
	"chatserver/db"
	"chatserver/models"
	"chatserver/utils"
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
    collection := db.GetCollection("users")

    var foundUser models.User
    err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
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
