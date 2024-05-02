package repositoryinterface

import "youchoose/internal/entity"

type ListFavoriteRepositoryInterface interface {
	Create(listFavorite *entity.ListFavorite) error
	Update(listFavorite *entity.ListFavorite) error
	GetByID(listFavoriteID string) (bool, entity.ListFavorite, error)
	GetByChooserIDAndListID(chooserID, listID string) (bool, entity.ListFavorite, error)
	Deactivate(listFavorite *entity.ListFavorite) error
	GetAll() ([]entity.ListFavorite, error)
	GetAllByListID(listID string) ([]entity.ListFavorite, error)
}
