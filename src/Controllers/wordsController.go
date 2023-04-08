package Handlers

import (
	Models "WordMixBack/src/Model"
	"WordMixBack/src/Services"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) AddNewWords(w http.ResponseWriter, r *http.Request) {
	var newWords []Models.Word
	err := json.NewDecoder(r.Body).Decode(&newWords)
	if err != nil {
		return
	}

	err = Services.AddNewWords(h.Client, newWords)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", newWords)
}

func (h *Handler) GetWords(w http.ResponseWriter, r *http.Request) {
	words, err := Services.GetWords(h.Client)
	if err != nil {
		return
	}

	wordsJson, err := ParseToJSON(words)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%+v", wordsJson)
}
