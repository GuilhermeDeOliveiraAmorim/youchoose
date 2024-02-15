package entity

import (
	"testing"
)

func TestVotation_StartVotation(t *testing.T) {
	combinations := []Combination{
		{VotationID: "1", FirstMovieID: "A", SecondMovieID: "B"},
		{VotationID: "1", FirstMovieID: "C", SecondMovieID: "D"},
	}

	votation := NewVotation("chooserID", "listID", combinations)

	if votation.StartTime.IsZero() {
		t.Error("Expected non-zero start time after starting the votation.")
	}

	if votation.Status != INPROGRESS {
		t.Error("Expected votation status to be 'In progress' after starting.")
	}
}

func TestVotation_EndVotation(t *testing.T) {
	combinations := []Combination{
		{VotationID: "1", FirstMovieID: "A", SecondMovieID: "B"},
		{VotationID: "1", FirstMovieID: "C", SecondMovieID: "D"},
	}

	votation := NewVotation("chooserID", "listID", combinations)
	votation.EndVotation()

	if votation.EndTime.IsZero() {
		t.Error("Expected non-zero end time after ending the votation.")
	}

	if votation.Status != FINISHED {
		t.Error("Expected votation status to be 'Finished' after ending.")
	}
}

func TestVotation_IsVotationInProgress(t *testing.T) {
	combinations := []Combination{
		{VotationID: "1", FirstMovieID: "A", SecondMovieID: "B"},
		{VotationID: "1", FirstMovieID: "C", SecondMovieID: "D"},
	}

	votation := NewVotation("chooserID", "listID", combinations)

	if !votation.IsVotationInProgress() {
		t.Error("Expected votation to be in progress.")
	}

	votation.EndVotation()

	if votation.IsVotationInProgress() {
		t.Error("Expected votation to be finished after ending.")
	}
}

func TestVotation_Vote(t *testing.T) {
	combinations := []Combination{
		{VotationID: "1", FirstMovieID: "A", SecondMovieID: "B"},
		{VotationID: "1", FirstMovieID: "C", SecondMovieID: "D"},
	}

	votation := NewVotation("chooserID", "listID", combinations)
	chosenCombination := Combination{VotationID: "1", FirstMovieID: "A", SecondMovieID: "B", ChosenMovieID: "A"}

	votation.Vote(chosenCombination)

	if (votation.ChosenCombination.VotationID != chosenCombination.VotationID) || (votation.ChosenCombination.FirstMovieID != chosenCombination.FirstMovieID) || (votation.ChosenCombination.SecondMovieID != chosenCombination.SecondMovieID) {
		t.Error("Expected votation to have the chosen combination.")
	}
}

func TestVotation_GetAvailableCombinations(t *testing.T) {
	combinations := []Combination{
		{VotationID: "1", FirstMovieID: "A", SecondMovieID: "B"},
		{VotationID: "1", FirstMovieID: "C", SecondMovieID: "D"},
	}

	votation := NewVotation("chooserID", "listID", combinations)

	availableCombinations := votation.GetAvailableCombinations()

	if len(availableCombinations) != len(combinations) {
		t.Error("Expected all combinations to be available.")
	}
}

func TestVotation_GetVotedCombinations(t *testing.T) {
	combinations := []Combination{
		{VotationID: "1", FirstMovieID: "A", SecondMovieID: "B", ChosenMovieID: "A"},
		{VotationID: "1", FirstMovieID: "C", SecondMovieID: "D", ChosenMovieID: "D"},
		{VotationID: "1", FirstMovieID: "E", SecondMovieID: "F"},
	}

	votation := NewVotation("chooserID", "listID", combinations)

	votedCombinations := votation.GetVotedCombinations()

	if len(votedCombinations) != 2 {
		t.Error("Expected two voted combinations.")
	}
}

func TestVotation_GetUnvotedCombinations(t *testing.T) {
	combinations := []Combination{
		{VotationID: "1", FirstMovieID: "A", SecondMovieID: "B", ChosenMovieID: "A"},
		{VotationID: "1", FirstMovieID: "C", SecondMovieID: "D", ChosenMovieID: "D"},
		{VotationID: "1", FirstMovieID: "E", SecondMovieID: "F"},
	}

	votation := NewVotation("chooserID", "listID", combinations)

	unvotedCombinations := votation.GetUnvotedCombinations()

	if len(unvotedCombinations) != 1 {
		t.Error("Expected one unvoted combination.")
	}
}
