package repositoryinterface

import "youchoose/internal/entity"

type ActorRepositoryInterface interface {
	Create(actor *entity.Actor) error
	CreateMany(actors *[]entity.Actor) error
	Update(actor *entity.Actor) error
	GetByID(actorID string) (entity.Actor, error)
	GetAll() ([]entity.Actor, error)
	Deactivate(actor *entity.Actor) error
	DoTheseActorsExist(actorIDs []string) (bool, []entity.Actor, error)
	DoTheseActorsAreIncludedInTheMovie(movieID string, actorsIDs []string) (bool, []entity.Actor, error)
}
