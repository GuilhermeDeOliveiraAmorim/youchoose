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
	image := &Image{
		Name:     name,
		Type:    imageType,
		Size:    size,
		SharedEntity: *NewSharedEntity(),
	}

	validationErrors := ValidateImage(image)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return image, nil
}


func ValidateImage(image *Image) []ProblemDetails {
	var validationErrors []ProblemDetails

	if image.Name == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nome de imagem inválido",
			Status: http.StatusBadRequest,
			Detail: "A imagem deve ter um nome válido.",
		})
	}

	if image.Size <= 0 || image.Size > 100000 {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Tamanho de imagem inválido",
			Status: http.StatusBadRequest,
			Detail: "O tamanho da imagem deve estar entre 1 e 100000 bytes.",
		})
	}

	return validationErrors
}