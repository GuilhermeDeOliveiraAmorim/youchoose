package internal

type MovieRepositoryInterface interface {
	Create(movie *Movie) error
	Update(movie *Movie) error
	GetByID(movieID string) (Movie, error)
	GetAll() ([]Movie, error)
	GetByActorID(actorID string) ([]Movie, error)
	GetByDirectorID(directorID string) ([]Movie, error)
	GetByGenreID(genreID string) ([]Movie, error)
	GetByWriterID(writerID string) ([]Movie, error)
	Deactivate(movieID string) error
}
