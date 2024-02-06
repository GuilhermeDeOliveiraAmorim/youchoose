package internal

import (
	"net/http"
)

type Nationality struct {
	CountryName string `json:"country_name"`
	Flag        string `json:"flag"`
}

func NewNationality(countryName, flag string) (*Nationality, []ProblemDetails) {
	validationErrors := ValidateNationality(countryName)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Nationality{
		CountryName: countryName,
		Flag:        flag,
	}, nil
}

func ValidateNationality(country string) []ProblemDetails {
	var validationErrors []ProblemDetails

	for _, c := range countries {
		if c.Name == country {
			return validationErrors
		}
	}

	validationErrors = append(validationErrors, ProblemDetails{
		Type:     "ValidationError",
		Title:    "País inválido",
		Status:   http.StatusBadRequest,
		Detail:   "Por favor, forneça um país válido.",
		Instance: RFC400,
	})

	return validationErrors
}
