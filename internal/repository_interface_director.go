package internal

type DirectorRepositoryInterface interface {
	Create(director *Director) error
	Update(director *Director) error
	GetByID(directorID string) (Director, error)
	GetAll() ([]Director, error)
	Deactivate(directorID string) error
}
