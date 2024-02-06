package internal

import (
	"testing"
)

func TestMovieWriterEquals(t *testing.T) {
	mw1 := NewMovieWriter("aopishjd", "paosijdh")
	mw2 := NewMovieWriter("aopishjd", "paosijdh")
	
	if !mw1.Equals(mw2) {
		t.Errorf("Os MovieWriters deveriam ser iguais, mas não são.")
	}
	
	mw2.WriterID = "aksjgdd"
	
	if mw1.Equals(mw2) {
		t.Errorf("Os MovieWriters não deveriam ser iguais, mas são.")
	}
}