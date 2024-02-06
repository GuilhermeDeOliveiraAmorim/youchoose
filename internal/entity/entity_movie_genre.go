package entity

type MovieGenre struct {
	SharedEntity
	MovieID string `json:"movie_id"`
	GenreID string `json:"genre_id"`
}

func NewMovieGenre(movieID, genreID string) *MovieGenre {
	return &MovieGenre{
		SharedEntity: *NewSharedEntity(),
		MovieID:      movieID,
		GenreID:      genreID,
	}
}

func (mg *MovieGenre) Equals(other *MovieGenre) bool {
	return mg.MovieID == other.MovieID && mg.GenreID == other.GenreID
}
