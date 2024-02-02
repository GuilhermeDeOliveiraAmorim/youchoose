package internal

type MovieGenre struct {
	SharedEntity
	MovieID int `json:"movie_id"`
	GenreID int `json:"genre_id"`
}

func NewMovieGenre(movieID, genreID int) *MovieGenre {
	return &MovieGenre{
		SharedEntity: *NewSharedEntity(),
		MovieID:      movieID,
		GenreID:      genreID,
	}
}

func (mg *MovieGenre) Equals(other *MovieGenre) bool {
	return mg.MovieID == other.MovieID && mg.GenreID == other.GenreID
}
