package repositoryinterface

import "youchoose/internal/entity"

type WriterRepositoryInterface interface {
	Create(writer *entity.Writer) error
	Update(writer *entity.Writer) error
	GetByID(writerID string) (entity.Writer, error)
	GetAll() ([]entity.Writer, error)
	Deactivate(writer *entity.Writer) error
	DoTheseWritersExist(writerIDs []string) (bool, []entity.Writer, error)
}
