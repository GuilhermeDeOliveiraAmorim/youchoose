package internal

type MovieDirectorRepositoryInterface interface {
	Create(movieDirector MovieDirector) error
	Update(movieDirector MovieDirector) error
	GetByID(movieDirectorID string) (MovieDirector, error)
	GetAll() ([]MovieDirector, error)
	GetAllByMovieID(movieID string) ([]MovieDirector, error)
	GetAllByDirectorID(directorID string) ([]MovieDirector, error)
	Deactivate(movieDirectorID string) error
}
