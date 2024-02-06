package internal

type MovieWriterRepositoryInterface interface {
	Create(movieWriter MovieWriter) error
	Update(movieWriter MovieWriter) error
	GetByID(movieWriterID string) (MovieWriter, error)
	GetAll() ([]MovieWriter, error)
	GetAllByMovieID(movieID string) ([]MovieWriter, error)
	GetAllByWriterID(writerID string) ([]MovieWriter, error)
	Deactivate(movieWriterID string) error
}
