package internal

import (
	"net/http"
)

type Genre struct {
	SharedEntity
	Name        string      `json:"name"`
	ImageID     string      `json:"image_id"`
}

func NewGenre(name, imageID string) (*Genre, []ProblemDetails) {
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

func ValidateGenre(name, imageID string) []ProblemDetails {
	var validationErrors []ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nome do gênero inválido",
			Status: http.StatusBadRequest,
			Detail: "O nome do gênero não pode estar vazio.",
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nome do gênero inválido",
			Status: http.StatusBadRequest,
			Detail: "O nome do gênero não pode ter mais do que 100 caracteres.",
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "ID de imagem do gênero inválido",
			Status: http.StatusBadRequest,
			Detail: "O ID de imagem do gênero não pode estar vazio.",
		})
	}

	return validationErrors
}
