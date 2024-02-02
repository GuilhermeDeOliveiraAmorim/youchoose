package internal

import (
	"testing"
)

func TestNewAddress_ValidAddress_ReturnsAddress(t *testing.T) {
	city := "Rio de Janeiro"
	state := "RJ"
	country := "Brazil"
	
	address, err := NewAddress(city, state, country)
	
	if err != nil {
		t.Errorf("Expected no validation errors, but got errors: %v", err)
	}

	if address == nil {
		t.Error("Expected non-nil address, but got nil")
		return
	}
	
	if address.City != city || address.State != state || address.Country != country {
		t.Error("Address attributes do not match the expected values")
	}
}

func TestNewAddress_InvalidCity_ReturnsValidationError(t *testing.T) {
	invalidCity := ""
	
	address, err := NewAddress(invalidCity, "SP", "Brazil")
	
	if err == nil {
		t.Error("Expected validation errors, but got nil")
	}

	if address != nil {
		t.Error("Expected nil address, but got an address instance")
	}

	expectedErrorMsg := "A cidade não pode estar vazia."
	if len(err) != 1 || err[0].Detail != expectedErrorMsg {
		t.Errorf("Expected validation error with detail '%s', but got: %v", expectedErrorMsg, err)
	}
}

func TestNewAddress_InvalidState_ReturnsValidationError(t *testing.T) {
	invalidState := ""
	
	address, err := NewAddress("São Paulo", invalidState, "Brazil")
	
	if err == nil {
		t.Error("Expected validation errors, but got nil")
	}

	if address != nil {
		t.Error("Expected nil address, but got an address instance")
	}

	expectedErrorMsg := "O estado não pode estar vazio."
	if len(err) != 1 || err[0].Detail != expectedErrorMsg {
		t.Errorf("Expected validation error with detail '%s', but got: %v", expectedErrorMsg, err)
	}
}

func TestNewAddress_InvalidCountry_ReturnsValidationError(t *testing.T) {
	invalidCountry := "InvalidCountry"
	
	address, err := NewAddress("Paris", "Île-de-France", invalidCountry)
	
	if err == nil {
		t.Error("Expected validation errors, but got nil")
	}

	if address != nil {
		t.Error("Expected nil address, but got an address instance")
	}

	expectedErrorMsg := "Por favor, forneça um país válido."
	if len(err) != 1 || err[0].Detail != expectedErrorMsg {
		t.Errorf("Expected validation error with detail '%s', but got: %v", expectedErrorMsg, err)
	}
}
