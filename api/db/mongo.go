package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientInstance *mongo.Client
    clientInstanceError error
    mongoOnce sync.Once  
)


func ConnectMongoDB() (*mongo.Client, error) {
    err := godotenv.Load(".env")
    if err != nil{
     log.Fatalf("Error loading .env file: %s", err)
    }   
	uri := os.Getenv("DB_CONNECTION")
    mongoOnce.Do(func() {
        clientOptions := options.Client().ApplyURI(uri)
        var err error
        clientInstance, err = mongo.Connect(context.TODO(), clientOptions)
        if err != nil {
            clientInstanceError = err
            return
        }
        if clientInstanceError != nil {
            log.Println("Error connecting to MongoDB:", clientInstanceError)
            return
        }

        err = clientInstance.Ping(context.TODO(), nil)
        if err != nil {
            clientInstanceError = err
            log.Println("error connnecting to mongodb",err)
            return
        }
        fmt.Println("Connected to MongoDB")
    })

    return clientInstance, clientInstanceError
}

func GetCollection(name string) *mongo.Collection {
   
    if clientInstance == nil {
        // If client is nil, attempt to connect
        _, err := ConnectMongoDB()
        if err != nil {
            log.Fatalf("Failed to initialize MongoDB client: %v", err)
        }
    }


    // Get the database (you can pass your database name here)
    db := clientInstance.Database("chatserver")

    // Get the collection by name
    collection := db.Collection(name)
    return collection
}
