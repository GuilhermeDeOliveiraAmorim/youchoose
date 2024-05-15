package gorm

import (
	"errors"
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type MovieActorRepository struct {
	gorm *gorm.DB
}

func NewMovieActorRepository(gorm *gorm.DB) *MovieActorRepository {
	return &MovieActorRepository{
		gorm: gorm,
	}
}

func (m *MovieActorRepository) Create(movieActor *entity.MovieActor) error {
	panic("unimplemented")
}

func (m *MovieActorRepository) CreateMany(movieActors *[]entity.MovieActor) error {
	var movieActorsModel []MovieActors

	for _, movieActor := range *movieActors {
		movieActorsModel = append(movieActorsModel, MovieActors{
			ID:            movieActor.ID,
			Active:        movieActor.Active,
			CreatedAt:     movieActor.CreatedAt,
			UpdatedAt:     movieActor.UpdatedAt,
			DeactivatedAt: movieActor.DeactivatedAt,
			MovieID:       movieActor.MovieID,
			ActorID:       movieActor.ActorID,
		})
	}

	if err := m.gorm.Create(movieActorsModel).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (m *MovieActorRepository) Deactivate(movieActor *entity.MovieActor) error {
	panic("unimplemented")
}

func (m *MovieActorRepository) GetAll() ([]entity.MovieActor, error) {
	panic("unimplemented")
}

func (m *MovieActorRepository) GetAllByActorID(actorID string) ([]entity.MovieActor, error) {
	panic("unimplemented")
}

func (m *MovieActorRepository) GetAllByMovieID(movieID string) ([]entity.MovieActor, error) {
	panic("unimplemented")
}

func (m *MovieActorRepository) GetByID(movieActorID string) (entity.MovieActor, error) {
	panic("unimplemented")
}

func (m *MovieActorRepository) Update(movieActor *entity.MovieActor) error {
	panic("unimplemented")
}
