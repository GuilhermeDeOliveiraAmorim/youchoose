package internal

type CombinationRepositoryInterface interface {
	Create(combination *Combination) error
	Update(combination *Combination) error
	GetByID(combinationID string) (Combination, error)
	GetAll() ([]Combination, error)
	GetAllByVotationID(votationID string) ([]Combination, error)
	Deactivate(combinationID string) error
}
