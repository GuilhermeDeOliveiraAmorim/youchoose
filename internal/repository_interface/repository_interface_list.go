package repositoryinterface

import "youchoose/internal/entity"

type ListRepositoryInterface interface {
	Create(list *entity.List) error
	Update(list *entity.List) error
	GetByID(listID string) (bool, entity.List, error)
	GetAll() ([]entity.List, error)
	Deactivate(list *entity.List) error
	GetAllMoviesByListID(listID string) ([]entity.Movie, error)
	GetAllMoviesCombinationsByListID(listID string) ([][]entity.Movie, error)
}
