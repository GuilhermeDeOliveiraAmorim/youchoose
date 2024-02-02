package internal

import (
	"net/http"
)

type Image struct {
	Name     string `json:"name"`
	Type    string `json:"type"`
	Size    int    `json:"size"`
	SharedEntity
}

func NewImage(name, imageType string, size int) (*Image, []ProblemDetails) {
	validationErrors := ValidateImage(name, size)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	image := &Image{
		Name:     name,
		Type:    imageType,
		Size:    size,
		SharedEntity: *NewSharedEntity(),
	}

	return image, nil
}


func ValidateImage(name string, size int) []ProblemDetails {
	var validationErrors []ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nome de imagem inválido",
			Status: http.StatusBadRequest,
			Detail: "A imagem deve ter um nome válido.",
		})
	}

	if size <= 0 || size > 100000 {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Tamanho de imagem inválido",
			Status: http.StatusBadRequest,
			Detail: "O tamanho da imagem deve estar entre 1 e 100000 bytes.",
		})
	}

	return validationErrors
}