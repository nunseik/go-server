package main

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"github.com/nunseik/go-server/internal/database"
)

func (cfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event   string `json:"event"`
		Data    struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithError(w, 204, "unsupported event", nil)
		return
	}

	userID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = cfg.dbQueries.UpdateUserChirpyRed(r.Context(), database.UpdateUserChirpyRedParams{
		IsChirpyRed: true,
		ID:          userID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, 204, nil)


}