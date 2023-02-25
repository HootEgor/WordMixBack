package Handlers

import (
	Models "WordMixBack/src/Model"
	"WordMixBack/src/Services"
	"cloud.google.com/go/firestore"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func Register(client *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser Models.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			return
		}

		err = Services.AddNewUser(client, newUser)
		if err != nil {
			return
		}

		token, err := GenerateJWT(newUser)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "%+v", token)
	}
}

func Login(client *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser Models.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			return
		}

		user, err := Services.LoginUser(client, newUser)
		if err != nil || user == nil {
			http.Error(w, "user not found", 404)
			return
		}

		token, err := GenerateJWT(user[0])
		if err != nil {
			return
		}

		fmt.Fprintf(w, "%+v", token)
	}
}

func GenerateJWT(user Models.User) (string, error) {

	claims := jwt.MapClaims{}
	claims["login"] = user.Login
	claims["password"] = user.Password
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("my-secret-key")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(tokenString)
	if err != nil {
		return "", err
	}

	jsonString := string(jsonData)

	return jsonString, nil
}
