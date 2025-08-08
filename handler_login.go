package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nunseik/go-server/internal/auth"
	"github.com/nunseik/go-server/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
	}

	type LoginResponse struct {
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Email     string `json:"email"`
		Token     string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "error decoding parameters", err)
		return
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

	JWTDuration := time.Duration(3600) * time.Second

	token, err := auth.MakeJWT(user.ID, cfg.secretKey, JWTDuration)
	if err != nil {
		respondWithError(w, 500, "error generating token", err)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, 500, "error generating refresh token", err)
		return
	}

	_, err = cfg.dbQueries.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		UserID: user.ID,
		Token: refreshToken,
	})
	if err != nil {
		respondWithError(w, 500, "error creating refresh token", err)
		return
	}

	respondWithJSON(w, 200, LoginResponse{
		ID:        user.ID.String(),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Email:     user.Email,
		Token:     token,
		RefreshToken: refreshToken,
	})
	
}