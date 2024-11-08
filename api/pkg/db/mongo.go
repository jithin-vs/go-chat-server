package db

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientInstance *mongo.Client
    clientInstanceError error
    mongoOnce sync.Once  
)


func ConnectMongoDB() (*mongo.Client, error) {
	uri := "mongodb+srv://jithinvs1045:122333@cluster0.563o5.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
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
            return
        }
        fmt.Println("Connected to MongoDB")
    })

    return clientInstance, clientInstanceError
}