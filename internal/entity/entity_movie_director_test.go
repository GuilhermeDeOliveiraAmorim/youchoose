package entity

import (
	"testing"
)

func TestMovieDirectorEquals(t *testing.T) {
	md1 := NewMovieDirector("asdasd", "alksjd")
	md2 := NewMovieDirector("asdasd", "alksjd")

	if !md1.Equals(md2) {
		t.Errorf("Os MovieDirectors deveriam ser iguais, mas não são.")
	}

	md2.DirectorID = "alskjd"

	if md1.Equals(md2) {
		t.Errorf("Os MovieDirectors não deveriam ser iguais, mas são.")
	}
}
