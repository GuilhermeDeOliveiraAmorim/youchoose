package valueobject

import (
	"net/http"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/util"
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
			Type:     "ValidationError",
			Title:    "Data de nascimento inválida",
			Status:   http.StatusBadRequest,
			Detail:   "Por favor, forneça uma data de nascimento válida.",
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
