package valueobject

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/util"
)

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func NewAddress(city, state, country string) (*Address, []util.ProblemDetails) {
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

func ValidateAddress(city, state, country string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if city == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "Cidade inválida",
			Status:   http.StatusBadRequest,
			Detail:   "A cidade não pode estar vazia.",
			Instance: util.RFC400,
		})
	}

	if state == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "Estado inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O estado não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	if country == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "País inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O país não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	if !isCountryValid(country) {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "País inválido",
			Status:   http.StatusBadRequest,
			Detail:   "Por favor, forneça um país válido.",
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
