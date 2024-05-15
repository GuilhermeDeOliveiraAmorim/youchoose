package gorm

import (
	"gorm.io/gorm"
	"youchoose/internal/entity"
)

type MovieActorRepository struct {
	gorm *gorm.DB
}

// Create implements repositoryinterface.MovieActorRepositoryInterface.
func (m *MovieActorRepository) Create(movieActor *entity.MovieActor) error {
	panic("unimplemented")
}

// CreateMany implements repositoryinterface.MovieActorRepositoryInterface.
func (m *MovieActorRepository) CreateMany(movieActors *[]entity.MovieActor) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.MovieActorRepositoryInterface.
func (m *MovieActorRepository) Deactivate(movieActor *entity.MovieActor) error {
	panic("unimplemented")
}

// GetAll implements repositoryinterface.MovieActorRepositoryInterface.
func (m *MovieActorRepository) GetAll() ([]entity.MovieActor, error) {
	panic("unimplemented")
}

// GetAllByActorID implements repositoryinterface.MovieActorRepositoryInterface.
func (m *MovieActorRepository) GetAllByActorID(actorID string) ([]entity.MovieActor, error) {
	panic("unimplemented")
}

// GetAllByMovieID implements repositoryinterface.MovieActorRepositoryInterface.
func (m *MovieActorRepository) GetAllByMovieID(movieID string) ([]entity.MovieActor, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.MovieActorRepositoryInterface.
func (m *MovieActorRepository) GetByID(movieActorID string) (entity.MovieActor, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.MovieActorRepositoryInterface.
func (m *MovieActorRepository) Update(movieActor *entity.MovieActor) error {
	panic("unimplemented")
}

func NewMovieActorRepository(gorm *gorm.DB) *MovieActorRepository {
	return &MovieActorRepository{
		gorm: gorm,
	}
}
