package internal

import (
	"testing"
)

func TestListMovieEquals(t *testing.T) {
	lm1 := NewListMovie("oaishydf", "pposjddd")
	lm2 := NewListMovie("oaishydf", "pposjddd")

	if !lm1.Equals(lm2) {
		t.Errorf("Os ListMovies deveriam ser iguais, mas não são.")
	}

	lm2.MovieID = "apsoujd"

	if lm1.Equals(lm2) {
		t.Errorf("Os ListMovies não deveriam ser iguais, mas são.")
	}
}
