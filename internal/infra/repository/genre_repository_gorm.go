package gorm

import (
	"errors"
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type GenreRepository struct {
	gorm *gorm.DB
}

func (g *GenreRepository) Create(genre *entity.Genre) error {
	if err := g.gorm.Create(&Genres{
		ID:            genre.ID,
		Active:        genre.Active,
		CreatedAt:     genre.CreatedAt,
		UpdatedAt:     genre.UpdatedAt,
		DeactivatedAt: genre.DeactivatedAt,
		Name:          genre.Name,
	}).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (g *GenreRepository) CreateMany(genres *[]entity.Genre) error {
	var genresModel []Genres

	for _, genre := range *genres {
		genresModel = append(genresModel, Genres{
			ID:            genre.ID,
			Active:        genre.Active,
			CreatedAt:     genre.CreatedAt,
			UpdatedAt:     genre.UpdatedAt,
			DeactivatedAt: genre.DeactivatedAt,
			Name:          genre.Name,
		})
	}

	if err := g.gorm.Create(genresModel).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (g *GenreRepository) Deactivate(genre *entity.Genre) error {
	panic("unimplemented")
}

func (g *GenreRepository) DoTheseGenresAreIncludedInTheMovie(movieID string, genresIDs []string) (bool, []entity.Genre, error) {
	panic("unimplemented")
}

func (g *GenreRepository) DoTheseGenresExist(genreIDs []string) (bool, []entity.Genre, error) {
	var genresModel []Genres
	result := g.gorm.Where("id IN ?", genreIDs).Find(&genresModel)

	if result.Error != nil {
		return false, nil, result.Error
	}

	var genres []entity.Genre

	if result.RowsAffected != int64(len(genreIDs)) {
		return false, genres, nil
	}

	for _, genreModel := range genresModel {
		genres = append(genres, entity.Genre{
			SharedEntity: entity.SharedEntity{
				ID:            genreModel.ID,
				Active:        genreModel.Active,
				CreatedAt:     genreModel.CreatedAt,
				UpdatedAt:     genreModel.UpdatedAt,
				DeactivatedAt: genreModel.DeactivatedAt,
			},
			Name: genreModel.Name,
		})
	}

	return true, genres, nil
}

func (g *GenreRepository) GetAll() ([]entity.Genre, error) {
	panic("unimplemented")
}

func (g *GenreRepository) GetAllByMovieID(movieID string) ([]entity.Genre, error) {
	panic("unimplemented")
}

func (g *GenreRepository) GetByID(genreID string) (entity.Genre, error) {
	panic("unimplemented")
}

func (g *GenreRepository) Update(genre *entity.Genre) error {
	panic("unimplemented")
}

func NewGenreRepository(gorm *gorm.DB) *GenreRepository {
	return &GenreRepository{
		gorm: gorm,
	}
}
