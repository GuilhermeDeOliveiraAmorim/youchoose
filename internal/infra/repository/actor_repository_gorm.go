package gorm

import (
	"gorm.io/gorm"
	"youchoose/internal/entity"
)

type ActorRepository struct {
	gorm *gorm.DB
}

// Create implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) Create(actor *entity.Actor) error {
	panic("unimplemented")
}

// CreateMany implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) CreateMany(actors *[]entity.Actor) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) Deactivate(actor *entity.Actor) error {
	panic("unimplemented")
}

// DoTheseActorsAreIncludedInTheMovie implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) DoTheseActorsAreIncludedInTheMovie(movieID string, actorsIDs []string) (bool, []entity.Actor, error) {
	panic("unimplemented")
}

// DoTheseActorsExist implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) DoTheseActorsExist(actorIDs []string) (bool, []entity.Actor, error) {
	panic("unimplemented")
}

// GetAll implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) GetAll() ([]entity.Actor, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) GetByID(actorID string) (entity.Actor, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) Update(actor *entity.Actor) error {
	panic("unimplemented")
}

func NewActorRepository(gorm *gorm.DB) *ActorRepository {
	return &ActorRepository{
		gorm: gorm,
	}
}
