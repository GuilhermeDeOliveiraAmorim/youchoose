package entity

import (
	"net/http"
	"youchoose/internal/util"
)

type Votation struct {
	SharedEntity
	ChooserID     string `json:"chooser_id"`
	ListID        string `json:"list_id"`
	FirstMovieID  string `json:"first_movie_id"`
	SecondMovieID string `json:"second_movie_id"`
	ChosenMovieID string `json:"chosen_movie_id"`
}

func NewVotation(chooserID, listID, firstMovieID, secondMovieID, chosenMovieID string) (*Votation, []util.ProblemDetails) {
	validationErrors := ValidateVotation(chooserID, listID, firstMovieID, secondMovieID, chosenMovieID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Votation{
		SharedEntity:  *NewSharedEntity(),
		ChooserID:     chooserID,
		ListID:        listID,
		FirstMovieID:  firstMovieID,
		SecondMovieID: secondMovieID,
		ChosenMovieID: chosenMovieID,
	}, nil
}

func ValidateVotation(chooserID, listID, firstMovieID, secondMovieID, chosenMovieID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if chooserID == "" || listID == "" || firstMovieID == "" || secondMovieID == "" || chosenMovieID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidIDs,
			Status:   http.StatusBadRequest,
			Detail:   util.VoteErrorDetailEmptyIDs,
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
