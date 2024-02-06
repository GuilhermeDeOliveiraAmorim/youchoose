package repositoryinterface

import "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/entity"

type MovieActorRepositoryInterface interface {
	Create(movieActor *entity.MovieActor) error
	Update(movieActor *entity.MovieActor) error
	GetByID(movieActorID string) (entity.MovieActor, error)
	GetAll() ([]entity.MovieActor, error)
	GetAllByMovieID(movieID string) ([]entity.MovieActor, error)
	GetAllByActorID(actorID string) ([]entity.MovieActor, error)
	Deactivate(movieActorID string) error
}
