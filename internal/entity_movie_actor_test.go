package internal

import "testing"

func TestMovieActorEquals(t *testing.T) {
	ma1 := NewMovieActor(1, 2)
	ma2 := NewMovieActor(1, 2)
	
	if !ma1.Equals(ma2) {
		t.Errorf("Os MovieActors deveriam ser iguais, mas não são.")
	}
	
	ma2.ActorID = 3
	
	if ma1.Equals(ma2) {
		t.Errorf("Os MovieActors não deveriam ser iguais, mas são.")
	}
}
