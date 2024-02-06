package repositoryinterface

import "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/entity"

type MovieDirectorRepositoryInterface interface {
	Create(movieDirector *entity.MovieDirector) error
	Update(movieDirector *entity.MovieDirector) error
	GetByID(movieDirectorID string) (entity.MovieDirector, error)
	GetAll() ([]entity.MovieDirector, error)
	GetAllByMovieID(movieID string) ([]entity.MovieDirector, error)
	GetAllByDirectorID(directorID string) ([]entity.MovieDirector, error)
	Deactivate(movieDirectorID string) error
}
