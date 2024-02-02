package valueobject

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/youchoose/internal"
)

type Nationality struct {
	CountryName string `json:"country_name"`
	Flag        string `json:"flag"`
}

func NewNationality(countryName, flag string) (*Nationality, []internal.ProblemDetails) {
	validationErrors := ValidateNationality(countryName)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Nationality{
		CountryName: countryName,
		Flag:        flag,
	}, nil
}