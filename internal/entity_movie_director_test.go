package internal

import (
	"testing"
)

func TestMovieDirectorEquals(t *testing.T) {
	md1 := NewMovieDirector(1, 2)
	md2 := NewMovieDirector(1, 2)
	
	if !md1.Equals(md2) {
		t.Errorf("Os MovieDirectors deveriam ser iguais, mas não são.")
	}
	
	md2.DirectorID = 3
	
	if md1.Equals(md2) {
		t.Errorf("Os MovieDirectors não deveriam ser iguais, mas são.")
	}
}