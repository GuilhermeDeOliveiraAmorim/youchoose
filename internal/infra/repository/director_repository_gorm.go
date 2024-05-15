package gorm

import (
	"errors"
	"youchoose/internal/entity"
	valueobject "youchoose/internal/value_object"

	"gorm.io/gorm"
)

type DirectorRepository struct {
	gorm *gorm.DB
}

func NewDirectorRepository(gorm *gorm.DB) *DirectorRepository {
	return &DirectorRepository{
		gorm: gorm,
	}
}

func (d *DirectorRepository) Create(director *entity.Director) error {
	if err := d.gorm.Create(&Directors{
		ID:            director.ID,
		Active:        director.Active,
		CreatedAt:     director.CreatedAt,
		UpdatedAt:     director.UpdatedAt,
		DeactivatedAt: director.DeactivatedAt,
		Name:          director.Name,
		Day:           director.BirthDate.Day,
		Month:         director.BirthDate.Month,
		Year:          director.BirthDate.Year,
		CountryName:   director.Nationality.CountryName,
		Flag:          director.Nationality.Flag,
		ImageID:       director.ImageID,
	}).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (d *DirectorRepository) CreateMany(directors *[]entity.Director) error {
	panic("unimplemented")
}

func (d *DirectorRepository) Deactivate(director *entity.Director) error {
	panic("unimplemented")
}

func (d *DirectorRepository) DoTheseDirectorsAreIncludedInTheMovie(movieID string, directorsIDs []string) (bool, []entity.Director, error) {
	panic("unimplemented")
}

func (d *DirectorRepository) DoTheseDirectorsExist(directorIDs []string) (bool, []entity.Director, error) {
	var directorsModel []Directors
	result := d.gorm.Where("id IN ?", directorIDs).Find(&directorsModel)

	if result.Error != nil {
		return false, nil, result.Error
	}

	var directors []entity.Director

	if result.RowsAffected != int64(len(directorIDs)) {
		return false, directors, nil
	}

	for _, directorModel := range directorsModel {
		directors = append(directors, entity.Director{
			SharedEntity: entity.SharedEntity{
				ID:            directorModel.ID,
				Active:        directorModel.Active,
				CreatedAt:     directorModel.CreatedAt,
				UpdatedAt:     directorModel.UpdatedAt,
				DeactivatedAt: directorModel.DeactivatedAt,
			},
			Name:    directorModel.Name,
			ImageID: directorModel.ImageID,
			BirthDate: &valueobject.BirthDate{
				Day:   directorModel.Day,
				Month: directorModel.Month,
				Year:  directorModel.Year,
			},
			Nationality: &valueobject.Nationality{
				CountryName: directorModel.CountryName,
				Flag:        directorModel.Flag,
			},
		})
	}

	return true, directors, nil
}

func (d *DirectorRepository) GetAll() ([]entity.Director, error) {
	panic("unimplemented")
}

func (d *DirectorRepository) GetByID(directorID string) (entity.Director, error) {
	panic("unimplemented")
}

func (d *DirectorRepository) Update(director *entity.Director) error {
	panic("unimplemented")
}
