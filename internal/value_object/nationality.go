package valueobject

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/youchoose/internal"
)

type Nationality struct {
	CountryName string `json:"country_name"`
	Flag        string `json:"flag"`
}

func NewNationality(countryName, flag string) (*Nationality, *internal.ProblemDetails) {
	if !isCountryValid(countryName) {
		return nil, &internal.ProblemDetails{
			Type:   "ValidationError",
			Title:  "País inválido",
			Status: http.StatusBadRequest,
			Detail: "Por favor, forneça um país válido.",
		}
	}

	return &Nationality{
		CountryName: countryName,
		Flag:        flag,
	}, nil
}