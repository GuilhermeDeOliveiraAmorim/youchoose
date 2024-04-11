package usecase

import (
	"errors"
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetListsUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockRepositoryRepo := mock.NewMockListRepositoryInterface(ctrl)
	getListsUseCase := NewGetListsUseCase(mockChooserRepo, mockRepositoryRepo)

	list1, _ := entity.NewList("Minha Lista 1", "Descrição da Lista 1", uuid.NewString(), uuid.NewString(), uuid.NewString())
	list2, _ := entity.NewList("Minha Lista 2", "Descrição da Lista 2", uuid.NewString(), uuid.NewString(), uuid.NewString())

	expectedLists := []entity.List{*list1, *list2}

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "Aracaju", State: "Sergipe", Country: "Brazil"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	mockChooserRepo.EXPECT().GetByID(gomock.Any()).Return(true, *chooser, nil).Times(2)
	mockRepositoryRepo.EXPECT().GetAll().Return(expectedLists, nil)

	input := GetListsInputDTO{}
	output, problemDetails := getListsUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, len(expectedLists), len(output.Lists))

	for i, list := range output.Lists {
		assert.Equal(t, expectedLists[i].ID, list.ChooserID)
		assert.Equal(t, expectedLists[i].Title, list.Title)
		assert.Equal(t, expectedLists[i].Description, list.Description)
		assert.Equal(t, expectedLists[i].ChooserID, list.ChooserID)
		assert.Equal(t, expectedLists[i].ProfileImageID, list.ProfileImageID)
		assert.Equal(t, expectedLists[i].CoverImageID, list.CoverImageID)
		assert.Equal(t, expectedLists[i].Votes, list.Votes)
	}

	mockRepositoryRepo.EXPECT().GetAll().Return(nil, errors.New("database error"))

	output, problemDetails = getListsUseCase.Execute(input)

	assert.Empty(t, output.Lists)
	assert.NotEmpty(t, problemDetails.ProblemDetails)
	assert.Equal(t, util.RFC503, problemDetails.ProblemDetails[0].Instance)
}
