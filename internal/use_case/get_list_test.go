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

func TestGetListUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockListRepo := mock.NewMockListRepositoryInterface(ctrl)
	getListUseCase := NewGetListUseCase(mockChooserRepo, mockListRepo)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	list, _ := entity.NewList("Minha Lista 1", "Descrição da Lista 1", uuid.NewString(), uuid.NewString(), chooser.ID)

	input := GetListInputDTO{
		ChooserID: chooser.ID,
		ListID:    list.ID,
	}

	outputExpected := ListOutputDTO{
		ListID:         list.ID,
		CreatedAt:      list.CreatedAt,
		Title:          list.Title,
		Description:    list.Description,
		ChooserID:      list.ChooserID,
		ProfileImageID: list.ProfileImageID,
		CoverImageID:   list.CoverImageID,
		Votes:          list.Votes,
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *chooser, nil)
	mockListRepo.EXPECT().GetByID(input.ListID).Return(true, *list, nil)

	output, problems := getListUseCase.Execute(input)

	assert.Empty(t, problems.ProblemDetails)
	assert.Equal(t, outputExpected.ListID, output.ListID)
	assert.Equal(t, outputExpected.CreatedAt, output.CreatedAt)
	assert.Equal(t, outputExpected.Title, output.Title)
	assert.Equal(t, outputExpected.Description, output.Description)
	assert.Equal(t, outputExpected.ChooserID, output.ChooserID)
	assert.Equal(t, outputExpected.ProfileImageID, output.ProfileImageID)
	assert.Equal(t, outputExpected.CoverImageID, output.CoverImageID)
	assert.Equal(t, outputExpected.Votes, output.Votes)
}
