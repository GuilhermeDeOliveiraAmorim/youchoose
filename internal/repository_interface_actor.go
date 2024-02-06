package internal

type ActorRepositoryInterface interface {
	Create(actor *Actor) error
	Update(actor *Actor) error
	GetByID(actorID string) (Actor, error)
	GetAll() ([]Actor, error)
	Deactivate(actorID string) error
}
