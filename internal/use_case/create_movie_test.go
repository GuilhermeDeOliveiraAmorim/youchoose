package usecase_test

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
	"youchoose/internal/entity"
	usecase "youchoose/internal/use_case"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovieUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActorRepo := mock.NewMockActorRepositoryInterface(ctrl)
	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockDirectorRepo := mock.NewMockDirectorRepositoryInterface(ctrl)
	mockGenreRepo := mock.NewMockGenreRepositoryInterface(ctrl)
	mockImageRepo := mock.NewMockImageRepositoryInterface(ctrl)
	mockMovieActor := mock.NewMockMovieActorRepositoryInterface(ctrl)
	mockMovieDirector := mock.NewMockMovieDirectorRepositoryInterface(ctrl)
	mockMovieGenre := mock.NewMockMovieGenreRepositoryInterface(ctrl)
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)
	mockMovieWriter := mock.NewMockMovieWriterRepositoryInterface(ctrl)
	mockWriterRepo := mock.NewMockWriterRepositoryInterface(ctrl)

	createMovieUC := usecase.NewCreateMovieUseCase(
		mockChooserRepo,
		mockMovieRepo,
		mockImageRepo,
		mockGenreRepo,
		mockDirectorRepo,
		mockActorRepo,
		mockWriterRepo,
		mockMovieGenre,
		mockMovieDirector,
		mockMovieActor,
		mockMovieWriter,
	)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	file1, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao file1: %v", myError)
	}

	file2, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao criar file2: %v", myError)
	}

	file3, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao criar file2: %v", myError)
	}

	file4, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao criar file2: %v", myError)
	}

	file5, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao criar file2: %v", myError)
	}

	genre := usecase.Genre{
		Name:      "Comedy",
		ImageFile: file1,
		ImageHandler: &multipart.FileHeader{
			Filename: "cover.jpg",
			Size:     150,
		},
	}

	director := usecase.Director{
		Name:        "Guilherme",
		Day:         10,
		Month:       10,
		Year:        1990,
		CountryName: "Brasil",
		Flag:        "🇺🇸",
		ImageFile:   file2,
		ImageHandler: &multipart.FileHeader{
			Filename: "cover.jpg",
			Size:     150,
		},
	}

	actor := usecase.Actor{
		Name:        "Nayara",
		Day:         10,
		Month:       10,
		Year:        1990,
		CountryName: "Brasil",
		Flag:        "🇺🇸",
		ImageFile:   file3,
		ImageHandler: &multipart.FileHeader{
			Filename: "cover.jpg",
			Size:     150,
		},
	}

	writer := usecase.Writer{
		Name:        "Helder",
		Day:         10,
		Month:       10,
		Year:        1990,
		CountryName: "Brasil",
		Flag:        "🇺🇸",
		ImageFile:   file4,
		ImageHandler: &multipart.FileHeader{
			Filename: "cover.jpg",
			Size:     150,
		},
	}

	createMovieInput := usecase.CreateMovieInputDTO{
		ChooserID:   uuid.NewString(),
		Title:       "Test Movie",
		CountryName: "Brazil",
		Flag:        "🇺🇸",
		ReleaseYear: 2023,
		ImageFile:   file5,
		ImageHandler: &multipart.FileHeader{
			Filename: "cover.jpg",
			Size:     150,
		},
		Genres:    []usecase.Genre{genre},
		Directors: []usecase.Director{director},
		Actors:    []usecase.Actor{actor},
		Writers:   []usecase.Writer{writer},
	}

	mockChooserRepo.EXPECT().GetByID(gomock.Any()).Return(true, *chooser, nil)
	mockImageRepo.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()
	mockImageRepo.EXPECT().CreateMany(gomock.Any()).Return(nil)
	mockMovieActor.EXPECT().CreateMany(gomock.Any()).Return(nil)
	mockMovieDirector.EXPECT().CreateMany(gomock.Any()).Return(nil)
	mockMovieWriter.EXPECT().CreateMany(gomock.Any()).Return(nil)
	mockMovieGenre.EXPECT().CreateMany(gomock.Any()).Return(nil)
	mockMovieRepo.EXPECT().Create(gomock.Any()).Return(nil)

	output, problemDetails := createMovieUC.Execute(createMovieInput)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.NotNil(t, output)
	assert.Equal(t, createMovieInput.Title, output.Title)
	assert.Equal(t, createMovieInput.CountryName, output.Nationality.CountryName)
	assert.Equal(t, createMovieInput.Flag, output.Nationality.Flag)
	assert.Equal(t, createMovieInput.ReleaseYear, output.ReleaseYear)
	assert.NotEmpty(t, output.ID)
	assert.NotEmpty(t, output.ImageID)
	assert.NotNil(t, output.Genres)
	assert.NotNil(t, output.Directors)
	assert.NotNil(t, output.Actors)
	assert.NotNil(t, output.Writers)

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
