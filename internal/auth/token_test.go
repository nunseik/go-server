package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken_Valid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer sometoken123")
	token, err := GetBearerToken(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "sometoken123" {
		t.Errorf("expected token 'sometoken123', got '%s'", token)
	}
}

func TestGetBearerToken_MissingHeader(t *testing.T) {
	headers := http.Header{}
	_, err := GetBearerToken(headers)
	if err == nil {
		t.Error("expected error for missing Authorization header, got nil")
	}
}

func TestGetBearerToken_InvalidFormat(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Token sometoken123")
	_, err := GetBearerToken(headers)
	if err == nil {
		t.Error("expected error for invalid Authorization header format, got nil")
	}
}

func TestGetBearerToken_EmptyToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer ")
	_, err := GetBearerToken(headers)
	if err == nil {
		t.Error("expected error for empty token, got nil")
	}
}