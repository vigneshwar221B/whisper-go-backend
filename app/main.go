package main

import (
	"log"
	"net/http"
//	"os"
	"web-app/db"
	"web-app/helpers"

//	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // All origins
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		AllowedHeaders: []string{"*"},
	})

	// headersOk := handlers.AllowedHeaders([]string{"*"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//start the server
	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))

}
