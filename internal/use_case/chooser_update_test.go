package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateChooserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	updateChooserUseCase := NewUpdateChooserUseCase(mockRepository)

	imageID := uuid.New().String()

	createChooserInputDTO := CreateChooserInputDTO{
		Name:    "John Doe",
		City:    "Aracaju",
		State:   "Sergipe",
		Country: "Brasil",
		Day:     10,
		Month:   5,
		Year:    1990,
		ImageID: imageID,
	}

	login, _ := valueobject.NewLogin("emailvalido@email.com", "AS12a@@56")
	address, _ := valueobject.NewAddress(createChooserInputDTO.City, createChooserInputDTO.State, createChooserInputDTO.Country)
	birthDate, _ := valueobject.NewBirthDate(createChooserInputDTO.Day, createChooserInputDTO.Month, createChooserInputDTO.Year)

	existingUser, _ := entity.NewChooser(
		createChooserInputDTO.Name,
		login,
		address,
		birthDate,
		createChooserInputDTO.ImageID,
	)

	chooserID := existingUser.ID

	mockRepository.EXPECT().
		DoesTheChooserExist(chooserID).
		Return(true, *existingUser, nil).
		Times(1)

	mockRepository.EXPECT().Update(gomock.Any()).Return(nil).Do(func(user *entity.Chooser) {
		assert.Equal(t, createChooserInputDTO.Name, user.Name)
		assert.Equal(t, createChooserInputDTO.City, user.Address.City)
		assert.Equal(t, createChooserInputDTO.State, user.Address.State)
		assert.Equal(t, createChooserInputDTO.Country, user.Address.Country)
		assert.Equal(t, createChooserInputDTO.Day, user.BirthDate.Day)
		assert.Equal(t, createChooserInputDTO.Month, user.BirthDate.Month)
		assert.Equal(t, createChooserInputDTO.Year, user.BirthDate.Year)
		assert.Equal(t, createChooserInputDTO.ImageID, user.ImageID)
	})

	updateChooserInputDTO := UpdateChooserInputDTO{
		ID:      chooserID,
		Name:    "John Doe",
		City:    "Aracaju",
		State:   "Sergipe",
		Country: "Brasil",
		Day:     10,
		Month:   5,
		Year:    1990,
		ImageID: imageID,
	}

	output, problems := updateChooserUseCase.Execute(updateChooserInputDTO)

	expectedOutput := UpdateChooserOutputDTO{
		ID:      chooserID,
		Name:    "John Doe",
		City:    "Aracaju",
		State:   "Sergipe",
		Country: "Brasil",
		Day:     10,
		Month:   5,
		Year:    1990,
		ImageID: imageID,
	}

	fmt.Println(chooserID)
	fmt.Println(expectedOutput)
	fmt.Println(output)

	assert.Equal(t, expectedOutput, output)

	assert.Equal(t, 0, len(problems.ProblemDetails))
}

func TestUpdateChooserUseCase_Execute_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	updateChooserUseCase := NewUpdateChooserUseCase(mockRepository)

	chooserID := uuid.New().String()

	input := UpdateChooserInputDTO{
		ID: chooserID,
	}

	mockRepository.EXPECT().DoesTheChooserExist(input.ID).Return(false, entity.Chooser{}, nil)

	output, problems := updateChooserUseCase.Execute(input)

	assert.Equal(t, UpdateChooserOutputDTO{}, output)

	expectedProblems := []util.ProblemDetails{
		{
			Type:     "Not Found",
			Title:    "Recurso não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Não foi possível encontrar o chooser de id " + input.ID,
			Instance: util.RFC404,
		},
	}
	assert.Equal(t, expectedProblems, problems.ProblemDetails)
}

func TestUpdateChooserUseCase_Execute_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	updateChooserUseCase := NewUpdateChooserUseCase(mockRepository)

	chooserID := uuid.New().String()

	input := UpdateChooserInputDTO{
		ID: chooserID,
	}

	mockRepository.EXPECT().DoesTheChooserExist(input.ID).Return(false, entity.Chooser{}, errors.New("database error"))

	output, problems := updateChooserUseCase.Execute(input)

	assert.Equal(t, UpdateChooserOutputDTO{}, output)

	expectedProblems := []util.ProblemDetails{
		{
			Type:     "Internal Server Error",
			Title:    "Erro ao buscar um chooser",
			Status:   http.StatusInternalServerError,
			Detail:   "database error",
			Instance: util.RFC500,
		},
	}
	assert.Equal(t, expectedProblems, problems.ProblemDetails)
}
