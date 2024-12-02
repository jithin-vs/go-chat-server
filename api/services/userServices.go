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
	
	response := map[string]interface{}{
		"message": "User registered successfully",
		"data": insertResult,
	}
	return response, nil
}

func LoginUser(ctx context.Context, user models.User) (interface{}, error) {
    // Step 1: Get the user collection
    collection := db.GetCollection("users")

    // Step 2: Find the user by email (or username)
    var foundUser models.User
    err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            log.Println("User not found")
            return nil, fmt.Errorf("user not found")
        }
        log.Println("Error finding user:", err)
        return nil, err
    }

    match := utils.CheckPasswordHash(user.Password, foundUser.Password) 
    if !match {
        log.Println("Incorrect password")
        return nil, fmt.Errorf("incorrect password")
    }

    response := map[string]interface{}{
        "message": fmt.Sprintf("User %s logged in successfully", foundUser.Username),
        "userID":  foundUser,
        // "token": token, // Uncomment if you're using JWT
    }
    return response, nil
}
