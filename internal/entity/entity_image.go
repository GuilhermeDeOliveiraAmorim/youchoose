package entity

import (
	"net/http"

	"youchoose/internal/util"
)

type Image struct {
	SharedEntity
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewImage(name, imageType string) (*Image, []util.ProblemDetails) {
	validationErrors := ValidateImage(name)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Image{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		Type:         imageType,
	}, nil
}

func ValidateImage(name string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidName,
			Status:   http.StatusBadRequest,
			Detail:   util.ImageErrorDetailEmptyName,
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
