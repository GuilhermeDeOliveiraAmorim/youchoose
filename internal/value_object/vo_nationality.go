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
	validationErrors := ValidateNationality(countryName, flag)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Nationality{
		CountryName: countryName,
		Flag:        flag,
	}, nil
}

func ValidateNationality(country, flag string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails
	countries := NewCountries()

	if country == "" || flag == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    util.NationalityErrorTitleCountryOrFlagEmpty,
			Status:   http.StatusBadRequest,
			Detail:   util.NationalityErrorDetailCountryOrFlagEmpty,
			Instance: util.RFC400,
		})
	}

	for _, c := range countries {
		if c.Name == country && c.Flag == flag {
			return validationErrors
		}
	}

	validationErrors = append(validationErrors, util.ProblemDetails{
		Type:     util.TypeValidationError,
		Title:    util.NationalityErrorTitleCountryOrFlagEmpty,
		Status:   http.StatusBadRequest,
		Detail:   util.NationalityErrorDetailCountryOrFlagEmpty,
		Instance: util.RFC400,
	})

	return validationErrors
}

func (na *Nationality) Equals(other *Nationality) bool {
	return na.CountryName == other.CountryName && na.Flag == other.Flag
}
