package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing Authorization header")
	}

	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	token := authHeader[7:]
	if token == "" {
		return "", fmt.Errorf("empty token in Authorization header")
	}

	return token, nil
	
}

func MakeRefreshToken() (string, error) {
	key := make([]byte, 32)
	rand.Read(key)
	encodedStr := hex.EncodeToString(key)
	return encodedStr, nil
}

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", fmt.Errorf("missing Authorization header")
	}

	if len(apiKey) < 7 || apiKey[:7] != "ApiKey " {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	return apiKey[7:], nil
}