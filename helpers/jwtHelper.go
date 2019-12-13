package helpers

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//just write anything you want
var mySigningKey = []byte("iamtheboneofmysword")

//GenerateJWT returns the JWT
func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
