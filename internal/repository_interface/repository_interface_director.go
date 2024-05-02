package repositoryinterface

import "youchoose/internal/entity"

type DirectorRepositoryInterface interface {
	Create(director *entity.Director) error
	Update(director *entity.Director) error
	GetByID(directorID string) (entity.Director, error)
	GetAll() ([]entity.Director, error)
	Deactivate(director *entity.Director) error
	DoTheseDirectorsExist(directorIDs []string) (bool, []entity.Director, error)
	DoTheseDirectorsAreIncludedInTheMovie(movieID string, directorsIDs []string) (bool, []entity.Director, error)
}
