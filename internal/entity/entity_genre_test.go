package entity

import (
	"testing"
)

func TestNewGenre(t *testing.T) {
	actor, err := NewGenre("GenreName", "ImageID")

	if err != nil {
		t.Errorf("Erro inesperado ao criar gênero válido: %v", err)
	}

	if actor == nil {
		t.Error("O gênero não deveria ser nulo para um gênero válido")
	}

	invalidGenre, err := NewGenre("", "ImageID")

	if err == nil {
		t.Error("Esperava-se um erro ao criar um gênero inválido (sem nome)")
	}

	if invalidGenre != nil {
		t.Error("O gênero deveria ser nulo para um gênero inválido")
	}
}

func TestValidateGenre(t *testing.T) {
	validationErrors := ValidateGenre("GenreName", "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar gênero válido: %v", validationErrors)
	}

	validationErrors = ValidateGenre("", "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um gênero inválido (sem nome)")
	}
}

func TestValidateGenre_NameLength(t *testing.T) {
	validationErrors := ValidateGenre("GenreWithValidName", "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar gênero com nome dentro do limite: %v", validationErrors)
	}

	longName := "GenreNameWithMoreThan100CharactersForTestingValidation" +
		"GenreNameWithMoreThan100CharactersForTestingValidation" +
		"GenreNameWithMoreThan100CharactersForTestingValidation"
	validationErrors = ValidateGenre(longName, "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um gênero com nome ultrapassando o limite")
	}
}
