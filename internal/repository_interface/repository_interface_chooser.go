package repositoryinterface

import "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/entity"

type ChooserRepositoryInterface interface {
	Create(chooser *entity.Chooser) error
	Update(chooser *entity.Chooser) error
	GetByID(chooserID string) (entity.Chooser, error)
	GetAll() ([]entity.Chooser, error)
	Deactivate(chooserID string) error
	GetLists(chooserID string) error
	GetVotation(chooserID, listID string) (entity.Votation, error)
	GetVotations(chooserID string) ([]entity.Votation, error)
}
