package internal

type ListMovieRepositoryInterface interface {
	Create(listMovie ListMovie) error
	Update(listMovie ListMovie) error
	GetByID(listMovieID string) (ListMovie, error)
	GetAll() ([]*ListMovie, error)
	GetAllByListID(listID string) ([]ListMovie, error)
	Deactivate(listMovieID string) error
}
