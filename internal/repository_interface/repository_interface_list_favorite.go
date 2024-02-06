package repositoryinterface

import "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/entity"

type ListFavoriteRepositoryInterface interface {
	Create(listFavorite *entity.ListFavorite) error
	Update(listFavorite *entity.ListFavorite) error
	GetByID(listFavoriteID string) (entity.ListFavorite, error)
	Deactivate(listFavoriteID string) error
	GetAll() ([]entity.ListFavorite, error)
	GetAllByListID(listID string) ([]entity.ListFavorite, error)
}
