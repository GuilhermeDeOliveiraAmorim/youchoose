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

func TestGetChoosersUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	getChoosersUseCase := NewGetChoosersUseCase(mockRepository)

	login1, _ := valueobject.NewLogin("email@email.com", "12@#asd89")
	login2, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address1, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")
	address2, _ := valueobject.NewAddress("Maceió", "Alagoas", "Brasil")

	birthDate1, _ := valueobject.NewBirthDate(1, 1, 1990)
	birthDate2, _ := valueobject.NewBirthDate(2, 2, 1990)

	chooser1, _ := entity.NewChooser("Nome 1", login1, address1, birthDate1, uuid.New().String())
	chooser2, _ := entity.NewChooser("Nome 2", login2, address2, birthDate2, uuid.New().String())

	expectedChoosers := []entity.Chooser{*chooser1, *chooser2}

	mockRepository.EXPECT().GetAll().Return(expectedChoosers, nil)

	input := GetChoosersInputDTO{}
	output, problemDetails := getChoosersUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, len(expectedChoosers), len(output.Choosers))
	for i, chooser := range output.Choosers {
		assert.Equal(t, expectedChoosers[i].ID, chooser.ID)
		assert.Equal(t, expectedChoosers[i].Name, chooser.Name)
		assert.Equal(t, expectedChoosers[i].Address.City, chooser.City)
		assert.Equal(t, expectedChoosers[i].Address.State, chooser.State)
		assert.Equal(t, expectedChoosers[i].Address.Country, chooser.Country)
		assert.Equal(t, expectedChoosers[i].BirthDate.Day, chooser.Day)
		assert.Equal(t, expectedChoosers[i].BirthDate.Month, chooser.Month)
		assert.Equal(t, expectedChoosers[i].BirthDate.Year, chooser.Year)
		assert.Equal(t, expectedChoosers[i].ImageID, chooser.ImageID)
	}

	mockRepository.EXPECT().GetAll().Return(nil, errors.New("database error"))

	output, problemDetails = getChoosersUseCase.Execute(input)

	assert.Empty(t, output.Choosers)
	assert.NotEmpty(t, problemDetails.ProblemDetails)
	assert.Equal(t, util.RFC503, problemDetails.ProblemDetails[0].Instance)
}
