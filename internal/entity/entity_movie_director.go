package entity

type MovieDirector struct {
	SharedEntity
	MovieID    string `json:"movie_id"`
	DirectorID string `json:"director_id"`
}

func NewMovieDirector(movieID, directorID string) *MovieDirector {
	return &MovieDirector{
		SharedEntity: *NewSharedEntity(),
		MovieID:      movieID,
		DirectorID:   directorID,
	}
}

func (md *MovieDirector) Equals(other *MovieDirector) bool {
	return md.MovieID == other.MovieID && md.DirectorID == other.DirectorID
}
