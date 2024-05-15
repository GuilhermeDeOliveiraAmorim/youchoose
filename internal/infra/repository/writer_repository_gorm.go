package gorm

import (
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
	panic("unimplemented")
}

func (w *WriterRepository) CreateMany(writers *[]entity.Writer) error {
	panic("unimplemented")
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
