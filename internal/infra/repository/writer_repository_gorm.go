package gorm

import (
	"errors"
	"youchoose/internal/entity"
	valueobject "youchoose/internal/value_object"

	"gorm.io/gorm"
)

type WriterRepository struct {
	gorm *gorm.DB
}

func NewWriterRepository(gorm *gorm.DB) *WriterRepository {
	return &WriterRepository{
		gorm: gorm,
	}
}

func (w *WriterRepository) Create(writer *entity.Writer) error {
	if err := w.gorm.Create(&Writers{
		ID:            writer.ID,
		Active:        writer.Active,
		CreatedAt:     writer.CreatedAt,
		UpdatedAt:     writer.UpdatedAt,
		DeactivatedAt: writer.DeactivatedAt,
		Name:          writer.Name,
		Day:           writer.BirthDate.Day,
		Month:         writer.BirthDate.Month,
		Year:          writer.BirthDate.Year,
		CountryName:   writer.Nationality.CountryName,
		Flag:          writer.Nationality.Flag,
		ImageID:       writer.ImageID,
	}).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (w *WriterRepository) CreateMany(writers *[]entity.Writer) error {
	var writersModel []Writers

	for _, writer := range *writers {
		writersModel = append(writersModel, Writers{
			ID:            writer.ID,
			Active:        writer.Active,
			CreatedAt:     writer.CreatedAt,
			UpdatedAt:     writer.UpdatedAt,
			DeactivatedAt: writer.DeactivatedAt,
			Name:          writer.Name,
			Day:           writer.BirthDate.Day,
			Month:         writer.BirthDate.Month,
			Year:          writer.BirthDate.Year,
			CountryName:   writer.Nationality.CountryName,
			Flag:          writer.Nationality.Flag,
			ImageID:       writer.ImageID,
		})
	}

	if err := w.gorm.Create(writersModel).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (w *WriterRepository) Deactivate(writer *entity.Writer) error {
	panic("unimplemented")
}

func (w *WriterRepository) DoTheseWritersAreIncludedInTheMovie(movieID string, writersIDs []string) (bool, []entity.Writer, error) {
	panic("unimplemented")
}

func (w *WriterRepository) DoTheseWritersExist(writerIDs []string) (bool, []entity.Writer, error) {
	var writersModel []Writers
	result := w.gorm.Where("id IN ?", writerIDs).Find(&writersModel)

	if result.Error != nil {
		return false, nil, result.Error
	}

	var writers []entity.Writer

	if result.RowsAffected != int64(len(writerIDs)) {
		return false, writers, nil
	}

	for _, writerModel := range writersModel {
		writers = append(writers, entity.Writer{
			SharedEntity: entity.SharedEntity{
				ID:            writerModel.ID,
				Active:        writerModel.Active,
				CreatedAt:     writerModel.CreatedAt,
				UpdatedAt:     writerModel.UpdatedAt,
				DeactivatedAt: writerModel.DeactivatedAt,
			},
			Name:    writerModel.Name,
			ImageID: writerModel.ImageID,
			BirthDate: &valueobject.BirthDate{
				Day:   writerModel.Day,
				Month: writerModel.Month,
				Year:  writerModel.Year,
			},
			Nationality: &valueobject.Nationality{
				CountryName: writerModel.CountryName,
				Flag:        writerModel.Flag,
			},
		})
	}

	return true, writers, nil
}

func (w *WriterRepository) GetAll() ([]entity.Writer, error) {
	panic("unimplemented")
}

func (w *WriterRepository) GetByID(writerID string) (entity.Writer, error) {
	panic("unimplemented")
}

func (w *WriterRepository) Update(writer *entity.Writer) error {
	panic("unimplemented")
}
