package entity

import (
	"net/http"

	"youchoose/internal/util"
)

type Image struct {
	SharedEntity
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}

func NewImage(name, imageType string, size int64) (*Image, []util.ProblemDetails) {
	validationErrors := ValidateImage(name, size)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Image{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		Type:         imageType,
		Size:         size,
	}, nil
}

func ValidateImage(name string, size int64) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Nome de imagem inválido",
			Status:   http.StatusBadRequest,
			Detail:   "A imagem deve ter um nome válido.",
			Instance: util.RFC400,
		})
	}

	if size <= 0 || size > 100000 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Tamanho de imagem inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O tamanho da imagem deve estar entre 1 e 100000 bytes.",
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
