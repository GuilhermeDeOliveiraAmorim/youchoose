package gorm

import (
	"errors"
	"youchoose/internal/entity"
	valueobject "youchoose/internal/value_object"

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

func (cr *ChooserRepository) Deactivate(chooser *entity.Chooser) error {
	panic("unimplemented")
}

func (cr *ChooserRepository) GetAll() ([]entity.Chooser, error) {
	var choosersModel []Choosers
	if err := cr.gorm.Find(&choosersModel).Error; err != nil {
		return nil, err
	}

	var choosers []entity.Chooser

	if len(choosersModel) > 0 {
		for _, chooserModel := range choosersModel {
			chooser := entity.Chooser{
				SharedEntity: entity.SharedEntity{
					ID:            chooserModel.ID,
					Active:        chooserModel.Active,
					CreatedAt:     chooserModel.CreatedAt,
					UpdatedAt:     chooserModel.UpdatedAt,
					DeactivatedAt: chooserModel.DeactivatedAt,
				},
				Name: chooserModel.Name,
				Login: &valueobject.Login{
					Email:    chooserModel.Email,
					Password: chooserModel.Password,
				},
				Address: &valueobject.Address{
					City:    chooserModel.City,
					State:   chooserModel.State,
					Country: chooserModel.Country,
				},
				BirthDate: &valueobject.BirthDate{
					Day:   chooserModel.Day,
					Month: chooserModel.Month,
					Year:  chooserModel.Year,
				},
				ImageID: chooserModel.ImageID,
			}

			choosers = append(choosers, chooser)
		}
	}

	return choosers, nil
}

func (cr *ChooserRepository) GetByEmail(chooserEmail string) (entity.Chooser, error) {
	panic("unimplemented")
}

func (cr *ChooserRepository) GetByID(chooserID string) (bool, entity.Chooser, error) {
	var chooserModel Choosers
	result := cr.gorm.Model(&Choosers{}).Where("id = ?", chooserID).First(&chooserModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, entity.Chooser{}, nil
		}
		return false, entity.Chooser{}, result.Error
	}

	chooser := entity.Chooser{
		SharedEntity: entity.SharedEntity{
			ID:            chooserModel.ID,
			Active:        chooserModel.Active,
			CreatedAt:     chooserModel.CreatedAt,
			UpdatedAt:     chooserModel.UpdatedAt,
			DeactivatedAt: chooserModel.DeactivatedAt,
		},
		Name: chooserModel.Name,
		Login: &valueobject.Login{
			Email:    chooserModel.Email,
			Password: chooserModel.Password,
		},
		Address: &valueobject.Address{
			City:    chooserModel.City,
			State:   chooserModel.State,
			Country: chooserModel.Country,
		},
		BirthDate: &valueobject.BirthDate{
			Day:   chooserModel.Day,
			Month: chooserModel.Month,
			Year:  chooserModel.Year,
		},
		ImageID: chooserModel.ImageID,
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
