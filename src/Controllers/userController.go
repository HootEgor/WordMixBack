package Handlers

import (
	Models "WordMixBack/src/Model"
	"fmt"
	"net/http"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	id := 1
	user, err := Models.GetUserByID(id)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "%+v", user)
}

func GetLeadersByRegion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v", Models.Users)
}
