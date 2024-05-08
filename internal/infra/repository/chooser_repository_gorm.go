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
	if err := cr.gorm.Create(&Choosers{
		ID:            chooser.ID,
		Active:        chooser.Active,
		CreatedAt:     chooser.CreatedAt,
		UpdatedAt:     chooser.UpdatedAt,
		DeactivatedAt: chooser.DeactivatedAt,
		Name:          chooser.Name,
		Email:         chooser.Login.Email,
		EmailSalt:     chooser.Login.EmailSalt,
		Password:      chooser.Login.Password,
		PasswordSalt:  chooser.Login.PasswordSalt,
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

func (cr *ChooserRepository) ChooserAlreadyExists(chooserEmail string) (bool, error) {
	var count int64
	if err := cr.gorm.Model(&Choosers{}).Where("email = ?", chooserEmail).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (cr *ChooserRepository) Deactivate(chooser *entity.Chooser) error {
	panic("unimplemented")
}

func (cr *ChooserRepository) GetAll() ([]entity.Chooser, error) {
	panic("unimplemented")
}

func (cr *ChooserRepository) GetByEmail(chooserEmail string) (entity.Chooser, error) {
	panic("unimplemented")
}

func (cr *ChooserRepository) GetByID(chooserID string) (bool, entity.Chooser, error) {
	var chooser entity.Chooser
	result := cr.gorm.Model(&Choosers{}).Where("id = ?", chooserID).First(&chooser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, entity.Chooser{}, nil
		}
		return false, entity.Chooser{}, result.Error
	}
	return true, chooser, nil
}

func (cr *ChooserRepository) GetVotation(chooserID string, listID string) (entity.Votation, error) {
	panic("unimplemented")
}

func (cr *ChooserRepository) GetVotations(chooserID string) ([]entity.Votation, error) {
	panic("unimplemented")
}

func (cr *ChooserRepository) Update(chooser *entity.Chooser) error {
	panic("unimplemented")
}
