package gorm

import (
	"errors"
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type ImageRepository struct {
	gorm *gorm.DB
}

func NewImageRepository(gorm *gorm.DB) *ImageRepository {
	return &ImageRepository{
		gorm: gorm,
	}
}

func (i *ImageRepository) Create(image *entity.Image) error {
	if err := i.gorm.Create(&Images{
		ID:            image.ID,
		Active:        image.Active,
		CreatedAt:     image.CreatedAt,
		UpdatedAt:     image.UpdatedAt,
		DeactivatedAt: image.DeactivatedAt,
		Name:          image.Name,
		Type:          image.Type,
		Size:          image.Size,
	}).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (i *ImageRepository) CreateMany(images *[]entity.Image) error {
	var imagesModel []Images

	for _, image := range *images {
		imagesModel = append(imagesModel, Images{
			ID:            image.ID,
			Active:        image.Active,
			CreatedAt:     image.CreatedAt,
			UpdatedAt:     image.UpdatedAt,
			DeactivatedAt: image.DeactivatedAt,
			Name:          image.Name,
			Type:          image.Type,
			Size:          image.Size,
		})
	}

	if err := i.gorm.Create(imagesModel).Error; err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (i *ImageRepository) Deactivate(image *entity.Image) error {
	err := i.gorm.Model(&Images{}).Where("id = ?", image.ID).Updates(map[string]interface{}{
		"Active":        image.Active,
		"DeactivatedAt": image.DeactivatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (i *ImageRepository) GetAll() ([]entity.Image, error) {
	panic("unimplemented")
}

func (i *ImageRepository) GetByID(imageID string) (bool, entity.Image, error) {
	var imageModel Images

	result := i.gorm.Model(&Images{}).Where("id = ?", imageID).First(&imageModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, entity.Image{}, nil
		}
		return false, entity.Image{}, result.Error
	}

	image := entity.Image{
		SharedEntity: entity.SharedEntity{
			ID:            imageModel.ID,
			Active:        imageModel.Active,
			CreatedAt:     imageModel.CreatedAt,
			UpdatedAt:     imageModel.UpdatedAt,
			DeactivatedAt: imageModel.DeactivatedAt,
		},
		Name: imageModel.Name,
		Type: imageModel.Type,
		Size: imageModel.Size,
	}

	return true, image, nil
}

func (i *ImageRepository) Update(image *entity.Image) error {
	panic("unimplemented")
}
