package usecase

import (
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestCreateChooser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	createChooserUseCase := NewCreateChooserUseCase(mockRepository)

	name := "John Doe"
	login_1 := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	login_2 := &valueobject.Login{Email: "johnjohn@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser_1, _ := entity.NewChooser(name, login_1, address, birthDate, imageID)
	chooser_2, _ := entity.NewChooser(name, login_2, address, birthDate, imageID)

	var choosers []entity.Chooser

	choosers = append(choosers, *chooser_1)
	choosers = append(choosers, *chooser_2)

	mockRepository.EXPECT().GetByID(gomock.Any()).Return(true, *chooser_1, nil)
	mockRepository.EXPECT().GetAll().Return(choosers, nil)

	mockRepository.EXPECT().
		Create(gomock.Any()).
		Return(nil).
		Times(1)

	input := CreateChooserInputDTO{
		ChooserID: chooser_1.ID,
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Password:  "QWqw@#456",
		City:      "Aracaju",
		State:     "Sergipe",
		Country:   "Brasil",
		Day:       1,
		Month:     1,
		Year:      1990,
		ImageID:   uuid.New().String(),
	}

	output, problemDetailsOutputDTO := createChooserUseCase.Execute(input)

	if len(problemDetailsOutputDTO.ProblemDetails) > 0 {
		t.Errorf("Unexpected problems during chooser creation: %v", len(problemDetailsOutputDTO.ProblemDetails))
	}

	if output.ID == "" {
		t.Error("Expected non-empty ID in output")
	}
}
