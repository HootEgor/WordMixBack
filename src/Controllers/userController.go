package Handlers

import (
	Models "WordMixBack/src/Model"
	"WordMixBack/src/Services"
	"cloud.google.com/go/firestore"
	"encoding/json"
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

func NewUserScore(client *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newScore Models.Score
		err := json.NewDecoder(r.Body).Decode(&newScore)
		if err != nil {
			return
		}

		err = Services.NewUserScore(client, newScore)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "%+v", newScore)
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

func GetUserHistory(client *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		scores, err := Services.GetUserHistory(client, id)
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
