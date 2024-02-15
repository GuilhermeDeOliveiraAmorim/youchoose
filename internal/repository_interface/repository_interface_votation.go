package repositoryinterface

import "youchoose/internal/entity"

type VotationRepositoryInterface interface {
	Create(votation *entity.Votation) error
	Update(votation *entity.Votation) error
	GetByID(votationID string) (entity.Votation, error)
	GetAll() ([]entity.Votation, error)
	GetActivesByListID(listID string) ([]entity.Votation, error)
	Deactivate(votationID string) error
}
