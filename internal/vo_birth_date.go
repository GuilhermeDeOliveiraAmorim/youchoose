package internal

import (
	"net/http"
	"time"
)

type BirthDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

func NewBirthDate(day, month, year int) (*BirthDate, []ProblemDetails) {
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

func ValidateDate(day, month, year int) []ProblemDetails {
	var validationErrors []ProblemDetails

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	if date.Day() != day || int(date.Month()) != month || date.Year() != year {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:     "ValidationError",
			Title:    "Data de nascimento inválida",
			Status:   http.StatusBadRequest,
			Detail:   "Por favor, forneça uma data de nascimento válida.",
			Instance: RFC400,
		})
	}

	return validationErrors
}
