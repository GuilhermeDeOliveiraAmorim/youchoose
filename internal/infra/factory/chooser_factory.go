package factory

import (
	repository "youchoose/internal/infra/repository"
	usecase "youchoose/internal/use_case"

	"gorm.io/gorm"
)

type ChooserFactory struct {
	CreateChooser     *usecase.CreateChooserUseCase
	FindChooserByID   *usecase.GetChooserUseCase
	GetChoosers       *usecase.GetChoosersUseCase
	UpdateChooser     *usecase.UpdateChooserUseCase
	DeactivateChooser *usecase.DeactivateChooserUseCase
}

func NewChooserFactory(db *gorm.DB) *ChooserFactory {
	chooserRepository := repository.NewChooserRepository(db)
	imageRepository := repository.NewImageRepository(db)

	createChooser := usecase.NewCreateChooserUseCase(chooserRepository, imageRepository)
	findChooserByID := usecase.NewGetChooserUseCase(chooserRepository)
	getChoosers := usecase.NewGetChoosersUseCase(chooserRepository)
	updateChooser := usecase.NewUpdateChooserUseCase(chooserRepository, imageRepository)
	deactivateChooser := usecase.NewDeactivateChooserUseCase(chooserRepository)

	return &ChooserFactory{
		CreateChooser:     createChooser,
		FindChooserByID:   findChooserByID,
		GetChoosers:       getChoosers,
		UpdateChooser:     updateChooser,
		DeactivateChooser: deactivateChooser,
	}
}
