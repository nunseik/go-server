package main

import (
	"strings"
	"errors"
	"slices"
)

func removeBadWords (sentence string) (string, error) {
	if sentence == "" {
		return "", errors.New("sentence cannot be empty")
	}
	var newWords []string
	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	splited := strings.Split(sentence, " ")
	for _, word := range splited {
		if slices.Contains(badWords, strings.ToLower(word)) {
			newWords = append(newWords, "****")
		} else {
			newWords = append(newWords, word)
		}

	}
	return strings.Join(newWords, " "), nil
}