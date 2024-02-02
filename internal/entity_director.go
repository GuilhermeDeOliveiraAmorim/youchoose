package internal

import (
	"net/http"
)

type Director struct {
	SharedEntity
	Name        string      `json:"name"`
	BirthDate   *BirthDate  `json:"birth_date"`
	Nationality *Nationality `json:"nationality"`
	ImageID     string      `json:"image_id"`
}

func NewDirector(name string, birthDate *BirthDate, nationality *Nationality, imageID string) (*Director, []ProblemDetails) {
	validationErrors := ValidateDirector(name, birthDate, nationality, imageID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	Director := &Director{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		BirthDate:    birthDate,
		Nationality:  nationality,
		ImageID:      imageID,
	}

	return Director, nil
}

func ValidateDirector(name string, birthDate *BirthDate, nationality *Nationality, imageID string) []ProblemDetails {
	var validationErrors []ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nome do diretor inválido",
			Status: http.StatusBadRequest,
			Detail: "O nome do diretor não pode estar vazio.",
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nome do diretor inválido",
			Status: http.StatusBadRequest,
			Detail: "O nome do diretor não pode ter mais do que 100 caracteres.",
		})
	}

	if birthDate == nil {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Data de nascimento do diretor inválida",
			Status: http.StatusBadRequest,
			Detail: "A data de nascimento do diretor não pode ser nula.",
		})
	}

	if nationality == nil {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nacionalidade do diretor inválida",
			Status: http.StatusBadRequest,
			Detail: "A nacionalidade do diretor não pode ser nula.",
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "ID de imagem do diretor inválido",
			Status: http.StatusBadRequest,
			Detail: "O ID de imagem do diretor não pode estar vazio.",
		})
	}

	return validationErrors
}
