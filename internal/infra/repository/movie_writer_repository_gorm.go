package gorm

import (
	"gorm.io/gorm"
	"youchoose/internal/entity"
)

type MovieWriterRepository struct {
	gorm *gorm.DB
}

// Create implements repositoryinterface.MovieWriterRepositoryInterface.
func (m *MovieWriterRepository) Create(movieWriter *entity.MovieWriter) error {
	panic("unimplemented")
}

// CreateMany implements repositoryinterface.MovieWriterRepositoryInterface.
func (m *MovieWriterRepository) CreateMany(movieWriters *[]entity.MovieWriter) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.MovieWriterRepositoryInterface.
func (m *MovieWriterRepository) Deactivate(movieWriter *entity.MovieWriter) error {
	panic("unimplemented")
}

// GetAll implements repositoryinterface.MovieWriterRepositoryInterface.
func (m *MovieWriterRepository) GetAll() ([]entity.MovieWriter, error) {
	panic("unimplemented")
}

// GetAllByMovieID implements repositoryinterface.MovieWriterRepositoryInterface.
func (m *MovieWriterRepository) GetAllByMovieID(movieID string) ([]entity.MovieWriter, error) {
	panic("unimplemented")
}

// GetAllByWriterID implements repositoryinterface.MovieWriterRepositoryInterface.
func (m *MovieWriterRepository) GetAllByWriterID(writerID string) ([]entity.MovieWriter, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.MovieWriterRepositoryInterface.
func (m *MovieWriterRepository) GetByID(movieWriterID string) (entity.MovieWriter, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.MovieWriterRepositoryInterface.
func (m *MovieWriterRepository) Update(movieWriter *entity.MovieWriter) error {
	panic("unimplemented")
}

func NewMovieWriterRepository(gorm *gorm.DB) *MovieWriterRepository {
	return &MovieWriterRepository{
		gorm: gorm,
	}
}
