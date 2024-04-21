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

func TestRemoveListFavoriteUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	mockListRepository := mock.NewMockListRepositoryInterface(ctrl)
	mockListFavoriteRepository := mock.NewMockListFavoriteRepositoryInterface(ctrl)

	removeListFavoriteUseCase := NewRemoveListFavoriteUseCase(mockChooserRepository, mockListRepository, mockListFavoriteRepository)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	listFavorite := entity.NewListFavorite("chooser1", "list1")

	chooserID := chooser.ID
	listFavoriteID := listFavorite.ID

	input := RemoveListFavoriteInputDTO{
		ChooserID:      chooserID,
		ListFavoriteID: listFavoriteID,
	}

	mockListFavoriteRepository.EXPECT().GetByID(listFavoriteID).Return(true, *listFavorite, nil)
	mockChooserRepository.EXPECT().GetByID(chooserID).Return(true, *chooser, nil)
	mockListFavoriteRepository.EXPECT().Deactivate(gomock.Any()).Return(nil)

	output, problemDetails := removeListFavoriteUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, output.ID, listFavoriteID)
	assert.Equal(t, output.IsSuccess, true)
	assert.Equal(t, output.Message, "Lista removida das favoritas")
}
