package usecase

import (
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeactivateChooserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	deactivateChooserUseCase := NewDeactivateChooserUseCase(mockRepository)

	login, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")

	birthDate, _ := valueobject.NewBirthDate(1, 1, 1990)

	newChooser, _ := entity.NewChooser("Nome 1", login, address, birthDate, uuid.New().String())

	mockRepository.EXPECT().GetByID(newChooser.ID).Return(true, *newChooser, nil)
	mockRepository.EXPECT().Deactivate(newChooser.ID).Return(nil)

	input := DeactivateChooserInputDTO{ID: newChooser.ID}
	output, problemDetails := deactivateChooserUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, newChooser.ID, output.ID)
	assert.Equal(t, "Chooser desativado com sucesso", output.Message)
	assert.True(t, output.IsSuccess)

	mockRepository.EXPECT().GetByID(newChooser.ID).Return(false, entity.Chooser{}, nil)

	input = DeactivateChooserInputDTO{ID: newChooser.ID}
	output, problemDetails = deactivateChooserUseCase.Execute(input)

	assert.NotEmpty(t, problemDetails.ProblemDetails)
	assert.Equal(t, util.RFC404, problemDetails.ProblemDetails[0].Instance)
}
