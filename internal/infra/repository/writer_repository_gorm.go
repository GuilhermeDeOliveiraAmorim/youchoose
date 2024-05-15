package gorm

import (
	"gorm.io/gorm"
	"youchoose/internal/entity"
)

type WriterRepository struct {
	gorm *gorm.DB
}

// Create implements repositoryinterface.WriterRepositoryInterface.
func (w *WriterRepository) Create(writer *entity.Writer) error {
	panic("unimplemented")
}

// CreateMany implements repositoryinterface.WriterRepositoryInterface.
func (w *WriterRepository) CreateMany(writers *[]entity.Writer) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.WriterRepositoryInterface.
func (w *WriterRepository) Deactivate(writer *entity.Writer) error {
	panic("unimplemented")
}

// DoTheseWritersAreIncludedInTheMovie implements repositoryinterface.WriterRepositoryInterface.
func (w *WriterRepository) DoTheseWritersAreIncludedInTheMovie(movieID string, writersIDs []string) (bool, []entity.Writer, error) {
	panic("unimplemented")
}

// DoTheseWritersExist implements repositoryinterface.WriterRepositoryInterface.
func (w *WriterRepository) DoTheseWritersExist(writerIDs []string) (bool, []entity.Writer, error) {
	panic("unimplemented")
}

// GetAll implements repositoryinterface.WriterRepositoryInterface.
func (w *WriterRepository) GetAll() ([]entity.Writer, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.WriterRepositoryInterface.
func (w *WriterRepository) GetByID(writerID string) (entity.Writer, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.WriterRepositoryInterface.
func (w *WriterRepository) Update(writer *entity.Writer) error {
	panic("unimplemented")
}

func NewWriterRepository(gorm *gorm.DB) *WriterRepository {
	return &WriterRepository{
		gorm: gorm,
	}
}
