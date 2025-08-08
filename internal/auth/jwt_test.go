package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	secret := "supersecretkey"
	userID := uuid.New()
	expiresIn := 2 * time.Minute

	token, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	parsedID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}
	if parsedID != userID {
		t.Errorf("ValidateJWT returned wrong userID: got %v, want %v", parsedID, userID)
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	secret := "supersecretkey"
	userID := uuid.New()
	expiresIn := -1 * time.Second // already expired

	token, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Error("ValidateJWT did not fail for expired token")
	}
}

func TestValidateJWT_WrongSecret(t *testing.T) {
	secret := "supersecretkey"
	wrongSecret := "wrongsecret"
	userID := uuid.New()
	expiresIn := 2 * time.Minute

	token, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	_, err = ValidateJWT(token, wrongSecret)
	if err == nil {
		t.Error("ValidateJWT did not fail for token signed with wrong secret")
	}
}


