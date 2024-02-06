package internal

type ListMovie struct {
	SharedEntity
	ListID  string `json:"list_id"`
	MovieID string `json:"movie_id"`
}

func NewListMovie(listID, movieID string) *ListMovie {
	return &ListMovie{
		SharedEntity: *NewSharedEntity(),
		ListID:  listID,
		MovieID: movieID,
	}
}

func (lm *ListMovie) Equals(other *ListMovie) bool {
    return lm.ListID == other.ListID && lm.MovieID == other.MovieID
}