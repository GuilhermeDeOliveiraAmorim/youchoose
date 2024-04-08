package entity

import "youchoose/internal/util"

type ListMovie struct {
	SharedEntity
	ListID    string `json:"list_id"`
	MovieID   string `json:"movie_id"`
	ChooserID string `json:"chooser_id"`
}

func NewListMovie(listID, movieID, chooserID string) (*ListMovie, []util.ProblemDetails) {
	return &ListMovie{
		SharedEntity: *NewSharedEntity(),
		ListID:       listID,
		MovieID:      movieID,
		ChooserID:    chooserID,
	}, nil
}

func (lm *ListMovie) Equals(other *ListMovie) bool {
	return lm.ListID == other.ListID && lm.MovieID == other.MovieID
}
