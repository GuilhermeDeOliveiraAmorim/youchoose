package entity

import (
	"testing"

	valueobject "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/value_object"
)

func TestNewList(t *testing.T) {
	// Testando a criação de uma nova lista válida
	list, err := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")
	if err != nil {
		t.Errorf("Erro ao criar uma nova lista válida: %v", err)
	}

	// Verificando se a lista foi criada corretamente
	if list.Title != "Minha Lista" || list.Description != "Descrição da Lista" || list.ChooserID != "chooser123" {
		t.Errorf("A lista não foi criada corretamente. Detalhes da lista: %v", list)
	}

	// Testando a criação de uma nova lista inválida (título vazio)
	_, err = NewList("profile123", "cover123", "", "Descrição da Lista", "chooser123")
	if err == nil {
		t.Error("Criou uma lista com título vazio, mas deveria ter retornado um erro.")
	}
}

func TestList_AddMovies(t *testing.T) {
	// Criando instância para BirthDate
	birthDate, _ := valueobject.NewBirthDate(10, 5, 2010)

	// Criando instância para Nationality
	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	// Criando instância para ator
	actor, _ := NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")

	// Criando instância para gênero
	genre, _ := NewGenre("Ação", "image_id_genre")

	// Criando instância para diretor
	director, _ := NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")

	// Criando instância para filme
	movie, _ := NewMovie("Inception", *nationality, []Genre{*genre}, []Director{*director}, []Actor{*actor}, []Writer{}, 2010, "image123")

	// Criando instância para lista
	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	// Adicionando filme à lista
	list.AddMovies([]Movie{*movie})

	// Verificando se o filme foi adicionado corretamente
	if len(list.Movies) != 1 || list.Movies[0].Title != "Inception" {
		t.Errorf("Erro ao adicionar filme à lista. Detalhes da lista: %v", list)
	}
}

func TestList_RemoveMovies(t *testing.T) {
	// Criando instância para BirthDate
	birthDate, _ := valueobject.NewBirthDate(10, 5, 1990)

	// Criando instância para Nationality
	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	// Criando instância para ator
	actor, _ := NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")

	// Criando instância para gênero
	genre, _ := NewGenre("Ação", "image_id_genre")

	// Criando instância para diretor
	director, _ := NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")

	// Criando instância para filme
	movie1, _ := NewMovie("Inception", *nationality, []Genre{*genre}, []Director{*director}, []Actor{*actor}, []Writer{}, 2010, "image123")

	// Criando instância para filme
	movie2, _ := NewMovie("Interstellar", *nationality, []Genre{*genre}, []Director{*director}, []Actor{*actor}, []Writer{}, 2014, "image456")

	// Criando instância para lista
	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	// Adicionando filmes à lista
	list.AddMovies([]Movie{*movie1, *movie2})

	// Removendo filme da lista
	list.RemoveMovies([]Movie{*movie1})

	// Verificando se o filme foi removido corretamente
	if len(list.Movies) != 1 || list.Movies[0].Title != "Interstellar" {
		t.Errorf("Erro ao remover filme da lista. Detalhes da lista: %v", list)
	}
}

func TestList_GetAvailableMoviesCombinations(t *testing.T) {
	// Criando instância para BirthDate
	birthDate, _ := valueobject.NewBirthDate(15, 5, 1990)

	// Criando instância para Nationality
	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	// Criando instância para ator
	actor, _ := NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")

	// Criando instância para gênero
	genre, _ := NewGenre("Ação", "image_id_genre")

	// Criando instância para diretor
	director, _ := NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")

	// Criando instância para filme
	movie1, _ := NewMovie("Inception", *nationality, []Genre{*genre}, []Director{*director}, []Actor{*actor}, []Writer{}, 2010, "image123")

	// Criando instância para filme
	movie2, _ := NewMovie("Interstellar", *nationality, []Genre{*genre}, []Director{*director}, []Actor{*actor}, []Writer{}, 2014, "image456")

	// Criando instância para filme
	movie3, _ := NewMovie("The Dark Knight", *nationality, []Genre{*genre}, []Director{*director}, []Actor{*actor}, []Writer{}, 2008, "image789")

	// Criando instância para lista
	list, _ := NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	// Adicionando filmes à lista
	list.AddMovies([]Movie{*movie1, *movie2, *movie3})

	// Obtendo combinações de filmes
	combinations := list.GetAvailableMoviesCombinations()

	// Verificando se as combinações foram geradas corretamente
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
