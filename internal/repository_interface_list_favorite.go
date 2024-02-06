package internal

type ListFavoriteRepositoryInterface interface {
	Create(listFavorite *ListFavorite) error
	Update(listFavorite *ListFavorite) error
	GetByID(listFavoriteID string) (ListFavorite, error)
	Deactivate(listFavoriteID string) error
	GetAll() ([]ListFavorite, error)
	GetAllByListID(listID string) ([]ListFavorite, error)
}
