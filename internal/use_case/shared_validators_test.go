package usecase

import (
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

func TestChooserValidator_ChooserFoundAndActive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	useCaseName := "TestChooserValidator"

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	expectedChooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	mockChooserRepo.EXPECT().GetByID(expectedChooser.ID).Return(true, *expectedChooser, nil)

	output, problems := chooserValidator(mockChooserRepo, expectedChooser.ID, useCaseName)

	assert.Empty(t, problems.ProblemDetails)
	assert.Equal(t, expectedChooser, &output)
}

func TestChooserValidator_ChooserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	chooserID := "1"
	useCaseName := "TestChooserValidator"

	mockChooserRepo.EXPECT().GetByID(chooserID).Return(false, entity.Chooser{}, nil)

	result, problemDetails := chooserValidator(mockChooserRepo, chooserID, useCaseName)

	expectedProblemDetails := util.ProblemDetailsOutputDTO{
		ProblemDetails: []util.ProblemDetails{
			{
				Type:     util.TypeNotFound,
				Title:    "Chooser não encontrado",
				Status:   http.StatusNotFound,
				Detail:   "Nenhum chooser com o ID " + chooserID + " foi encontrado",
				Instance: util.RFC404,
			},
		},
	}

	assert.Equal(t, expectedProblemDetails, problemDetails)
	assert.Equal(t, entity.Chooser{}, result)
}

func TestChooserValidator_ChooserInactive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	chooserID := "1"
	useCaseName := "TestChooserValidator"

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	inactiveChooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	inactiveChooser.Deactivate()

	mockChooserRepo.EXPECT().GetByID(chooserID).Return(true, *inactiveChooser, nil)

	result, problemDetails := chooserValidator(mockChooserRepo, chooserID, useCaseName)

	expectedProblemDetails := util.ProblemDetailsOutputDTO{
		ProblemDetails: []util.ProblemDetails{
			{
				Type:     util.TypeNotFound,
				Title:    "Chooser não encontrado",
				Status:   http.StatusNotFound,
				Detail:   "O chooser com o ID " + chooserID + " está desativado",
				Instance: util.RFC404,
			},
		},
	}

	assert.Equal(t, expectedProblemDetails, problemDetails)
	assert.Equal(t, entity.Chooser{}, result)
}

func TestChooserValidator_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	chooserID := "1"
	useCaseName := "TestChooserValidator"

	internalError := fmt.Errorf("Erro interno ao buscar chooser")

	mockChooserRepo.EXPECT().GetByID(chooserID).Return(false, entity.Chooser{}, internalError)

	result, problemDetails := chooserValidator(mockChooserRepo, chooserID, useCaseName)

	expectedProblemDetails := util.ProblemDetailsOutputDTO{
		ProblemDetails: []util.ProblemDetails{
			{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao resgatar chooser de ID " + chooserID,
				Status:   http.StatusInternalServerError,
				Detail:   internalError.Error(),
				Instance: util.RFC503,
			},
		},
	}

	assert.Equal(t, expectedProblemDetails, problemDetails)
	assert.Equal(t, entity.Chooser{}, result)
}
