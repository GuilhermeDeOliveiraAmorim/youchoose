package gorm

import (
	"gorm.io/gorm"
	"youchoose/internal/entity"
)

type MovieDirectorRepository struct {
	gorm *gorm.DB
}

// Create implements repositoryinterface.MovieDirectorRepositoryInterface.
func (m *MovieDirectorRepository) Create(movieDirector *entity.MovieDirector) error {
	panic("unimplemented")
}

// CreateMany implements repositoryinterface.MovieDirectorRepositoryInterface.
func (m *MovieDirectorRepository) CreateMany(movieDirectors *[]entity.MovieDirector) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.MovieDirectorRepositoryInterface.
func (m *MovieDirectorRepository) Deactivate(movieDirector *entity.MovieDirector) error {
	panic("unimplemented")
}

// GetAll implements repositoryinterface.MovieDirectorRepositoryInterface.
func (m *MovieDirectorRepository) GetAll() ([]entity.MovieDirector, error) {
	panic("unimplemented")
}

// GetAllByDirectorID implements repositoryinterface.MovieDirectorRepositoryInterface.
func (m *MovieDirectorRepository) GetAllByDirectorID(directorID string) ([]entity.MovieDirector, error) {
	panic("unimplemented")
}

// GetAllByMovieID implements repositoryinterface.MovieDirectorRepositoryInterface.
func (m *MovieDirectorRepository) GetAllByMovieID(movieID string) ([]entity.MovieDirector, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.MovieDirectorRepositoryInterface.
func (m *MovieDirectorRepository) GetByID(movieDirectorID string) (entity.MovieDirector, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.MovieDirectorRepositoryInterface.
func (m *MovieDirectorRepository) Update(movieDirector *entity.MovieDirector) error {
	panic("unimplemented")
}

func NewMovieDirectorRepository(gorm *gorm.DB) *MovieDirectorRepository {
	return &MovieDirectorRepository{
		gorm: gorm,
	}
}
