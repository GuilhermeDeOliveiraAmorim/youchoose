package entity

import (
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestListMovieEquals(t *testing.T) {
	lm1, _ := NewListMovie("oaishydf", "pposjddd", ulid.Make().String())
	lm2, _ := NewListMovie("oaishydf", "pposjddd", ulid.Make().String())

	if !lm1.Equals(lm2) {
		t.Errorf("Os ListMovies deveriam ser iguais, mas não são.")
	}

	lm2.MovieID = "apsoujd"

	if lm1.Equals(lm2) {
		t.Errorf("Os ListMovies não deveriam ser iguais, mas são.")
	}
}
