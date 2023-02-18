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
		fmt.Fprintf(w, "%+v", user)
	}

}

//func GetLeadersByRegion(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "%+v", Models.Users)
//}
