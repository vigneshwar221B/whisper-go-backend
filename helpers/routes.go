package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-app/db"
	"web-app/model"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func IsAuthorized(endpoint http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		fmt.Println()
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				json.NewEncoder(w).Encode("error")
			}

			if token.Valid {
				endpoint.ServeHTTP(w, r)
			}
		} else {
			json.NewEncoder(w).Encode("not authorized")
		}
	})
}

//SaveUser is for '/register' post
func SaveUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t model.Register
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	//check if the user already register
	user := db.FindUserdb(t.Email)

	if user.Email != "" {
		json.NewEncoder(w).Encode("user already exists")
		return
	}

	//hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(t.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	t.Password = string(hashedPassword)

	//save it to DB
	db.AddUser(t)

	fmt.Println(t)

	//generate jwt and send it
	token, err := GenerateJWT(t.Email)

	if err != nil {
		fmt.Println("Failed to generate token")
	}

	json.NewEncoder(w).Encode(token)

}

func FindUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t model.Login
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	//find it in the db

	user := db.FindUserdb(t.Email)
	fmt.Println(user.Email)

	//check if the user exists
	if user.Email == "" {
		json.NewEncoder(w).Encode("can't find the user")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(t.Password))

	if err != nil {
		//incorrect password
		json.NewEncoder(w).Encode("incorrect password")
		return
	}
	//generate jwt and send it
	token, err := GenerateJWT(t.Email)

	if err != nil {
		fmt.Println("Failed to generate token")
	}

	json.NewEncoder(w).Encode(token)

}

func AddPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t model.Post
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	db.AddNewPostDB(t)

	json.NewEncoder(w).Encode("Post added successfully")
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts := db.GetAllPostsDB()

	json.NewEncoder(w).Encode(posts)
}
