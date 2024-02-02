package internal

import (
	"net/http"
)

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func NewAddress(city, state, country string) (*Address, []ProblemDetails) {
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

func ValidateAddress(city, state, country string) []ProblemDetails {
	var validationErrors []ProblemDetails

	if city == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Cidade inválida",
			Status: http.StatusBadRequest,
			Detail: "A cidade não pode estar vazia.",
		})
	}

	if state == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Estado inválido",
			Status: http.StatusBadRequest,
			Detail: "O estado não pode estar vazio.",
		})
	}

	if country == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "País inválido",
			Status: http.StatusBadRequest,
			Detail: "O país não pode estar vazio.",
		})
	}

	if !isCountryValid(country) {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "País inválido",
			Status: http.StatusBadRequest,
			Detail: "Por favor, forneça um país válido.",
		})
	}

	return validationErrors
}