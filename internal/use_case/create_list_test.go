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

func TestCreateListUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListRepo := mock.NewMockListRepositoryInterface(ctrl)
	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockImageRepo := mock.NewMockImageRepositoryInterface(ctrl)
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)
	mockListMovieRepo := mock.NewMockListMovieRepositoryInterface(ctrl)

	useCase := usecase.NewCreateListUseCase(mockListRepo, mockChooserRepo, mockImageRepo, mockMovieRepo, mockListMovieRepo)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "Aracaju", State: "Sergipe", Country: "Brazil"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	birthDate, problem := valueobject.NewBirthDate(15, 5, 2010)
	if problem != nil {
		fmt.Println(problem)
	}

	nationality, problem := valueobject.NewNationality("United States", "🇺🇸")
	if problem != nil {
		fmt.Println(problem)
	}

	tomHardy, problem := entity.NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")
	if problem != nil {
		fmt.Println(problem)
	}

	ellenPage, problem := entity.NewActor("Ellen Page", birthDate, nationality, "ellen_page_image")
	if problem != nil {
		fmt.Println(problem)
	}

	nolan, problem := entity.NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")
	if problem != nil {
		fmt.Println(problem)
	}

	nolanWriter, problem := entity.NewWriter("Christopher Nolan", birthDate, nationality, "nolan_image")
	if problem != nil {
		fmt.Println(problem)
	}

	sciFi, problem := entity.NewGenre("Sci-Fi", "image_id_genre")
	if problem != nil {
		fmt.Println(problem)
	}

	movie, problem := entity.NewMovie("Inception", *nationality, []entity.Genre{*sciFi}, []entity.Director{*nolan}, []entity.Actor{*tomHardy, *ellenPage}, []entity.Writer{*nolanWriter}, 2010, "image123")
	if problem != nil {
		t.Errorf("Erro ao criar um novo filme válido: %v", problem)
	}

	var movies []entity.Movie

	movies = append(movies, *movie)
	movies = append(movies, *movie)
	movies = append(movies, *movie)

	mockChooserRepo.EXPECT().GetByID(gomock.Any()).Return(true, *chooser, nil)
	mockListRepo.EXPECT().Create(gomock.Any()).Return(nil)
	mockImageRepo.EXPECT().Create(gomock.Any()).Return(nil).Times(2)
	mockMovieRepo.EXPECT().DoTheseMoviesExist(gomock.Any()).Return(true, nil)
	mockListMovieRepo.EXPECT().Create(gomock.Any()).Return(nil)
	mockListRepo.EXPECT().GetAllMoviesByListID(gomock.Any()).Return(movies, nil)

	file1, myError := os.Open("image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao file1: %v", myError)
	}

	file2, myError := os.Open("image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao criar file2: %v", myError)
	}

	input := usecase.CreateListInputDTO{
		Title:            "Minha Lista",
		ProfileImageFile: file1,
		ProfileImageHandler: &multipart.FileHeader{
			Filename: "profile.jpg",
			Size:     100,
		},
		CoverImageFile: file2,
		CoverImageHandler: &multipart.FileHeader{
			Filename: "cover.jpg",
			Size:     150,
		},
		Description: "Esta é uma lista de filmes",
		Movies:      []string{movie.ID, movie.ID, movie.ID},
		ChooserID:   "chooser123",
	}

	output, err := useCase.Execute(input)

	if len(err.ProblemDetails) > 0 {
		t.Errorf("Erro ao criar uma nova lista: %v", err)
	}

	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, input.Title, output.Title)
	assert.Equal(t, input.Description, output.Description)
	assert.NotEmpty(t, output.ProfileImageID)
	assert.NotEmpty(t, output.CoverImageID)
	assert.Equal(t, input.ChooserID, output.ChooserID)
	assert.Len(t, output.Movies, 3)

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
