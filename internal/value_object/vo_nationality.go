package valueobject

import (
	"net/http"

	"youchoose/internal/util"
)

type Nationality struct {
	CountryName string `json:"country_name"`
	Flag        string `json:"flag"`
}

func NewNationality(countryName, flag string) (*Nationality, []util.ProblemDetails) {
	validationErrors := ValidateNationality(countryName)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Nationality{
		CountryName: countryName,
		Flag:        flag,
	}, nil
}

func ValidateNationality(country string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	for _, c := range countries {
		if c.Name == country {
			return validationErrors
		}
	}

	validationErrors = append(validationErrors, util.ProblemDetails{
		Type:     "Validation Error",
		Title:    "País inválido",
		Status:   http.StatusBadRequest,
		Detail:   "Por favor, forneça um país válido.",
		Instance: util.RFC400,
	})

	return validationErrors
}
