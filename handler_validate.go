package main

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) (string, uuid.UUID) {
	type parameters struct {
		Body string `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "error decoding parameters", err)
		return "", uuid.UUID{}
	}
	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long", nil)
		return "", uuid.UUID{}
	}
	cleanedBody, err := removeBadWords(params.Body)
	if err != nil {
		respondWithError(w, 500, "error cleaning body", err)
	}

	return cleanedBody, params.UserID
}