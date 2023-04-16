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

type Service interface {
	GetUserInfo(client *firestore.Client, id string) (Models.User, error)
}

type Handler struct {
	Service Service
	Client  *firestore.Client
}

func (h *Handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := Services.GetUserInfo(h.Client, id)
	if err != nil {
		return
	}

	userJson, err := ParseToJSON(user)
	if err != nil {
		return
	}

	fprintf, err := fmt.Fprintf(w, "%+v", userJson)
	if err != nil {
		fmt.Errorf("rr.Body.String(): %v", fprintf)
		return
	}
}

func (h *Handler) NewUserScore(w http.ResponseWriter, r *http.Request) {
	var newScore Models.Score
	err := json.NewDecoder(r.Body).Decode(&newScore)
	if err != nil {
		return
	}

	err = Services.NewUserScore(h.Client, newScore)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", newScore)
}

func (h *Handler) GetLeaders(w http.ResponseWriter, r *http.Request) {
	scores, err := Services.GetLeaders(h.Client)
	if err != nil {
		return
	}

	scoresJson, err := ParseToJSON(scores)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "%+v", scoresJson)
}

func (h *Handler) GetUserHistory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	scores, err := Services.GetUserHistory(h.Client, id)
	if err != nil {
		return
	}

	scoresJson, err := ParseToJSON(scores)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "%+v", scoresJson)
}
