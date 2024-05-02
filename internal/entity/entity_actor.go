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

	return &Actor{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		BirthDate:    birthDate,
		Nationality:  nationality,
		ImageID:      imageID,
	}, nil
}

func ValidateActor(name string, birthDate *valueobject.BirthDate, nationality *valueobject.Nationality, imageID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidName,
			Status:   http.StatusBadRequest,
			Detail:   util.ActorErrorDetailEmptyName,
			Instance: util.RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidName,
			Status:   http.StatusBadRequest,
			Detail:   util.ActorErrorDetailMaxLengthName,
			Instance: util.RFC400,
		})
	}

	if birthDate == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidBirthDate,
			Status:   http.StatusBadRequest,
			Detail:   util.ActorErrorDetailNotNullBirthDate,
			Instance: util.RFC400,
		})
	}

	if nationality == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidNationality,
			Status:   http.StatusBadRequest,
			Detail:   util.ActorErrorDetailNotNullNationality,
			Instance: util.RFC400,
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidImageID,
			Status:   http.StatusBadRequest,
			Detail:   util.ActorErrorDetailEmptyImageID,
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
