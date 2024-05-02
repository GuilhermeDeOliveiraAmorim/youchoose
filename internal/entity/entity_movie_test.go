package entity

import (
	"testing"

	valueobject "youchoose/internal/value_object"
)

func TestNewMovie(t *testing.T) {
	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	movie, err := NewMovie("Inception", *nationality, 2010, "image123")
	if err != nil {
		t.Errorf("Erro ao criar um novo filme válido: %v", err)
	}

	if movie.Title != "Inception" || movie.Nationality.CountryName != "United States" || movie.ReleaseYear != 2010 || movie.ImageID != "image123" {
		t.Errorf("O filme não foi criado corretamente. Detalhes do filme: %v", movie)
	}

	_, err = NewMovie("", valueobject.Nationality{CountryName: "United States", Flag: "🇺🇸"}, 2010, "image123")
	if err == nil {
		t.Error("Criou um filme com título vazio, mas deveria ter retornado um erro.")
	}
}

func TestMovie_AddActors(t *testing.T) {
	movie, _ := NewMovie("Inception", valueobject.Nationality{CountryName: "USA", Flag: "🇺🇸"}, 2010, "image123")

	newActors := []Actor{{Name: "Tom Hardy"}, {Name: "Ellen Page"}}
	movie.AddActors(newActors)

	if len(movie.Actors) != 2 || movie.Actors[0].Name != "Tom Hardy" || movie.Actors[1].Name != "Ellen Page" {
		t.Errorf("Erro ao adicionar atores. Detalhes do filme: %v", movie)
	}
}

func TestMovie_RemoveDirectors(t *testing.T) {
	birthDate, _ := valueobject.NewBirthDate(15, 5, 2010)

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	nolan, _ := NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")

	movie, _ := NewMovie("Inception", *nationality, 2010, "image123")

	movie.AddDirectors([]Director{*nolan})

	directorsToRemove := []Director{*nolan}
	movie.RemoveDirectors(directorsToRemove)

	if len(movie.Directors) != 1 {
		t.Errorf("Erro ao remover diretor. Detalhes do filme: %v", movie)
	}
}

func TestMovie_IncrementVotes(t *testing.T) {
	movie, _ := NewMovie("Inception", valueobject.Nationality{CountryName: "USA", Flag: "🇺🇸"}, 2010, "image123")

	movie.IncrementVotes()

	if movie.Votes != 1 {
		t.Errorf("Erro ao incrementar votos. Detalhes do filme: %v", movie)
	}
}
