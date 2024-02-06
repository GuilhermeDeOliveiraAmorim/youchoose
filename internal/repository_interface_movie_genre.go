package internal

type MovieGenreRepositoryInterface interface {
	Create(movieGenre MovieGenre) error
	Update(movieGenre MovieGenre) error
	GetByID(movieGenreID string) (MovieGenre, error)
	GetAll() ([]MovieGenre, error)
	GetAllByMovieID(movieID string) ([]MovieGenre, error)
	GetAllByGenreID(genreID string) ([]MovieGenre, error)
	Deactivate(movieGenreID string) error
}
