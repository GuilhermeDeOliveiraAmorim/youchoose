package repositoryinterface

import "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/entity"

type ActorRepositoryInterface interface {
	Create(actor *entity.Actor) error
	Update(actor *entity.Actor) error
	GetByID(actorID string) (entity.Actor, error)
	GetAll() ([]entity.Actor, error)
	Deactivate(actorID string) error
}
