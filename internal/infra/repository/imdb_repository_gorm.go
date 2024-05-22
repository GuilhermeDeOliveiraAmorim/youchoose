package gorm

import (
	"errors"
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type IMDBRepository struct {
	gorm *gorm.DB
}

func NewIMDBRepository(gorm *gorm.DB) *IMDBRepository {
	return &IMDBRepository{
		gorm: gorm,
	}
}

func (i *IMDBRepository) Create(imdb *entity.IMDB) error {
	if err := i.gorm.Create(&IMDBs{
		ID:            imdb.ID,
		Active:        imdb.Active,
		CreatedAt:     imdb.CreatedAt,
		UpdatedAt:     imdb.UpdatedAt,
		DeactivatedAt: imdb.DeactivatedAt,
		IMDBID:        imdb.IMDBID,
	}).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (a *IMDBRepository) CreateMany(imdbs *[]entity.IMDB) error {
	var imdbsModel []IMDBs

	for _, imdb := range *imdbs {
		imdbsModel = append(imdbsModel, IMDBs{
			ID:            imdb.ID,
			Active:        imdb.Active,
			CreatedAt:     imdb.CreatedAt,
			UpdatedAt:     imdb.UpdatedAt,
			DeactivatedAt: imdb.DeactivatedAt,
			IMDBID:        imdb.IMDBID,
		})
	}

	if err := a.gorm.Create(imdbsModel).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (i *IMDBRepository) GetByID(IMDBID string) (bool, error) {
	var imdbModel IMDBs

	result := i.gorm.Model(&IMDBs{}).Where("id = ?", IMDBID).First(&imdbModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (a *IMDBRepository) DoTheseIMDBsExist(IMDBIDs []string) (bool, error) {
	var IMDBsModel []IMDBs
	result := a.gorm.Where("id IN ?", IMDBIDs).Find(&IMDBsModel)

	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected != int64(len(IMDBIDs)) {
		return false, nil
	}

	return true, nil
}
