package usecase

import (
	"testing"
	"youchoose/internal/use_case/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestCreateChooser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	createChooserUseCase := NewCreateChooserUseCase(mockRepository)

	mockRepository.EXPECT().
		ChooserAlreadyExists(gomock.Eq("john.doe@example.com")).
		Return(false, nil).
		Times(1)

	mockRepository.EXPECT().
		Create(gomock.Any()).
		Return(nil).
		Times(1)

	input := CreateChooserInputDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "QWqw@#456",
		City:     "Aracaju",
		State:    "Sergipe",
		Country:  "Brasil",
		Day:      1,
		Month:    1,
		Year:     1990,
		ImageID:  uuid.New().String(),
	}

	output, problemsDetails := createChooserUseCase.Execute(input)

	if len(problemsDetails) > 0 {
		t.Errorf("Unexpected problems during chooser creation: %v", problemsDetails)
	}

	if output.ID == "" {
		t.Error("Expected non-empty ID in output")
	}
}
