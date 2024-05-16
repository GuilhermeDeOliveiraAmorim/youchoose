package gorm

import (
	"errors"
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

func (m *MovieRepository) Create(movie *entity.Movie) error {
	if err := m.gorm.Create(&Movies{
		ID:            movie.ID,
		Active:        movie.Active,
		CreatedAt:     movie.CreatedAt,
		UpdatedAt:     movie.UpdatedAt,
		DeactivatedAt: movie.DeactivatedAt,
		Title:         movie.Title,
		CountryName:   movie.Nationality.CountryName,
		Flag:          movie.Nationality.Flag,
		ReleaseYear:   movie.ReleaseYear,
		ImageID:       movie.ImageID,
		Votes:         movie.Votes,
	}).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (m *MovieRepository) Deactivate(movie *entity.Movie) error {
	panic("unimplemented")
}

func (m *MovieRepository) DoTheseMoviesExist(movieIDs []string) (bool, []entity.Movie, error) {
	panic("unimplemented")
}

func (m *MovieRepository) GetAll() ([]entity.Movie, error) {
	panic("unimplemented")
}

func (m *MovieRepository) GetByActorID(actorID string) ([]entity.Movie, error) {
	panic("unimplemented")
}

func (m *MovieRepository) GetByDirectorID(directorID string) ([]entity.Movie, error) {
	panic("unimplemented")
}

func (m *MovieRepository) GetByGenreID(genreID string) ([]entity.Movie, error) {
	panic("unimplemented")
}

func (m *MovieRepository) GetByID(movieID string) (bool, entity.Movie, error) {
	panic("unimplemented")
}

func (m *MovieRepository) GetByWriterID(writerID string) ([]entity.Movie, error) {
	panic("unimplemented")
}

func (m *MovieRepository) Update(movie *entity.Movie) error {
	panic("unimplemented")
}
