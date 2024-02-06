package repositoryinterface

import "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/entity"

type GenreRepositoryInterface interface {
	Create(genre *entity.Genre) error
	Update(genre *entity.Genre) error
	GetByID(genreID string) (entity.Genre, error)
	GetAll() ([]entity.Genre, error)
	GetAllByMovieID(movieID string) ([]entity.Genre, error)
	Deactivate(genreID string) error
}
