package main

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/nunseik/go-server/internal/auth"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.dbQueries.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, 500, "error getting chirps", err)
	}

	chirps := make([]Chirp, len(dbChirps))
	for i, dbChirp := range dbChirps {
		chirps[i] = Chirp{
			ID: dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body: dbChirp.Body,
			UserID: dbChirp.UserID,
		}
	}

	respondWithJSON(w, 200, chirps)
}

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirpIDStr := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		respondWithError(w, 404, "invalid chirp ID", err)
		return
	}
	chirp, err := cfg.dbQueries.GetSingleChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, 404, "error getting chirp by id", err)
		return
	}
	respondWithJSON(w, 200, Chirp{
		ID: chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body: chirp.Body,
		UserID: chirp.UserID,
	})
}

func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, r *http.Request) {
	chirpIDStr := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		respondWithError(w, 404, "invalid chirp ID", err)
		return
	}

	tokenHeader, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "unauthorized", err)
		return
	}

	userID, err := auth.ValidateJWT(tokenHeader, cfg.secretKey)
	if err != nil {
		respondWithError(w, 401, "unauthorized", err)
		return
	}

	chirp, err := cfg.dbQueries.GetSingleChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, 404, "error getting chirp by id", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, 403, "forbidden", nil)
		return
	}

	err = cfg.dbQueries.DeleteChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, 500, "error deleting chirp", err)
		return
	}

	respondWithJSON(w, 204, nil)
}