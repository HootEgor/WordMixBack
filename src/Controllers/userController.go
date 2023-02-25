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

		userJson, err := ParseUserToJSON(user)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "%+v", userJson)
	}

}
