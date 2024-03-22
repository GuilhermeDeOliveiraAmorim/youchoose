package repositoryinterface

import "youchoose/internal/entity"

type ListMovieRepositoryInterface interface {
	Create(listMovies []entity.ListMovie) error
	Update(listMovie *entity.ListMovie) error
	GetByID(listMovieID string) (entity.ListMovie, error)
	GetAll() ([]*entity.ListMovie, error)
	GetAllByListID(listID string) ([]entity.ListMovie, error)
	Deactivate(listMovieID string) error
}
