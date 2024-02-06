package internal

import "testing"

func TestMovieGenreEquals(t *testing.T) {
	mg1 := NewMovieGenre(1, 2)
	mg2 := NewMovieGenre(1, 2)
	
	if !mg1.Equals(mg2) {
		t.Errorf("Os MovieGenres deveriam ser iguais, mas não são.")
	}
	
	mg2.GenreID = 3
	
	if mg1.Equals(mg2) {
		t.Errorf("Os MovieGenres não deveriam ser iguais, mas são.")
	}
}
