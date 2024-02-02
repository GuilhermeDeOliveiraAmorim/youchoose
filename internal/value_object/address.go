package valueobject

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/youchoose/internal"
)

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func NewAddress(city, state, country string) (*Address, []internal.ProblemDetails) {
	validationErrors := ValidateAddress(city, state, country)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	address := &Address{
		City:    city,
		State:   state,
		Country: country,
	}

	return address, nil
}

func ValidateAddress(city, state, country string) []internal.ProblemDetails {
	var validationErrors []internal.ProblemDetails

	if city == "" {
		validationErrors = append(validationErrors, internal.ProblemDetails{
			Type:   "ValidationError",
			Title:  "Cidade inválida",
			Status: http.StatusBadRequest,
			Detail: "A cidade não pode estar vazia.",
		})
	}

	if state == "" {
		validationErrors = append(validationErrors, internal.ProblemDetails{
			Type:   "ValidationError",
			Title:  "Estado inválido",
			Status: http.StatusBadRequest,
			Detail: "O estado não pode estar vazio.",
		})
	}

	if country == "" {
		validationErrors = append(validationErrors, internal.ProblemDetails{
			Type:   "ValidationError",
			Title:  "País inválido",
			Status: http.StatusBadRequest,
			Detail: "O país não pode estar vazio.",
		})
	}

	if !isCountryValid(country) {
		validationErrors = append(validationErrors, internal.ProblemDetails{
			Type:   "ValidationError",
			Title:  "País inválido",
			Status: http.StatusBadRequest,
			Detail: "Por favor, forneça um país válido.",
		})
	}

	return validationErrors
}