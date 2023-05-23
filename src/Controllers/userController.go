package Handlers

import (
	Models "WordMixBack/src/Model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type UserService interface {
	GetUserInfo(ctx context.Context, id string) (Models.User, error)
	GetLeaders(ctx context.Context) ([]Models.Score, error)
	GetUserHistory(ctx context.Context, id string) ([]Models.Score, error)
	NewUserScore(ctx context.Context, score Models.Score) error
}

func (h *HttpHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := h.userService.GetUserInfo(r.Context(), id)
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

func (h *HttpHandler) NewUserScore(w http.ResponseWriter, r *http.Request) {
	var newScore Models.Score
	err := json.NewDecoder(r.Body).Decode(&newScore)
	if err != nil {
		return
	}

	err = h.userService.NewUserScore(r.Context(), newScore)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", newScore)
}

func (h *HttpHandler) GetLeaders(w http.ResponseWriter, r *http.Request) {
	scores, err := h.userService.GetLeaders(r.Context())
	if err != nil {
		return
	}

	scoresJson, err := ParseToJSON(scores)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "%+v", scoresJson)
}

func (h *HttpHandler) GetUserHistory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	scores, err := h.userService.GetUserHistory(r.Context(), id)
	if err != nil {
		return
	}

	scoresJson, err := ParseToJSON(scores)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "%+v", scoresJson)
}
