package entity

import (
	"net/http"

	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"
)

type Actor struct {
	SharedEntity
	Name        string                   `json:"name"`
	BirthDate   *valueobject.BirthDate   `json:"birth_date"`
	Nationality *valueobject.Nationality `json:"nationality"`
	ImageID     string                   `json:"image_id"`
}

func NewActor(name string, birthDate *valueobject.BirthDate, nationality *valueobject.Nationality, imageID string) (*Actor, []util.ProblemDetails) {
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

func ValidateActor(name string, birthDate *valueobject.BirthDate, nationality *valueobject.Nationality, imageID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Nome do ator inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do ator não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Nome do ator inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do ator não pode ter mais do que 100 caracteres.",
			Instance: util.RFC400,
		})
	}

	if birthDate == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Data de nascimento do ator inválida",
			Status:   http.StatusBadRequest,
			Detail:   "A data de nascimento do ator não pode ser nula.",
			Instance: util.RFC400,
		})
	}

	if nationality == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Nacionalidade do ator inválida",
			Status:   http.StatusBadRequest,
			Detail:   "A nacionalidade do ator não pode ser nula.",
			Instance: util.RFC400,
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "ID de imagem do ator inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O ID de imagem do ator não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
