package internal

import (
	"net/http"
)

type Actor struct {
	SharedEntity
	Name        string       `json:"name"`
	BirthDate   *BirthDate   `json:"birth_date"`
	Nationality *Nationality `json:"nationality"`
	ImageID     string       `json:"image_id"`
}

func NewActor(name string, birthDate *BirthDate, nationality *Nationality, imageID string) (*Actor, []ProblemDetails) {
	validationErrors := ValidateActor(name, birthDate, nationality, imageID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	actor := &Actor{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		BirthDate:    birthDate,
		Nationality:  nationality,
		ImageID:      imageID,
	}

	return actor, nil
}

func ValidateActor(name string, birthDate *BirthDate, nationality *Nationality, imageID string) []ProblemDetails {
	var validationErrors []ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:     "ValidationError",
			Title:    "Nome do ator inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do ator não pode estar vazio.",
			Instance: RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:     "ValidationError",
			Title:    "Nome do ator inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do ator não pode ter mais do que 100 caracteres.",
			Instance: RFC400,
		})
	}

	if birthDate == nil {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:     "ValidationError",
			Title:    "Data de nascimento do ator inválida",
			Status:   http.StatusBadRequest,
			Detail:   "A data de nascimento do ator não pode ser nula.",
			Instance: RFC400,
		})
	}

	if nationality == nil {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:     "ValidationError",
			Title:    "Nacionalidade do ator inválida",
			Status:   http.StatusBadRequest,
			Detail:   "A nacionalidade do ator não pode ser nula.",
			Instance: RFC400,
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:     "ValidationError",
			Title:    "ID de imagem do ator inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O ID de imagem do ator não pode estar vazio.",
			Instance: RFC400,
		})
	}

	return validationErrors
}
