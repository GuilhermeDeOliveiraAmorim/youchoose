package repositoryinterface

import "youchoose/internal/entity"

type VotationRepositoryInterface interface {
	Create(votation *entity.Votation) error
	Update(votation *entity.Votation) error
	GetByID(votationID string) (bool, entity.Votation, error)
	GetAll() ([]entity.Votation, error)
	GetAllByListIDAndChooserID(listID, chooserID string) ([]entity.Votation, error)
	Deactivate(votationID string) error
}
