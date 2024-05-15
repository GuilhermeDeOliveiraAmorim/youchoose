package gorm

import (
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type MovieRepository struct {
	gorm *gorm.DB
}

func NewMovieRepository(gorm *gorm.DB) *MovieRepository {
	return &MovieRepository{
		gorm: gorm,
	}
}

// Create implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) Create(movie *entity.Movie) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) Deactivate(movie *entity.Movie) error {
	panic("unimplemented")
}

// DoTheseMoviesExist implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) DoTheseMoviesExist(movieIDs []string) (bool, []entity.Movie, error) {
	panic("unimplemented")
}

// GetAll implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) GetAll() ([]entity.Movie, error) {
	panic("unimplemented")
}

// GetByActorID implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) GetByActorID(actorID string) ([]entity.Movie, error) {
	panic("unimplemented")
}

// GetByDirectorID implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) GetByDirectorID(directorID string) ([]entity.Movie, error) {
	panic("unimplemented")
}

// GetByGenreID implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) GetByGenreID(genreID string) ([]entity.Movie, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) GetByID(movieID string) (bool, entity.Movie, error) {
	panic("unimplemented")
}

// GetByWriterID implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) GetByWriterID(writerID string) ([]entity.Movie, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.MovieRepositoryInterface.
func (m *MovieRepository) Update(movie *entity.Movie) error {
	panic("unimplemented")
}
