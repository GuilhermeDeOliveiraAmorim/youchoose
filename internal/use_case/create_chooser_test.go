package usecase

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
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
	mockImageRepository := mock.NewMockImageRepositoryInterface(ctrl)
	createChooserUseCase := NewCreateChooserUseCase(mockRepository, mockImageRepository)

	name := "John Doe"
	login_1 := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	login_2 := &valueobject.Login{Email: "johnjohn@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	file1, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao file1: %v", myError)
	}

	chooser_1, _ := entity.NewChooser(name, login_1, address, birthDate, imageID)
	chooser_2, _ := entity.NewChooser(name, login_2, address, birthDate, imageID)

	var choosers []entity.Chooser

	choosers = append(choosers, *chooser_1)
	choosers = append(choosers, *chooser_2)

	mockRepository.EXPECT().GetByID(gomock.Any()).Return(true, *chooser_1, nil)
	mockRepository.EXPECT().GetAll().Return(choosers, nil)
	mockImageRepository.EXPECT().Create(gomock.Any()).Return(nil)
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
		ImageFile: file1,
		ImageHandler: &multipart.FileHeader{
			Filename: "profile.jpg",
			Size:     100,
		},
	}

	output, problemDetailsOutputDTO := createChooserUseCase.Execute(input)

	if len(problemDetailsOutputDTO.ProblemDetails) > 0 {
		t.Errorf("Unexpected problems during chooser creation: %v", len(problemDetailsOutputDTO.ProblemDetails))
	}

	if output.ID == "" {
		t.Error("Expected non-empty ID in output")
	}

	dirPath := "/home/guilherme/Workspace/youchoose/internal/upload/"

	fileError := filepath.Walk(dirPath, func(path string, info os.FileInfo, fileError error) error {
		if fileError != nil {
			return fileError
		}

		if !info.IsDir() {
			deleteErr := os.Remove(path)
			if deleteErr != nil {
				fmt.Println("Erro ao excluir arquivo:", deleteErr)
			}
		}

		return nil
	})

	if fileError != nil {
		fmt.Println("Erro ao percorrer o diretório:", fileError)
	}
}
