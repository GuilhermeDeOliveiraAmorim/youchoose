package valueobject

import (
	"net/http"

	"youchoose/internal/util"
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
			Type:     util.TypeValidationError,
			Title:    util.AddressErrorTitleEmptyCity,
			Status:   http.StatusBadRequest,
			Detail:   util.AddressErrorDetailEmptyCity,
			Instance: util.RFC400,
		})
	}

	if state == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.AddressErrorTitleEmptyState,
			Status:   http.StatusBadRequest,
			Detail:   util.AddressErrorDetailEmptyState,
			Instance: util.RFC400,
		})
	}

	if country == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidCountry,
			Status:   http.StatusBadRequest,
			Detail:   util.AddressErrorDetailInvalidCountry,
			Instance: util.RFC400,
		})
	}

	if !isCountryValid(country) {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.AddressErrorTitleInvalidCountry,
			Status:   http.StatusBadRequest,
			Detail:   util.AddressErrorDetailInvalidCountry,
			Instance: util.RFC400,
		})
	}

	return validationErrors
}

func (ad *Address) Equals(other *Address) bool {
	return ad.City == other.City && ad.Country == other.Country && ad.State == other.State
}
