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

func TestGetChooserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	getChooserUseCase := NewGetChooserUseCase(mockRepository)

	login1, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address1, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")

	birthDate1, _ := valueobject.NewBirthDate(1, 1, 1990)

	newChooser, _ := entity.NewChooser("Nome 1", login1, address1, birthDate1, uuid.New().String())

	outputExpected := GetChooserOutputDTO {
		ID: newChooser.ID,
		Name: newChooser.Name,
        City: newChooser.Address.City,
        State: newChooser.Address.State,
        Country: newChooser.Address.Country,
        Day: newChooser.BirthDate.Day,
		Month: newChooser.BirthDate.Month,
        Year: newChooser.BirthDate.Year,
        ImageID: newChooser.ImageID,
	}

	mockRepository.EXPECT().GetByID(newChooser.ID).Return(true, *newChooser, nil)

	input := GetChooserInputDTO{
		ID: newChooser.ID,
	}

	output, problemDetails := getChooserUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, output, outputExpected)

	mockRepository.EXPECT().GetByID(newChooser.ID).Return(false, entity.Chooser{}, errors.New("database error"))

	output, problemDetails = getChooserUseCase.Execute(input)

	assert.Empty(t, output)
	assert.NotEmpty(t, problemDetails.ProblemDetails)
	assert.Equal(t, util.RFC503, problemDetails.ProblemDetails[0].Instance)
}
