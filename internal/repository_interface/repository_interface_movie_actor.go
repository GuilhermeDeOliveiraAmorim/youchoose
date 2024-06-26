package repositoryinterface

import "youchoose/internal/entity"

type MovieActorRepositoryInterface interface {
	Create(movieActor *entity.MovieActor) error
	CreateMany(movieActors *[]entity.MovieActor) error
	Update(movieActor *entity.MovieActor) error
	GetByID(movieActorID string) (entity.MovieActor, error)
	GetAll() ([]entity.MovieActor, error)
	GetAllByMovieID(movieID string) ([]entity.MovieActor, error)
	GetAllByActorID(actorID string) ([]entity.MovieActor, error)
	Deactivate(movieActor *entity.MovieActor) error
}
