package valueobject

import (
	"net/http"
	"time"

	"youchoose/internal/util"
)

type BirthDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

func NewBirthDate(day, month, year int) (*BirthDate, []util.ProblemDetails) {
	validationErrors := ValidateDate(day, month, year)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &BirthDate{
		Day:   day,
		Month: month,
		Year:  year,
	}, nil
}

func ValidateDate(day, month, year int) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	if date.Day() != day || int(date.Month()) != month || date.Year() != year {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidBirthDate,
			Status:   http.StatusBadRequest,
			Detail:   util.BirthDateErrorDetailInvalidDate,
			Instance: util.RFC400,
		})
	}

	return validationErrors
}

func (bd *BirthDate) Equals(other *BirthDate) bool {
	return bd.Day == other.Day && bd.Month == other.Month && bd.Year == other.Year
}
