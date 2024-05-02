package usecase

import (
	"fmt"
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRemoveListFavoriteUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	mockListRepository := mock.NewMockListRepositoryInterface(ctrl)
	mockListFavoriteRepository := mock.NewMockListFavoriteRepositoryInterface(ctrl)

	removeListFavoriteUseCase := NewRemoveListFavoriteUseCase(mockChooserRepository, mockListRepository, mockListFavoriteRepository)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "Aracaju", State: "SE", Country: "Brasil"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)
	list, _ := entity.NewList("Minha Lista", "Descrição da Lista", "Minha Lista", "Descrição da Lista", "chooser123")

	listFavorite := entity.NewListFavorite(chooser.ID, list.ID)

	input := RemoveListFavoriteInputDTO{
		ChooserID: chooser.ID,
		ListID:    list.ID,
	}

	mockChooserRepository.EXPECT().GetByID(chooser.ID).Return(true, *chooser, nil)
	mockListRepository.EXPECT().GetByID(list.ID).Return(true, *list, nil)
	mockListFavoriteRepository.EXPECT().GetByChooserIDAndListID(chooser.ID, list.ID).Return(true, *listFavorite, nil)
	mockListFavoriteRepository.EXPECT().Deactivate(gomock.Any()).Return(nil)
	mockListRepository.EXPECT().Update(gomock.Any()).Return(nil)

	output, problemDetails := removeListFavoriteUseCase.Execute(input)

	fmt.Println(output, problemDetails)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, output.ID, listFavorite.ID)
	assert.Equal(t, output.IsSuccess, true)
	assert.Equal(t, output.Message, "Lista removida das favoritas")
}
