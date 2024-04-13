package entity

import (
	"net/http"

	"youchoose/internal/util"
)

type Genre struct {
	SharedEntity
	Name    string `json:"name"`
	ImageID string `json:"image_id"`
}

func NewGenre(name, imageID string) (*Genre, []util.ProblemDetails) {
	validationErrors := ValidateGenre(name, imageID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Genre{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		ImageID:      imageID,
	}, nil
}

func ValidateGenre(name, imageID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Nome do gênero inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do gênero não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Nome do gênero inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do gênero não pode ter mais do que 100 caracteres.",
			Instance: util.RFC400,
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "ID de imagem do gênero inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O ID de imagem do gênero não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
