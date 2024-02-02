package internal

import (
	"testing"
)

func TestNewDirector(t *testing.T) {
	birthDate := &BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &Nationality{CountryName: "Country", Flag: "Flag"}
	actor, err := NewDirector("DirectorName", birthDate, nationality, "ImageID")

	if err != nil {
		t.Errorf("Erro inesperado ao criar diretor válido: %v", err)
	}

	if actor == nil {
		t.Error("O diretor não deveria ser nulo para um diretor válido")
	}
	
	invalidDirector, err := NewDirector("", birthDate, nationality, "ImageID")

	if err == nil {
		t.Error("Esperava-se um erro ao criar um diretor inválido (sem nome)")
	}

	if invalidDirector != nil {
		t.Error("O diretor deveria ser nulo para um diretor inválido")
	}
}

func TestValidateDirector(t *testing.T) {
	birthDate := &BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &Nationality{CountryName: "Country", Flag: "Flag"}
	validationErrors := ValidateDirector("DirectorName", birthDate, nationality, "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar diretor válido: %v", validationErrors)
	}
	
	validationErrors = ValidateDirector("", birthDate, nationality, "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um diretor inválido (sem nome)")
	}
}

func TestValidateDirector_NameLength(t *testing.T) {
	birthDate := &BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &Nationality{CountryName: "Country", Flag: "Flag"}
	validationErrors := ValidateDirector("DirectorWithValidName", birthDate, nationality, "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar diretor com nome dentro do limite: %v", validationErrors)
	}
	
	longName := "DirectorNameWithMoreThan100CharactersForTestingValidation" +
		"DirectorNameWithMoreThan100CharactersForTestingValidation" +
		"DirectorNameWithMoreThan100CharactersForTestingValidation"
	validationErrors = ValidateDirector(longName, birthDate, nationality, "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um diretor com nome ultrapassando o limite")
	}
}