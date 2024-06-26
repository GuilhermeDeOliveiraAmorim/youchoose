package repositoryinterface

import "youchoose/internal/entity"

type ListMovieRepositoryInterface interface {
	CreateMany(listMovies *[]entity.ListMovie) error
	Update(listMovie *entity.ListMovie) error
	GetByID(listMovieID string) (entity.ListMovie, error)
	GetByListIDAndMovieIDAndChooserID(listID, movieID, chooserID string) (bool, entity.ListMovie, error)
	GetAll() ([]*entity.ListMovie, error)
	GetAllByListID(listID string) ([]entity.ListMovie, error)
	Deactivate(listMovie *entity.ListMovie) error
	DeactivateAll(movies *[]entity.Movie) error
}
