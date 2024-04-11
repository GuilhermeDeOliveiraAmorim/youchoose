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

func TestGetChooserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	getChooserUseCase := NewGetChooserUseCase(mockRepository)

	login, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")

	birthDate, _ := valueobject.NewBirthDate(1, 1, 1990)

	chooser_1, _ := entity.NewChooser("Nome 1", login, address, birthDate, uuid.New().String())
	chooser_2, _ := entity.NewChooser("Nome 2", login, address, birthDate, uuid.New().String())

	outputExpected := ChooserOutputDTO{
		ID:        chooser_2.ID,
		CreatedAt: chooser_2.CreatedAt,
		Name:      chooser_2.Name,
		City:      chooser_2.Address.City,
		State:     chooser_2.Address.State,
		Country:   chooser_2.Address.Country,
		Day:       chooser_2.BirthDate.Day,
		Month:     chooser_2.BirthDate.Month,
		Year:      chooser_2.BirthDate.Year,
		ImageID:   chooser_2.ImageID,
	}

	mockRepository.EXPECT().GetByID(gomock.Any()).Return(true, *chooser_1, nil)
	mockRepository.EXPECT().GetByID(gomock.Any()).Return(true, *chooser_2, nil)

	input := GetChooserInputDTO{
		ChooserID: chooser_2.ID,
	}

	output, problemDetails := getChooserUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, output, outputExpected)
}
