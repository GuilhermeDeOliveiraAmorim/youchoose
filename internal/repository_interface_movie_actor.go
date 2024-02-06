package internal

type MovieActorRepositoryInterface interface {
	Create(movieActor *MovieActor) error
	Update(movieActor *MovieActor) error
	GetByID(movieActorID string) (MovieActor, error)
	GetAll() ([]MovieActor, error)
	GetAllByMovieID(movieID string) ([]MovieActor, error)
	GetAllByActorID(actorID string) ([]MovieActor, error)
	Deactivate(movieActorID string) error
}
