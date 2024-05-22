package repositoryinterface

import "youchoose/internal/entity"

type IMDBRepositoryInterface interface {
	Create(image *entity.IMDB) error
	CreateMany(images *[]entity.IMDB) error
	GetByID(imageID string) (bool, error)
}
