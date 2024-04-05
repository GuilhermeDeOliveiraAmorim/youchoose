package repositoryinterface

import "youchoose/internal/entity"

type GenreRepositoryInterface interface {
	Create(genre *entity.Genre) error
	Update(genre *entity.Genre) error
	GetByID(genreID string) (entity.Genre, error)
	GetAll() ([]entity.Genre, error)
	GetAllByMovieID(movieID string) ([]entity.Genre, error)
	Deactivate(genre *entity.Genre) error
	DoTheseGenresExist(genreIDs []string) (bool, []entity.Genre, error)
}
