package entity

import "testing"

func TestNewIMDBValid(t *testing.T) {
	image, err := NewIMDB("tt6900448")

	if err != nil {
		t.Errorf("Erro inesperado ao criar IMDB: %v", err)
	}

	if image == nil {
		t.Error("O IMDB não deveria ser nulo")
	}
}
