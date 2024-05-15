package gorm

import (
	"gorm.io/gorm"
	"youchoose/internal/entity"
)

type MovieGenreRepository struct {
	gorm *gorm.DB
}

// Create implements repositoryinterface.MovieGenreRepositoryInterface.
func (m *MovieGenreRepository) Create(movieGenre *entity.MovieGenre) error {
	panic("unimplemented")
}

// CreateMany implements repositoryinterface.MovieGenreRepositoryInterface.
func (m *MovieGenreRepository) CreateMany(movieGenres *[]entity.MovieGenre) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.MovieGenreRepositoryInterface.
func (m *MovieGenreRepository) Deactivate(movieGenre *entity.MovieGenre) error {
	panic("unimplemented")
}

// GetAll implements repositoryinterface.MovieGenreRepositoryInterface.
func (m *MovieGenreRepository) GetAll() ([]entity.MovieGenre, error) {
	panic("unimplemented")
}

// GetAllByGenreID implements repositoryinterface.MovieGenreRepositoryInterface.
func (m *MovieGenreRepository) GetAllByGenreID(genreID string) ([]entity.MovieGenre, error) {
	panic("unimplemented")
}

// GetAllByMovieID implements repositoryinterface.MovieGenreRepositoryInterface.
func (m *MovieGenreRepository) GetAllByMovieID(movieID string) ([]entity.MovieGenre, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.MovieGenreRepositoryInterface.
func (m *MovieGenreRepository) GetByID(movieGenreID string) (entity.MovieGenre, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.MovieGenreRepositoryInterface.
func (m *MovieGenreRepository) Update(movieGenre *entity.MovieGenre) error {
	panic("unimplemented")
}

func NewMovieGenreRepository(gorm *gorm.DB) *MovieGenreRepository {
	return &MovieGenreRepository{
		gorm: gorm,
	}
}
