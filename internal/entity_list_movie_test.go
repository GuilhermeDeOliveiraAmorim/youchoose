package internal

import (
	"testing"
)

func TestListMovieEquals(t *testing.T) {
	// Criar dois ListMovies iguais
	lm1 := NewListMovie(1, 2)
	lm2 := NewListMovie(1, 2)

	// Verificar se os ListMovies são iguais
	if !lm1.Equals(lm2) {
		t.Errorf("Os ListMovies deveriam ser iguais, mas não são.")
	}

	// Modificar um ListMovie para torná-lo diferente
	lm2.MovieID = 3

	// Verificar se os ListMovies são diferentes agora
	if lm1.Equals(lm2) {
		t.Errorf("Os ListMovies não deveriam ser iguais, mas são.")
	}
}