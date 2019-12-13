package db

import (
	"context"
	"fmt"
	"log"
	"web-app/model"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = Client.Database("webapp").Collection("user")

func AddUser(regObj model.Register) {
	//insert the user
	user := model.User{}
	user.Name = regObj.Name
	user.Email = regObj.Email
	user.Password = regObj.Password

	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func FindUserdb(email string) model.User{

	filter := bson.D{{"email", email}}
	var res = model.User{}

	err = collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
	//	log.Fatal(err)
		return model.User{}
	}

	fmt.Printf("Found a single document: %+v\n", res)
	return res
}
