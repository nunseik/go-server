package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nunseik/go-server/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}

	type LoginResponse struct {
		ID 	  string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Email string `json:"email"`
		Token string `json:"token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "error decoding parameters", err)
		return
	}

	if params.ExpiresInSeconds <= 0 || params.ExpiresInSeconds > 3600 { // 1 hour in seconds
		params.ExpiresInSeconds = 3600 // Default to 1 hour
	}

	user, err := cfg.dbQueries.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, 401, "incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, 401, "incorrect email or password", err)
		return
	}

	timeDuration := time.Duration(params.ExpiresInSeconds) * time.Second

	token, err := auth.MakeJWT(user.ID, cfg.secretKey, timeDuration)
	if err != nil {
		respondWithError(w, 500, "error generating token", err)
		return
	}
	respondWithJSON(w, 200, LoginResponse{
		ID:        user.ID.String(),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Email:     user.Email,
		Token:     token,
	})
	
}