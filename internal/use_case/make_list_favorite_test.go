package usecase

import (
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMakeListFavoriteUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListRepository := mock.NewMockListRepositoryInterface(ctrl)
	mockChooserRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	mockListFavoriteRepository := mock.NewMockListFavoriteRepositoryInterface(ctrl)

	makeListFavoriteUseCase := NewMakeListFavoriteUseCase(mockListRepository, mockChooserRepository, mockListFavoriteRepository)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)
	list, _ := entity.NewList("Minha Lista", "Descrição da Lista", "Minha Lista", "Descrição da Lista", "chooser123")

	input := MakeListFavoriteInputDTO{
		ChooserID: chooser.ID,
		ListID:    list.ID,
	}

	mockListsFavorites := []entity.ListFavorite{
		{
			ChooserID: uuid.NewString(),
			ListID:    uuid.NewString(),
		},
	}

	mockChooserRepository.EXPECT().GetByID(chooser.ID).Return(true, *chooser, nil)
	mockListRepository.EXPECT().GetByID(list.ID).Return(true, *list, nil)
	mockListFavoriteRepository.EXPECT().GetAllByListID(list.ID).Return(mockListsFavorites, nil)
	mockListFavoriteRepository.EXPECT().Create(gomock.Any()).Return(nil)

	output, problemDetails := makeListFavoriteUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, output.ID, list.ID)
	assert.Equal(t, output.IsSuccess, true)
	assert.Equal(t, output.Message, "Lista acrescentada às favoritas")
}
