package internal

import (
	"net/http"
)

type Writer struct {
	SharedEntity
	Name        string      `json:"name"`
	BirthDate   *BirthDate  `json:"birth_date"`
	Nationality *Nationality `json:"nationality"`
	ImageID     string      `json:"image_id"`
}

func NewWriter(name string, birthDate *BirthDate, nationality *Nationality, imageID string) (*Writer, []ProblemDetails) {
	validationErrors := ValidateWriter(name, birthDate, nationality, imageID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	actor := &Writer{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		BirthDate:    birthDate,
		Nationality:  nationality,
		ImageID:      imageID,
	}

	return actor, nil
}

func ValidateWriter(name string, birthDate *BirthDate, nationality *Nationality, imageID string) []ProblemDetails {
	var validationErrors []ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nome do(a) escritor(a) inválido",
			Status: http.StatusBadRequest,
			Detail: "O nome do(a) escritor(a) não pode estar vazio.",
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nome do(a) escritor(a) inválido",
			Status: http.StatusBadRequest,
			Detail: "O nome do(a) escritor(a) não pode ter mais do que 100 caracteres.",
		})
	}

	if birthDate == nil {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Data de nascimento do(a) escritor(a) inválida",
			Status: http.StatusBadRequest,
			Detail: "A data de nascimento do(a) escritor(a) não pode ser nula.",
		})
	}

	if nationality == nil {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Nacionalidade do(a) escritor(a) inválida",
			Status: http.StatusBadRequest,
			Detail: "A nacionalidade do(a) escritor(a) não pode ser nula.",
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "ID de imagem do(a) escritor(a) inválido",
			Status: http.StatusBadRequest,
			Detail: "O ID de imagem do(a) escritor(a) não pode estar vazio.",
		})
	}

	return validationErrors
}
