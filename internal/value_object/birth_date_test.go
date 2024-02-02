package valueobject

import (
	"net/http"
	"testing"
	"time"
)

func TestNewBirthDate_ValidDate(t *testing.T) {
	day, month, year := 31, 12, 2000

	birthDate, err := NewBirthDate(day, month, year)

	if err != nil {
		t.Fatalf("Erro inesperado ao criar uma data válida: %s", err.Detail)
	}

	if birthDate == nil {
		t.Fatal("BirthDate não deveria ser nulo para uma data válida.")
	}

	expectedDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	if birthDate.Day != day || birthDate.Month != month || birthDate.Year != year {
		t.Errorf("Data de nascimento inválida retornada: %d/%d/%d", birthDate.Day, birthDate.Month, birthDate.Year)
	}

	if birthDate.Day != expectedDate.Day() || birthDate.Month != int(expectedDate.Month()) || birthDate.Year != expectedDate.Year() {
		t.Errorf("Data de nascimento retornada não corresponde à data esperada.")
	}
}

func TestNewBirthDate_InvalidDate(t *testing.T) {
	// Data inválida (dia 32)
	day, month, year := 32, 12, 2000

	birthDate, err := NewBirthDate(day, month, year)

	if err == nil {
		t.Error("Esperava um erro, mas obteve uma data válida.")
		return
	}

	if birthDate != nil {
		t.Error("BirthDate deveria ser nulo para uma data inválida.")
		return
	}

	expectedErrorMessage := "Por favor, forneça uma data de nascimento válida."
	if err.Detail != expectedErrorMessage {
		t.Errorf("Mensagem de erro esperada: %s, mas obteve: %s", expectedErrorMessage, err.Detail)
	}

	if err.Status != http.StatusBadRequest {
		t.Errorf("Status HTTP esperado: %d, mas obteve: %d", http.StatusBadRequest, err.Status)
	}
}