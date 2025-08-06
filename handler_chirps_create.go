package main

import (
	"net/http"
	"time"
	"github.com/google/uuid"
	"github.com/nunseik/go-server/internal/database"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
	
	cleanedBody, userId := handlerChirpsValidate(w, r)

	chirp, err := cfg.dbQueries.CreateChirp(r.Context(), database.CreateChirpParams{
		Body: cleanedBody,
		UserID: userId,
	})
	if err != nil {
		respondWithError(w, 500, "error creating chirp", err)
	}

	respondWithJSON(w, 201, Chirp{
		ID: chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body: chirp.Body,
		UserID: chirp.UserID,
	})
}

	