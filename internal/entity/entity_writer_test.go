package entity

import (
	"testing"

	valueobject "youchoose/internal/value_object"
)

func TestNewWriter(t *testing.T) {
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &valueobject.Nationality{CountryName: "Country", Flag: "Flag"}
	writer, err := NewWriter("WriterName", birthDate, nationality, "ImageID")

	if err != nil {
		t.Errorf("Erro inesperado ao criar escritor válido: %v", err)
	}

	if writer == nil {
		t.Error("O escritor não deveria ser nulo para um escritor válido")
	}

	invalidWriter, err := NewWriter("", birthDate, nationality, "ImageID")

	if err == nil {
		t.Error("Esperava-se um erro ao criar um escritor inválido (sem nome)")
	}

	if invalidWriter != nil {
		t.Error("O escritor deveria ser nulo para um escritor inválido")
	}
}

func TestValidateWriter(t *testing.T) {
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &valueobject.Nationality{CountryName: "Country", Flag: "Flag"}
	validationErrors := ValidateWriter("WriterName", birthDate, nationality, "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar escritor válido: %v", validationErrors)
	}

	validationErrors = ValidateWriter("", birthDate, nationality, "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um escritor inválido (sem nome)")
	}
}

func TestValidateWriter_NameLength(t *testing.T) {
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 1990}
	nationality := &valueobject.Nationality{CountryName: "Country", Flag: "Flag"}
	validationErrors := ValidateWriter("WriterWithValidName", birthDate, nationality, "ImageID")

	if len(validationErrors) > 0 {
		t.Errorf("Erro inesperado ao validar escritor com nome dentro do limite: %v", validationErrors)
	}

	longName := "WriterNameWithMoreThan100CharactersForTestingValidation" +
		"WriterNameWithMoreThan100CharactersForTestingValidation" +
		"WriterNameWithMoreThan100CharactersForTestingValidation"
	validationErrors = ValidateWriter(longName, birthDate, nationality, "ImageID")

	if len(validationErrors) == 0 {
		t.Error("Esperava-se erros ao validar um escritor com nome ultrapassando o limite")
	}
}
