package internal

import (
	"testing"
)

func TestNewNationality_ValidCountry(t *testing.T) {
	countryName := "Brasil"
	flag := "🇧🇷"

	nationality, err := NewNationality(countryName, flag)
	if err != nil {
		t.Errorf("Esperava sucesso, mas obteve erro: %v", err)
	}

	if nationality.CountryName != countryName {
		t.Errorf("Esperava CountryName '%s', mas obteve '%s'", countryName, nationality.CountryName)
	}

	if nationality.Flag != flag {
		t.Errorf("Esperava Flag '%s', mas obteve '%s'", flag, nationality.Flag)
	}
}

func TestNewNationality_InvalidCountry(t *testing.T) {
	countryName := "PaísInválido"
	flag := "🏳️"

	nationality, err := NewNationality(countryName, flag)
	if err == nil {
		t.Error("Esperava erro, mas obteve sucesso")
	}

	if nationality != nil {
		t.Error("Nationality deveria ser nulo para um país inválido.")
		return
	}

	expectedErrorMessage := "Por favor, forneça um país válido."
	if err[0].Detail != expectedErrorMessage {
		t.Errorf("Mensagem de erro esperada: %s, mas obteve: %s", expectedErrorMessage, err[0].Detail)
	}
}

func TestNewNationality_InvalidCountry_NilInstance(t *testing.T) {
	countryName := "PaísInválido"
	flag := "🏳️"

	nationality, err := NewNationality(countryName, flag)
	if nationality != nil {
		t.Error("Esperava instância nula, mas obteve uma instância válida")
	}

	expectedErrorMessage := "Por favor, forneça um país válido."
	if err == nil || err[0].Detail != expectedErrorMessage {
		t.Errorf("Mensagem de erro esperada: %s, mas obteve: %v", expectedErrorMessage, err)
	}
}
