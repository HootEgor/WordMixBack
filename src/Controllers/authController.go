package Handlers

import (
	Models "WordMixBack/src/Model"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var newUser Models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	Models.AddNewUser(newUser.Login, newUser.Password, newUser.Language)
}
