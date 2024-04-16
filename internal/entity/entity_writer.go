package entity

import (
	"net/http"

	"youchoose/internal/util"

	valueobject "youchoose/internal/value_object"
)

type Writer struct {
	SharedEntity
	Name        string                   `json:"name"`
	BirthDate   *valueobject.BirthDate   `json:"birth_date"`
	Nationality *valueobject.Nationality `json:"nationality"`
	ImageID     string                   `json:"image_id"`
}

func NewWriter(name string, birthDate *valueobject.BirthDate, nationality *valueobject.Nationality, imageID string) (*Writer, []util.ProblemDetails) {
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

func ValidateWriter(name string, birthDate *valueobject.BirthDate, nationality *valueobject.Nationality, imageID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidName,
			Status:   http.StatusBadRequest,
			Detail:   util.WriterErrorDetailEmptyName,
			Instance: util.RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidName,
			Status:   http.StatusBadRequest,
			Detail:   util.WriterErrorDetailMaxLengthName,
			Instance: util.RFC400,
		})
	}

	if birthDate == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidBirthDate,
			Status:   http.StatusBadRequest,
			Detail:   util.WriterErrorDetailNotNullBirthDate,
			Instance: util.RFC400,
		})
	}

	if nationality == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidNationality,
			Status:   http.StatusBadRequest,
			Detail:   util.WriterErrorDetailNotNullNationality,
			Instance: util.RFC400,
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidImageID,
			Status:   http.StatusBadRequest,
			Detail:   util.WriterErrorDetailEmptyImageID,
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
