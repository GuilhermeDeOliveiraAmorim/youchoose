package entity

import (
	"testing"

	valueobject "youchoose/internal/value_object"
)

func TestNewList(t *testing.T) {
	list, err := NewList("Minha Lista", "Descrição da Lista", "Minha Lista", "Descrição da Lista", "chooser123")
	if err != nil {
		t.Errorf("Erro ao criar uma nova lista válida: %v", err)
	}

	if list.Title != "Minha Lista" || list.Description != "Descrição da Lista" || list.ChooserID != "chooser123" {
		t.Errorf("A lista não foi criada corretamente. Detalhes da lista: %v", list)
	}

	_, err = NewList("", "Descrição da Lista", "profile123", "cover123", "chooser123")
	if err == nil {
		t.Error("Criou uma lista com título vazio, mas deveria ter retornado um erro.")
	}
}

func TestList_AddMovies(t *testing.T) {
	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	movie, _ := NewMovie("Inception", *nationality, 2010, "image123")

	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	list.AddMovies([]Movie{*movie})

	if len(list.Movies) != 1 || list.Movies[0].Title != "Inception" {
		t.Errorf("Erro ao adicionar filme à lista. Detalhes da lista: %v", list)
	}
}

func TestList_RemoveMovies(t *testing.T) {
	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	movie1, _ := NewMovie("Inception", *nationality, 2010, "image123")

	movie2, _ := NewMovie("Interstellar", *nationality, 2014, "image456")

	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	list.AddMovies([]Movie{*movie1, *movie2})

	list.RemoveMovies([]Movie{*movie1})

	if len(list.Movies) != 1 || list.Movies[0].Title != "Interstellar" {
		t.Errorf("Erro ao remover filme da lista. Detalhes da lista: %v", list)
	}
}

func TestList_GetAvailableMoviesCombinations(t *testing.T) {
	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	movie1, _ := NewMovie("Inception", *nationality, 2010, "image123")

	movie2, _ := NewMovie("Interstellar", *nationality, 2014, "image456")

	movie3, _ := NewMovie("The Dark Knight", *nationality, 2008, "image789")

	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	list.AddMovies([]Movie{*movie1, *movie2, *movie3})

	combinations := list.GetAvailableMoviesCombinations()

	if len(combinations) != 3 {
		t.Errorf("Erro ao obter combinações de filmes. Número incorreto de combinações.")
	}
}

func TestList_ValidateList(t *testing.T) {
	validationErrors := ValidateList("Minha Lista", "Descrição da Lista", "chooser123")
	if len(validationErrors) > 0 {
		t.Errorf("Erro ao validar uma lista válida. Erros: %v", validationErrors)
	}

	validationErrors = ValidateList("", "Descrição da Lista", "chooser123")
	if len(validationErrors) == 0 {
		t.Error("Validou uma lista com título vazio, mas deveria ter retornado erro.")
	}
}

func TestList_IncrementVotes(t *testing.T) {
	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	list.IncrementVotes()

	if list.Votes != 1 {
		t.Errorf("Erro ao incrementar votos. Número de votos incorreto: %d", list.Votes)
	}
}

func TestList_ChangeProfileImageID(t *testing.T) {
	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	list.ChangeProfileImageID("new_profile_image")

	if list.ProfileImageID != "new_profile_image" {
		t.Errorf("Erro ao alterar imagem de perfil. Valor atual: %s", list.ProfileImageID)
	}
}

func TestList_ChangeCoverImageID(t *testing.T) {
	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	list.ChangeCoverImageID("new_cover_image")

	if list.CoverImageID != "new_cover_image" {
		t.Errorf("Erro ao alterar imagem de capa. Valor atual: %s", list.CoverImageID)
	}
}

func TestList_ChangeTitle(t *testing.T) {
	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	list.ChangeTitle("Nova Lista")

	if list.Title != "Nova Lista" {
		t.Errorf("Erro ao alterar título. Valor atual: %s", list.Title)
	}
}
