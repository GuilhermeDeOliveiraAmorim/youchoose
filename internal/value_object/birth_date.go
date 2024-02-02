package valueobject

import (
	"net/http"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/youchoose/internal"
)

type BirthDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

func NewBirthDate(day, month, year int) (*BirthDate, *internal.ProblemDetails) {
	if err := validateDate(day, month, year); err != nil {
		return nil, err
	}

	return &BirthDate{
		Day:   day,
		Month: month,
		Year:  year,
	}, nil
}

func validateDate(day, month, year int) *internal.ProblemDetails {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	
	if date.Day() != day || int(date.Month()) != month || date.Year() != year {
		return &internal.ProblemDetails{
			Type:   "ValidationError",
			Title:  "Data de nascimento inválida",
			Status: http.StatusBadRequest,
			Detail: "Por favor, forneça uma data de nascimento válida.",
		}
	}

	return nil
}