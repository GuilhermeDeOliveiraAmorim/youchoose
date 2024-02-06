package internal

type ListRepositoryInterface interface {
	Create(list *List) error
	Update(list *List) error
	GetByID(listID string) (List, error)
	GetAll() ([]List, error)
	Deactivate(listID string) error
	GetAllMoviesByListID(listID string) ([]Movie, error)
	GetAllMoviesCombinationsByListID(listID string) ([]Combination, error)
}
