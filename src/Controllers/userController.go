package Handlers

import (
	"WordMixBack/src/Services"
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserInfo(client *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		user, err := Services.GetUserInfo(client, id)
		if err != nil {
			return
		}

		userJson, err := ParseToJSON(user)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "%+v", userJson)
	}
}

func GetLeaders(client *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scores, err := Services.GetLeaders(client)
		if err != nil {
			return
		}

		scoresJson, err := ParseToJSON(scores)
		if err != nil {
			return
		}
		fmt.Fprintf(w, "%+v", scoresJson)
	}
}
