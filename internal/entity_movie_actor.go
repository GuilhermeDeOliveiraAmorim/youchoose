package internal

type MovieActor struct {
	SharedEntity
	MovieID string `json:"movie_id"`
	ActorID string `json:"actor_id"`
}

func NewMovieActor(movieID, actorID string) *MovieActor {
	return &MovieActor{
		SharedEntity: *NewSharedEntity(),
		MovieID:      movieID,
		ActorID:      actorID,
	}
}

func (ma *MovieActor) Equals(other *MovieActor) bool {
	return ma.MovieID == other.MovieID && ma.ActorID == other.ActorID
}
