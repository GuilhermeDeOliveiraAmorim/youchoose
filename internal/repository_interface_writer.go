package internal

type WriterRepositoryInterface interface {
	Create(writer *Writer) error
	Update(writer *Writer) error
	GetByID(writerID string) (Writer, error)
	GetAll() ([]Writer, error)
	Deactivate(writerID string) error
}
