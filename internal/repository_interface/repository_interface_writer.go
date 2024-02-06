package repositoryinterface

import "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/entity"

type WriterRepositoryInterface interface {
	Create(writer *entity.Writer) error
	Update(writer *entity.Writer) error
	GetByID(writerID string) (entity.Writer, error)
	GetAll() ([]entity.Writer, error)
	Deactivate(writerID string) error
}
