package repositoryinterface

import "youchoose/internal/entity"

type CombinationRepositoryInterface interface {
	Create(combination *entity.Combination) error
	Update(combination *entity.Combination) error
	GetByID(combinationID string) (entity.Combination, error)
	GetAll() ([]entity.Combination, error)
	GetAllByVotationID(votationID string) ([]entity.Combination, error)
	Deactivate(combinationID string) error
}
