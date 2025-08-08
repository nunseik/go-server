package main

import (
	"net/http"
	"time"

	"github.com/nunseik/go-server/internal/auth"
)

func (cfg *apiConfig) handlerTokenRefresh(w http.ResponseWriter, r *http.Request) {
	
	refreshTokenHeader, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "unauthorized", err)
		return
	}

	refreshTokenDB, err := cfg.dbQueries.GetRefreshToken(r.Context(), refreshTokenHeader)
	if err != nil {
		respondWithError(w, 401, "unauthorized", err)
		return
	}

	if refreshTokenDB.ExpiresAt.Before(time.Now()) {
		respondWithError(w, 401, "refresh token expired", nil)
		return
	}

	if refreshTokenDB.RevokedAt.Valid {
		respondWithError(w, 401, "refresh token revoked", nil)
		return
	}

	accessToken, err := auth.MakeJWT(refreshTokenDB.UserID, cfg.secretKey, 3600*time.Second)
	if err != nil {
		respondWithError(w, 500, "internal server error", err)
		return
	}

	respondWithJSON(w, 200, map[string]string{
		"token": accessToken,
	})

}

func (cfg *apiConfig) handlerTokenRevoke(w http.ResponseWriter, r *http.Request) {
	refreshTokenHeader, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "unauthorized", err)
		return
	}

	err = cfg.dbQueries.RevokeRefreshToken(r.Context(), refreshTokenHeader)
	if err != nil {
		respondWithError(w, 500, "error revoking token", err)
		return
	}

	respondWithJSON(w, 204, nil)
}