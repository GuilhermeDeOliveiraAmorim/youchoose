package internal

import (
	"testing"
)

func TestMovieWriterEquals(t *testing.T) {
	mw1 := NewMovieWriter(1, 2)
	mw2 := NewMovieWriter(1, 2)
	
	if !mw1.Equals(mw2) {
		t.Errorf("Os MovieWriters deveriam ser iguais, mas não são.")
	}
	
	mw2.WriterID = 3
	
	if mw1.Equals(mw2) {
		t.Errorf("Os MovieWriters não deveriam ser iguais, mas são.")
	}
}