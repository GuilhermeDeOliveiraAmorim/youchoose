package internal

type MovieDirector struct {
	SharedEntity
	MovieID     int `json:"movie_id"`
	DirectorID  int `json:"director_id"`
}

func NewMovieDirector(movieID, directorID int) *MovieDirector {
	return &MovieDirector{
		SharedEntity: *NewSharedEntity(),
		MovieID:      movieID,
		DirectorID:   directorID,
	}
}

func (md *MovieDirector) Equals(other *MovieDirector) bool {
	return md.MovieID == other.MovieID && md.DirectorID == other.DirectorID
}
