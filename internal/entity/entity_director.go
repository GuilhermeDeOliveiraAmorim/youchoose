package entity

import (
	"net/http"

	"youchoose/internal/util"

	valueobject "youchoose/internal/value_object"
)

type Director struct {
	SharedEntity
	Name        string                   `json:"name"`
	BirthDate   *valueobject.BirthDate   `json:"birth_date"`
	Nationality *valueobject.Nationality `json:"nationality"`
	ImageID     string                   `json:"image_id"`
}

func NewDirector(name string, birthDate *valueobject.BirthDate, nationality *valueobject.Nationality, imageID string) (*Director, []util.ProblemDetails) {
	validationErrors := ValidateDirector(name, birthDate, nationality, imageID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Director{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		BirthDate:    birthDate,
		Nationality:  nationality,
		ImageID:      imageID,
	}, nil
}

func ValidateDirector(name string, birthDate *valueobject.BirthDate, nationality *valueobject.Nationality, imageID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Nome do(a) diretor(a) inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do(a) diretor(a) não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Nome do(a) diretor(a) inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do(a) diretor(a) não pode ter mais do que 100 caracteres.",
			Instance: util.RFC400,
		})
	}

	if birthDate == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Data de nascimento do(a) diretor(a) inválida",
			Status:   http.StatusBadRequest,
			Detail:   "A data de nascimento do(a) diretor(a) não pode ser nula.",
			Instance: util.RFC400,
		})
	}

	if nationality == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Nacionalidade do(a) diretor(a) inválida",
			Status:   http.StatusBadRequest,
			Detail:   "A nacionalidade do(a) diretor(a) não pode ser nula.",
			Instance: util.RFC400,
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "ID de imagem do(a) diretor(a) inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O ID de imagem do(a) diretor(a) não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
