package internal

type VotationRepositoryInterface interface {
	Create(votation *Votation) error
	Update(votation *Votation) error
	GetByID(votationID string) (Votation, error)
	GetAll() ([]Votation, error)
	GetActivesByListID(listID string) ([]Votation, error)
	Deactivate(votationID string) error
}
