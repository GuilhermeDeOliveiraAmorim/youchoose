package factory

import (
	repository "youchoose/internal/infra/repository"
	usecase "youchoose/internal/use_case"

	"gorm.io/gorm"
)

type ChooserFactory struct {
	CreateChooser *usecase.CreateChooserUseCase
}

func NewChooserFactory(db *gorm.DB) *ChooserFactory {
	chooserRepository := repository.NewChooserRepository(db)

	createChooser := usecase.NewCreateChooserUseCase(chooserRepository)

	return &ChooserFactory{
		CreateChooser: createChooser,
	}
}
