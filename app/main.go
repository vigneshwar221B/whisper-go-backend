package main

import (
	"log"
	"net/http"
	"web-app/db"
	"web-app/helpers"

	"github.com/gorilla/mux"
)

func main() {
	//connect to db
	db.Connectmongo()

	// Init router
	r := mux.NewRouter()
	secretHandler := http.HandlerFunc(helpers.AddPost)

	// Route handles & endpoints
	r.HandleFunc("/", helpers.GetAllPosts).Methods("GET")
	r.HandleFunc("/register", helpers.SaveUser).Methods("POST")
	r.HandleFunc("/login", helpers.FindUser).Methods("POST")
	r.Handle("/addpost", helpers.IsAuthorized(secretHandler)).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))

}
