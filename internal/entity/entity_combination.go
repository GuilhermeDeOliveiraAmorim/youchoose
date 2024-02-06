package entity

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/util"
)

type Combination struct {
	SharedEntity
	VotationID    string `json:"votation_id"`
	FirstMovieID  string `json:"first_movie_id"`
	SecondMovieID string `json:"second_movie_id"`
	ChosenMovieID string `json:"chosen_movie_id"`
}

func NewCombination(votationID, firstMovieID, secondMovieID, chosenMovieID string) (*Combination, []util.ProblemDetails) {
	validationErrors := ValidateCombination(votationID, firstMovieID, secondMovieID, chosenMovieID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	combination := &Combination{
		SharedEntity:  *NewSharedEntity(),
		VotationID:    votationID,
		FirstMovieID:  firstMovieID,
		SecondMovieID: secondMovieID,
		ChosenMovieID: chosenMovieID,
	}

	return combination, nil
}

func ValidateCombination(votationID, firstMovieID, secondMovieID, chosenMovieID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if firstMovieID == "" || secondMovieID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "IDs dos filmes na combinação inválidos",
			Status:   http.StatusBadRequest,
			Detail:   "Os IDs dos filmes na combinação não podem estar vazios.",
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
