package gorm

import (
	"errors"
	"fmt"
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type MovieWriterRepository struct {
	gorm *gorm.DB
}

func NewMovieWriterRepository(gorm *gorm.DB) *MovieWriterRepository {
	return &MovieWriterRepository{
		gorm: gorm,
	}
}

func (m *MovieWriterRepository) Create(movieWriter *entity.MovieWriter) error {
	panic("unimplemented")
}

func (m *MovieWriterRepository) CreateMany(movieWriters *[]entity.MovieWriter) error {
	fmt.Println("---")
	fmt.Println(movieWriters)
	fmt.Println("---")
	var movieWritersModel []MovieWriters

	for _, movieWriter := range *movieWriters {
		movieWritersModel = append(movieWritersModel, MovieWriters{
			ID:            movieWriter.ID,
			Active:        movieWriter.Active,
			CreatedAt:     movieWriter.CreatedAt,
			UpdatedAt:     movieWriter.UpdatedAt,
			DeactivatedAt: movieWriter.DeactivatedAt,
			MovieID:       movieWriter.MovieID,
			WriterID:      movieWriter.WriterID,
		})
	}
	
	for _, v := range movieWritersModel {
		fmt.Println(v)
	}

	if err := m.gorm.Create(movieWritersModel).Error; err != nil {
		fmt.Println("OPA")
		return errors.New(err.Error())
	}

	return nil
}

func (m *MovieWriterRepository) Deactivate(movieWriter *entity.MovieWriter) error {
	panic("unimplemented")
}

func (m *MovieWriterRepository) GetAll() ([]entity.MovieWriter, error) {
	panic("unimplemented")
}

func (m *MovieWriterRepository) GetAllByMovieID(movieID string) ([]entity.MovieWriter, error) {
	panic("unimplemented")
}

func (m *MovieWriterRepository) GetAllByWriterID(writerID string) ([]entity.MovieWriter, error) {
	panic("unimplemented")
}

func (m *MovieWriterRepository) GetByID(movieWriterID string) (entity.MovieWriter, error) {
	panic("unimplemented")
}

func (m *MovieWriterRepository) Update(movieWriter *entity.MovieWriter) error {
	panic("unimplemented")
}
