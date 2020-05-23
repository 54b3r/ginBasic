package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var MySigningKey = os.Getenv("OAUTH_SIGNING_KEY")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Marvin Matos"
	// set the token exp time to 30 minutes from time issued
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	if MySigningKey == "" {
		MySigningKey = "supersecretkey"
	}

	tokenString, err := token.SignedString([]byte(MySigningKey))
	if err != nil {
		log.Printf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	fmt.Println("Client Application")

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error Generating token string")
	}
	fmt.Println(tokenString)
}
