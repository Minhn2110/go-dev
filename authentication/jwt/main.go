package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

var mySigningKey = []byte("my_secret_key")

// Create JWT Token
func CreateToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "john.doe",
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Fprintf(w, "Error creating the token")
		return
	}

	fmt.Fprintf(w, tokenString)
}

// Verify JWT Token
func VerifyToken(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Fprintf(w, "Invalid token")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Fprintf(w, "Token is valid. Username: %s", claims["username"])
	} else {
		fmt.Fprintf(w, "Invalid token")
	}
}

func main() {
	http.HandleFunc("/create", CreateToken)
	http.HandleFunc("/verify", VerifyToken)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
