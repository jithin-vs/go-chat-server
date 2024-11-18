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