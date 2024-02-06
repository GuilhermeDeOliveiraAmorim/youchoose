package repositoryinterface

import "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/entity"

type ImageRepositoryInterface interface {
	Create(image *entity.Image) error
	Update(image *entity.Image) error
	GetByID(imageID string) (entity.Image, error)
	GetAll() ([]entity.Image, error)
	Deactivate(imageID string) error
}
