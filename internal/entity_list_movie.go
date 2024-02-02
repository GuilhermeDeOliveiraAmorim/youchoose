package internal

type ListMovie struct {
	SharedEntity
	ListID  int `json:"list_id"`
	MovieID int `json:"movie_id"`
}

func NewListMovie(listID, movieID int) *ListMovie {
	return &ListMovie{
		SharedEntity: *NewSharedEntity(),
		ListID:  listID,
		MovieID: movieID,
	}
}

func (lm *ListMovie) Equals(other *ListMovie) bool {
    return lm.ListID == other.ListID && lm.MovieID == other.MovieID
}