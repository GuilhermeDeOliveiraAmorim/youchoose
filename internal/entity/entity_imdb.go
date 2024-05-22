package entity

import (
	"net/http"

	"youchoose/internal/util"
)

type IMDB struct {
	SharedEntity
	IMDBID string `json:"IMDBID"`
}

func NewIMDB(IMDBID string) (*IMDB, []util.ProblemDetails) {
	validationErrors := ValidateIMDB(IMDBID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &IMDB{
		SharedEntity: *NewSharedEntity(),
		IMDBID:       IMDBID,
	}, nil
}

func ValidateIMDB(IMDBID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if IMDBID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidID,
			Status:   http.StatusBadRequest,
			Detail:   util.IMDBErrorDetailEmptyID,
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
