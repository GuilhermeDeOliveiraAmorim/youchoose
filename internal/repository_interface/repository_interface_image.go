package repositoryinterface

import "youchoose/internal/entity"

type ImageRepositoryInterface interface {
	Create(image *entity.Image) error
	Update(image *entity.Image) error
	GetByID(imageID string) (entity.Image, error)
	GetAll() ([]entity.Image, error)
	Deactivate(image *entity.Image) error
}
