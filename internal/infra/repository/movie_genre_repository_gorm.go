package gorm

import (
	"errors"
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type MovieGenreRepository struct {
	gorm *gorm.DB
}

func NewMovieGenreRepository(gorm *gorm.DB) *MovieGenreRepository {
	return &MovieGenreRepository{
		gorm: gorm,
	}
}

func (m *MovieGenreRepository) Create(movieGenre *entity.MovieGenre) error {
	panic("unimplemented")
}

func (m *MovieGenreRepository) CreateMany(movieGenres *[]entity.MovieGenre) error {
	var movieGenresModel []MovieGenres

	for _, movieGenre := range *movieGenres {
		movieGenresModel = append(movieGenresModel, MovieGenres{
			ID:            movieGenre.ID,
			Active:        movieGenre.Active,
			CreatedAt:     movieGenre.CreatedAt,
			UpdatedAt:     movieGenre.UpdatedAt,
			DeactivatedAt: movieGenre.DeactivatedAt,
			MovieID:       movieGenre.MovieID,
			GenreID:       movieGenre.GenreID,
		})
	}

	if err := m.gorm.Create(movieGenresModel).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (m *MovieGenreRepository) Deactivate(movieGenre *entity.MovieGenre) error {
	panic("unimplemented")
}

func (m *MovieGenreRepository) GetAll() ([]entity.MovieGenre, error) {
	panic("unimplemented")
}

func (m *MovieGenreRepository) GetAllByGenreID(genreID string) ([]entity.MovieGenre, error) {
	panic("unimplemented")
}

func (m *MovieGenreRepository) GetAllByMovieID(movieID string) ([]entity.MovieGenre, error) {
	panic("unimplemented")
}

func (m *MovieGenreRepository) GetByID(movieGenreID string) (entity.MovieGenre, error) {
	panic("unimplemented")
}

func (m *MovieGenreRepository) Update(movieGenre *entity.MovieGenre) error {
	panic("unimplemented")
}
