package entity

import (
	"testing"
)

func TestNewCombination_ValidCombination(t *testing.T) {
	votationID := "votation123"
	firstMovieID := "movie1"
	secondMovieID := "movie2"
	chosenMovieID := "movie1"

	combination, err := NewCombination(votationID, firstMovieID, secondMovieID, chosenMovieID)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if combination == nil {
		t.Error("Expected a valid Combination, but got nil.")
	}

	if combination == nil ||
		combination.VotationID != votationID ||
		combination.FirstMovieID != firstMovieID ||
		combination.SecondMovieID != secondMovieID ||
		combination.ChosenMovieID != chosenMovieID {
		t.Error("Combination attributes do not match the expected values.")
	}
}

func TestNewCombination_InvalidMovieIDs(t *testing.T) {
	votationID := "votation123"
	firstMovieID := ""
	secondMovieID := "movie2"
	chosenMovieID := "movie1"

	_, err := NewCombination(votationID, firstMovieID, secondMovieID, chosenMovieID)

	if err == nil {
		t.Error("Expected an error for invalid MovieIDs, but got nil.")
	}
}
