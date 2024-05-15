package gorm

import (
	"errors"
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type MovieDirectorRepository struct {
	gorm *gorm.DB
}

func NewMovieDirectorRepository(gorm *gorm.DB) *MovieDirectorRepository {
	return &MovieDirectorRepository{
		gorm: gorm,
	}
}

func (m *MovieDirectorRepository) Create(movieDirector *entity.MovieDirector) error {
	panic("unimplemented")
}

func (m *MovieDirectorRepository) CreateMany(movieDirectors *[]entity.MovieDirector) error {
	var movieDirectorsModel []MovieDirectors

	for _, movieDirector := range *movieDirectors {
		movieDirectorsModel = append(movieDirectorsModel, MovieDirectors{
			ID:            movieDirector.ID,
			Active:        movieDirector.Active,
			CreatedAt:     movieDirector.CreatedAt,
			UpdatedAt:     movieDirector.UpdatedAt,
			DeactivatedAt: movieDirector.DeactivatedAt,
			MovieID:       movieDirector.MovieID,
			DirectorID:    movieDirector.DirectorID,
		})
	}

	if err := m.gorm.Create(movieDirectorsModel).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (m *MovieDirectorRepository) Deactivate(movieDirector *entity.MovieDirector) error {
	panic("unimplemented")
}

func (m *MovieDirectorRepository) GetAll() ([]entity.MovieDirector, error) {
	panic("unimplemented")
}

func (m *MovieDirectorRepository) GetAllByDirectorID(directorID string) ([]entity.MovieDirector, error) {
	panic("unimplemented")
}

func (m *MovieDirectorRepository) GetAllByMovieID(movieID string) ([]entity.MovieDirector, error) {
	panic("unimplemented")
}

func (m *MovieDirectorRepository) GetByID(movieDirectorID string) (entity.MovieDirector, error) {
	panic("unimplemented")
}

func (m *MovieDirectorRepository) Update(movieDirector *entity.MovieDirector) error {
	panic("unimplemented")
}
