package gorm

import (
	"youchoose/internal/entity"

	"gorm.io/gorm"
)

type GenreRepository struct {
	gorm *gorm.DB
}

// Create implements repositoryinterface.GenreRepositoryInterface.
func (g *GenreRepository) Create(genre *entity.Genre) error {
	panic("unimplemented")
}

// CreateMany implements repositoryinterface.GenreRepositoryInterface.
func (g *GenreRepository) CreateMany(genres *[]entity.Genre) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.GenreRepositoryInterface.
func (g *GenreRepository) Deactivate(genre *entity.Genre) error {
	panic("unimplemented")
}

// DoTheseGenresAreIncludedInTheMovie implements repositoryinterface.GenreRepositoryInterface.
func (g *GenreRepository) DoTheseGenresAreIncludedInTheMovie(movieID string, genresIDs []string) (bool, []entity.Genre, error) {
	panic("unimplemented")
}

// DoTheseGenresExist implements repositoryinterface.GenreRepositoryInterface.
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
			Name:    genreModel.Name,
			ImageID: genreModel.ImageID,
		})
	}

	return true, genres, nil
}

// GetAll implements repositoryinterface.GenreRepositoryInterface.
func (g *GenreRepository) GetAll() ([]entity.Genre, error) {
	panic("unimplemented")
}

// GetAllByMovieID implements repositoryinterface.GenreRepositoryInterface.
func (g *GenreRepository) GetAllByMovieID(movieID string) ([]entity.Genre, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.GenreRepositoryInterface.
func (g *GenreRepository) GetByID(genreID string) (entity.Genre, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.GenreRepositoryInterface.
func (g *GenreRepository) Update(genre *entity.Genre) error {
	panic("unimplemented")
}

func NewGenreRepository(gorm *gorm.DB) *GenreRepository {
	return &GenreRepository{
		gorm: gorm,
	}
}
