package gorm

import (
	"gorm.io/gorm"
	"youchoose/internal/entity"
)

type DirectorRepository struct {
	gorm *gorm.DB
}

// Create implements repositoryinterface.DirectorRepositoryInterface.
func (d *DirectorRepository) Create(director *entity.Director) error {
	panic("unimplemented")
}

// CreateMany implements repositoryinterface.DirectorRepositoryInterface.
func (d *DirectorRepository) CreateMany(directors *[]entity.Director) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.DirectorRepositoryInterface.
func (d *DirectorRepository) Deactivate(director *entity.Director) error {
	panic("unimplemented")
}

// DoTheseDirectorsAreIncludedInTheMovie implements repositoryinterface.DirectorRepositoryInterface.
func (d *DirectorRepository) DoTheseDirectorsAreIncludedInTheMovie(movieID string, directorsIDs []string) (bool, []entity.Director, error) {
	panic("unimplemented")
}

// DoTheseDirectorsExist implements repositoryinterface.DirectorRepositoryInterface.
func (d *DirectorRepository) DoTheseDirectorsExist(directorIDs []string) (bool, []entity.Director, error) {
	panic("unimplemented")
}

// GetAll implements repositoryinterface.DirectorRepositoryInterface.
func (d *DirectorRepository) GetAll() ([]entity.Director, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.DirectorRepositoryInterface.
func (d *DirectorRepository) GetByID(directorID string) (entity.Director, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.DirectorRepositoryInterface.
func (d *DirectorRepository) Update(director *entity.Director) error {
	panic("unimplemented")
}

func NewDirectorRepository(gorm *gorm.DB) *DirectorRepository {
	return &DirectorRepository{
		gorm: gorm,
	}
}
