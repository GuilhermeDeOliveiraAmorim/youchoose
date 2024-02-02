package internal

import (
	"testing"
)

func TestNewActor(t *testing.T) {
	birthDate := &BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &Nationality{CountryName: "Country", Flag: "Flag"}
	actor, err := NewActor("ActorName", birthDate, nationality, "ImageID")

	if err != nil {
		t.Errorf("Erro inesperado ao criar ator válido: %v", err)
	}

	if actor == nil {
		t.Error("O ator não deveria ser nulo para um ator válido")
	}
	
	invalidActor, err := NewActor("", birthDate, nationality, "ImageID")

	if err == nil {
		t.Error("Esperava-se um erro ao criar um ator inválido (sem nome)")
	}

	if invalidActor != nil {
		t.Error("O ator deveria ser nulo para um ator inválido")
	}
}

func TestValidateActor(t *testing.T) {
	birthDate := &BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &Nationality{CountryName: "Country", Flag: "Flag"}
	validationErrors := ValidateActor("ActorName", birthDate, nationality, "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar ator válido: %v", validationErrors)
	}
	
	validationErrors = ValidateActor("", birthDate, nationality, "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um ator inválido (sem nome)")
	}
}

func TestValidateActor_NameLength(t *testing.T) {
	birthDate := &BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &Nationality{CountryName: "Country", Flag: "Flag"}
	validationErrors := ValidateActor("ActorWithValidName", birthDate, nationality, "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar ator com nome dentro do limite: %v", validationErrors)
	}
	
	longName := "ActorNameWithMoreThan100CharactersForTestingValidation" +
		"ActorNameWithMoreThan100CharactersForTestingValidation" +
		"ActorNameWithMoreThan100CharactersForTestingValidation"
	validationErrors = ValidateActor(longName, birthDate, nationality, "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um ator com nome ultrapassando o limite")
	}
}