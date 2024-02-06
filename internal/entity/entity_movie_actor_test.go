package entity

import "testing"

func TestMovieActorEquals(t *testing.T) {
	ma1 := NewMovieActor("asdasd", "asdasd")
	ma2 := NewMovieActor("asdasd", "asdasd")

	if !ma1.Equals(ma2) {
		t.Errorf("Os MovieActors deveriam ser iguais, mas não são.")
	}

	ma2.ActorID = "sdfsdaf"

	if ma1.Equals(ma2) {
		t.Errorf("Os MovieActors não deveriam ser iguais, mas são.")
	}
}
