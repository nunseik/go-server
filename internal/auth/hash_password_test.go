package auth

import (
	"testing"
)

func TestHashPasswordAndCheckPasswordHash(t *testing.T) {
	password := "mySecretPassword123!"

	// Test hashing does not return error and produces a non-empty hash
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}
	if hash == "" {
		t.Fatal("HashPassword returned empty hash")
	}

	// Test correct password matches hash
	if err := CheckPasswordHash(password, hash); err != nil {
		t.Errorf("CheckPasswordHash failed for correct password: %v", err)
	}

	// Test incorrect password does not match hash
	wrongPassword := "wrongPassword"
	if err := CheckPasswordHash(wrongPassword, hash); err == nil {
		t.Error("CheckPasswordHash did not fail for incorrect password")
	}
}

func TestHashPasswordDifferentHashes(t *testing.T) {
	password := "samePassword"
	hash1, err1 := HashPassword(password)
	hash2, err2 := HashPassword(password)
	if err1 != nil || err2 != nil {
		t.Fatalf("HashPassword returned error: %v, %v", err1, err2)
	}
	if hash1 == hash2 {
		t.Error("HashPassword returned same hash for same password (should be different due to salt)")
	}
}
