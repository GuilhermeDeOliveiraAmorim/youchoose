package entity

import "testing"

func TestMovieGenreEquals(t *testing.T) {
	mg1 := NewMovieGenre("oaishd", "aoishd")
	mg2 := NewMovieGenre("oaishd", "aoishd")

	if !mg1.Equals(mg2) {
		t.Errorf("Os MovieGenres deveriam ser iguais, mas não são.")
	}

	mg2.GenreID = "apoisud"

	if mg1.Equals(mg2) {
		t.Errorf("Os MovieGenres não deveriam ser iguais, mas são.")
	}
}
