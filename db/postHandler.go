package db

import (
	"context"
	"fmt"
	"log"
	"web-app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

var postsCollection = Client.Database("webapp").Collection("posts")

func GetAllPostsDB() []*(model.Post){

	var posts []*(model.Post)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := postsCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

		// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return posts
}

func AddNewPostDB(obj model.Post) {
	post := model.Post{}

	post.Body = obj.Body
	post.Title = obj.Title
	post.ID = primitive.NewObjectID()

	insertResult, err := postsCollection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
