package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestListMovieEquals(t *testing.T) {
	lm1, _ := NewListMovie("oaishydf", "pposjddd", uuid.NewString())
	lm2, _ := NewListMovie("oaishydf", "pposjddd", uuid.NewString())

	if !lm1.Equals(lm2) {
		t.Errorf("Os ListMovies deveriam ser iguais, mas não são.")
	}

	lm2.MovieID = "apsoujd"

	if lm1.Equals(lm2) {
		t.Errorf("Os ListMovies não deveriam ser iguais, mas são.")
	}
}
