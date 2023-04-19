package Handlers

import (
	Models "WordMixBack/src/Model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WordService interface {
	AddNewWords(ctx context.Context, words []Models.Word) ([]Models.Word, error)
	GetWords(ctx context.Context) ([]Models.Word, error)
}

func (h *HttpHandler) AddNewWords(w http.ResponseWriter, r *http.Request) {
	var newWords []Models.Word

	err := json.NewDecoder(r.Body).Decode(&newWords)
	if err != nil {
		return
	}

	newWords, err = h.wordService.AddNewWords(r.Context(), newWords)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", newWords)
}

func (h *HttpHandler) GetWords(w http.ResponseWriter, r *http.Request) {
	words, err := h.wordService.GetWords(r.Context())
	if err != nil {
		return
	}

	wordsJson, err := ParseToJSON(words)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", wordsJson)
}
