package model

import "go.mongodb.org/mongo-driver/bson/primitive"
//Post is the structure of the Post Data
type Post struct{
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Title string `json:"title"`
	Body string `json:"body"`
}