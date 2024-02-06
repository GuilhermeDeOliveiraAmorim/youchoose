package repositoryinterface

import "youchoose/internal/entity"

type MovieGenreRepositoryInterface interface {
	Create(movieGenre *entity.MovieGenre) error
	Update(movieGenre *entity.MovieGenre) error
	GetByID(movieGenreID string) (entity.MovieGenre, error)
	GetAll() ([]entity.MovieGenre, error)
	GetAllByMovieID(movieID string) ([]entity.MovieGenre, error)
	GetAllByGenreID(genreID string) ([]entity.MovieGenre, error)
	Deactivate(movieGenreID string) error
}
