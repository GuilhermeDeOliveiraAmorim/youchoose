package internal

type ImageRepositoryInterface interface {
	Create(image *Image) error
    Update(image *Image) error
    GetByID(imageID string) (Image, error)
    GetAll() ([]Image, error)
    Deactivate(imageID string) error
}
