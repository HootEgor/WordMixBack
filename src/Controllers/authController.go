package Handlers

import (
	Models "WordMixBack/src/Model"
	"WordMixBack/src/Services"
	"cloud.google.com/go/firestore"
	"encoding/json"
	"net/http"
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
	}
}
