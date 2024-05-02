package repositoryinterface

import "youchoose/internal/entity"

type MovieRepositoryInterface interface {
	Create(movie *entity.Movie) error
	Update(movie *entity.Movie) error
	GetByID(movieID string) (bool, entity.Movie, error)
	GetAll() ([]entity.Movie, error)
	GetByActorID(actorID string) ([]entity.Movie, error)
	GetByDirectorID(directorID string) ([]entity.Movie, error)
	GetByGenreID(genreID string) ([]entity.Movie, error)
	GetByWriterID(writerID string) ([]entity.Movie, error)
	DoTheseMoviesExist(movieIDs []string) (bool, []entity.Movie, error)
	Deactivate(movie *entity.Movie) error
}
