package entity

import (
	"net/http"

	"youchoose/internal/util"
)

type Genre struct {
	SharedEntity
	Name string `json:"name"`
}

func NewGenre(name string) (*Genre, []util.ProblemDetails) {
	validationErrors := ValidateGenre(name)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Genre{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
	}, nil
}

func ValidateGenre(name string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidName,
			Status:   http.StatusBadRequest,
			Detail:   util.GenreErrorDetailEmptyName,
			Instance: util.RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidName,
			Status:   http.StatusBadRequest,
			Detail:   util.GenreErrorDetailMaxLengthName,
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
