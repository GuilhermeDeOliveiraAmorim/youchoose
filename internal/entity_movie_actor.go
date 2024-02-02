package internal

type MovieActor struct {
	SharedEntity
	MovieID int `json:"movie_id"`
	ActorID int `json:"actor_id"`
}

func NewMovieActor(movieID, actorID int) *MovieActor {
	return &MovieActor{
		SharedEntity: *NewSharedEntity(),
		MovieID:      movieID,
		ActorID:      actorID,
	}
}

func (ma *MovieActor) Equals(other *MovieActor) bool {
	return ma.MovieID == other.MovieID && ma.ActorID == other.ActorID
}
