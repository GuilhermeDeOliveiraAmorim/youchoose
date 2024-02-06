package internal

type ChooserRepositoryInterface interface {
	Create(chooser *Chooser) error
	Update(chooser *Chooser) error
	GetByID(chooserID string) (Chooser, error)
	GetAll() ([]Chooser, error)
	Deactivate(chooserID string) error
	GetLists(chooserID string) error
	GetVotation(chooserID, listID string) (Votation, error)
	GetVotations(chooserID string) ([]Votation, error)
}
