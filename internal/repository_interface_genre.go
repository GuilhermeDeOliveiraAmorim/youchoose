package internal

type GenreRepositoryInterface interface {
	Create(genre *Genre) error
	Update(genre *Genre) error
	GetByID(genreID string) (Genre, error)
	GetAll() ([]Genre, error)
	GetAllByMovieID(movieID string) ([]Genre, error)
	Deactivate(genreID string) error
}
