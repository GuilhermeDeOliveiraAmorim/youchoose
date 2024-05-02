package gorm

import (
	"errors"
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type ChooserRepository struct {
	gorm *gorm.DB
}

func NewChooserRepository(gorm *gorm.DB) *ChooserRepository {
	return &ChooserRepository{
		gorm: gorm,
	}
}

func (cr *ChooserRepository) Create(chooser *entity.Chooser) error {
	if err := cr.gorm.Create(&ChooserModel{
		ID:            chooser.ID,
		Active:        chooser.Active,
		CreatedAt:     chooser.CreatedAt,
		UpdatedAt:     chooser.UpdatedAt,
		DeactivatedAt: chooser.DeactivatedAt,
		Name:          chooser.Name,
		Email:         chooser.Login.Email,
		Password:      chooser.Login.Password,
		City:          chooser.Address.City,
		State:         chooser.Address.State,
		Country:       chooser.Address.Country,
		Day:           chooser.BirthDate.Day,
		Month:         chooser.BirthDate.Month,
		Year:          chooser.BirthDate.Year,
		ImageID:       chooser.ImageID,
	}).Error; err != nil {
		return errors.New("failed to create chooser: " + err.Error())
	}

	return nil
}
