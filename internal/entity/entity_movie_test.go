package entity

import (
	"fmt"
	"testing"

	valueobject "youchoose/internal/value_object"
)

func TestNewMovie(t *testing.T) {
	birthDate, err := valueobject.NewBirthDate(15, 5, 2010)
	if err != nil {
		fmt.Println(err)
	}

	nationality, err := valueobject.NewNationality("United States", "🇺🇸")
	if err != nil {
		fmt.Println(err)
	}

	tomHardy, err := NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")
	if err != nil {
		fmt.Println(err)
	}

	ellenPage, err := NewActor("Ellen Page", birthDate, nationality, "ellen_page_image")
	if err != nil {
		fmt.Println(err)
	}

	nolan, err := NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")
	if err != nil {
		fmt.Println(err)
	}

	nolanWriter, err := NewWriter("Christopher Nolan", birthDate, nationality, "nolan_image")
	if err != nil {
		fmt.Println(err)
	}

	sciFi, err := NewGenre("Sci-Fi", "image_id_genre")
	if err != nil {
		fmt.Println(err)
	}

	movie, err := NewMovie("Inception", *nationality, []Genre{*sciFi}, []Director{*nolan}, []Actor{*tomHardy, *ellenPage}, []Writer{*nolanWriter}, 2010, "image123")
	if err != nil {
		t.Errorf("Erro ao criar um novo filme válido: %v", err)
	}

	if movie.Title != "Inception" || movie.Nationality.CountryName != "United States" || movie.Genres[0].Name != "Sci-Fi" || movie.Directors[0].Name != "Christopher Nolan" || movie.Writers[0].Name != "Christopher Nolan" || movie.ReleaseYear != 2010 || movie.ImageID != "image123" {
		t.Errorf("O filme não foi criado corretamente. Detalhes do filme: %v", movie)
	}

	_, err = NewMovie("", valueobject.Nationality{CountryName: "United States", Flag: "🇺🇸"}, []Genre{{Name: "Sci-Fi"}}, []Director{{Name: "Christopher Nolan"}}, []Actor{{Name: "Leonardo DiCaprio"}}, []Writer{{Name: "Christopher Nolan"}}, 2010, "image123")
	if err == nil {
		t.Error("Criou um filme com título vazio, mas deveria ter retornado um erro.")
	}
}

func TestMovie_AddActors(t *testing.T) {
	movie, _ := NewMovie("Inception", valueobject.Nationality{CountryName: "USA", Flag: "🇺🇸"}, []Genre{{Name: "Sci-Fi"}}, []Director{{Name: "Christopher Nolan"}}, []Actor{{Name: "Leonardo DiCaprio"}}, []Writer{{Name: "Christopher Nolan"}}, 2010, "image123")

	newActors := []Actor{{Name: "Tom Hardy"}, {Name: "Ellen Page"}}
	movie.AddActors(newActors)

	if len(movie.Actors) != 3 || movie.Actors[1].Name != "Tom Hardy" || movie.Actors[2].Name != "Ellen Page" {
		t.Errorf("Erro ao adicionar atores. Detalhes do filme: %v", movie)
	}
}

func TestMovie_RemoveDirectors(t *testing.T) {
	birthDate, err := valueobject.NewBirthDate(15, 5, 2010)
	if err != nil {
		fmt.Println(err)
	}

	nationality, err := valueobject.NewNationality("United States", "🇺🇸")
	if err != nil {
		fmt.Println(err)
	}

	tomHardy, err := NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")
	if err != nil {
		fmt.Println(err)
	}

	ellenPage, err := NewActor("Ellen Page", birthDate, nationality, "ellen_page_image")
	if err != nil {
		fmt.Println(err)
	}

	nolan, err := NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")
	if err != nil {
		fmt.Println(err)
	}

	nolanWriter, err := NewWriter("Christopher Nolan", birthDate, nationality, "nolan_image")
	if err != nil {
		fmt.Println(err)
	}

	sciFi, err := NewGenre("Sci-Fi", "image_id_genre")
	if err != nil {
		fmt.Println(err)
	}

	movie, _ := NewMovie("Inception", *nationality, []Genre{*sciFi}, []Director{*nolan}, []Actor{*tomHardy, *ellenPage}, []Writer{*nolanWriter}, 2010, "image123")

	directorsToRemove := []Director{*nolan}
	movie.RemoveDirectors(directorsToRemove)

	if len(movie.Directors) != 1 {
		t.Errorf("Erro ao remover diretor. Detalhes do filme: %v", movie)
	}
}

func TestMovie_IncrementVotes(t *testing.T) {
	movie, _ := NewMovie("Inception", valueobject.Nationality{CountryName: "USA", Flag: "🇺🇸"}, []Genre{{Name: "Sci-Fi"}}, []Director{{Name: "Christopher Nolan"}}, []Actor{{Name: "Leonardo DiCaprio"}}, []Writer{{Name: "Christopher Nolan"}}, 2010, "image123")

	movie.IncrementVotes()

	if movie.Votes != 1 {
		t.Errorf("Erro ao incrementar votos. Detalhes do filme: %v", movie)
	}
}
