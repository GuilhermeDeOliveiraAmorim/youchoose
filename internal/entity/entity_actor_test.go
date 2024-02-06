package entity

import (
	"testing"

	valueobject "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/value_object"
)

func TestNewActor(t *testing.T) {
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &valueobject.Nationality{CountryName: "Country", Flag: "Flag"}
	actor, err := NewActor("ActorName", birthDate, nationality, "ImageID")

	if err != nil {
		t.Errorf("Erro inesperado ao criar ator(atriz) válido: %v", err)
	}

	if actor == nil {
		t.Error("O ator(atriz) não deveria ser nulo para um ator(atriz) válido")
	}

	invalidActor, err := NewActor("", birthDate, nationality, "ImageID")

	if err == nil {
		t.Error("Esperava-se um erro ao criar um ator(atriz) inválido (sem nome)")
	}

	if invalidActor != nil {
		t.Error("O ator(atriz) deveria ser nulo para um ator(atriz) inválido")
	}
}

func TestValidateActor(t *testing.T) {
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &valueobject.Nationality{CountryName: "Country", Flag: "Flag"}
	validationErrors := ValidateActor("ActorName", birthDate, nationality, "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar ator(atriz) válido: %v", validationErrors)
	}

	validationErrors = ValidateActor("", birthDate, nationality, "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um ator(atriz) inválido (sem nome)")
	}
}

func TestValidateActor_NameLength(t *testing.T) {
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &valueobject.Nationality{CountryName: "Country", Flag: "Flag"}
	validationErrors := ValidateActor("ActorWithValidName", birthDate, nationality, "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar ator(atriz) com nome dentro do limite: %v", validationErrors)
	}

	longName := "ActorNameWithMoreThan100CharactersForTestingValidation" +
		"ActorNameWithMoreThan100CharactersForTestingValidation" +
		"ActorNameWithMoreThan100CharactersForTestingValidation"
	validationErrors = ValidateActor(longName, birthDate, nationality, "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um ator(atriz) com nome ultrapassando o limite")
	}
}
