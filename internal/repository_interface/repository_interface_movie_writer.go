package repositoryinterface

import "youchoose/internal/entity"

type MovieWriterRepositoryInterface interface {
	Create(movieWriter *entity.MovieWriter) error
	Update(movieWriter *entity.MovieWriter) error
	GetByID(movieWriterID string) (entity.MovieWriter, error)
	GetAll() ([]entity.MovieWriter, error)
	GetAllByMovieID(movieID string) ([]entity.MovieWriter, error)
	GetAllByWriterID(writerID string) ([]entity.MovieWriter, error)
	Deactivate(movieWriter *entity.MovieWriter) error
}
