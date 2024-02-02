package internal

import (
	"testing"
)

func TestNewImageValid(t *testing.T) {
	name := "example.jpg"
	imageType := "jpeg"
	size := 50000

	image, err := NewImage(name, imageType, size)

	if err != nil {
		t.Errorf("Erro inesperado ao criar a imagem: %v", err)
	}

	if image == nil {
		t.Error("A imagem não deveria ser nula")
	}
}

func TestNewImageInvalidName(t *testing.T) {
	name := ""
	imageType := "jpeg"
	size := 50000

	image, err := NewImage(name, imageType, size)

	if err == nil {
		t.Error("Esperava um erro ao criar a imagem com nome inválido")
	}

	if image != nil {
		t.Error("A imagem deveria ser nula")
	}

	if len(err) == 0 {
		t.Error("Esperava notificações de erro, mas não encontrou nenhuma")
	}
}

func TestNewImageInvalidSize(t *testing.T) {
	name := "example.jpg"
	imageType := "jpeg"
	size := 200000

	image, err := NewImage(name, imageType, size)

	if err == nil {
		t.Error("Esperava um erro ao criar a imagem com tamanho inválido")
	}

	if image != nil {
		t.Error("A imagem deveria ser nula")
	}

	if len(err) == 0 {
		t.Error("Esperava notificações de erro, mas não encontrou nenhuma")
	}
}
