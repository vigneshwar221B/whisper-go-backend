package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//connect to mongo

// Set client options
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")

// Connect to MongoDB
var Client, err = mongo.Connect(context.TODO(), clientOptions)

//choose your collection
//var collection = client.Database("webapp").Collection("user")

func Connectmongo() {
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
