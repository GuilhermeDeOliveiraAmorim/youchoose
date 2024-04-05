package usecase

import (
	"mime/multipart"
	"os"
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateListUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListRepo := mock.NewMockListRepositoryInterface(ctrl)
	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockImageRepo := mock.NewMockImageRepositoryInterface(ctrl)
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)
	mockListMovieRepo := mock.NewMockListMovieRepositoryInterface(ctrl)

	useCase := NewUpdateListUseCase(mockListRepo, mockChooserRepo, mockImageRepo, mockMovieRepo, mockListMovieRepo)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "Aracaju", State: "Sergipe", Country: "Brazil"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	movie, _ := entity.NewMovie("Inception", *nationality, 2010, "image123")

	var movies []entity.Movie

	movies = append(movies, *movie)
	movies = append(movies, *movie)
	movies = append(movies, *movie)

	name = uuid.NewString()
	imageType := "jpeg"
	size := int64(50000)

	image_1, _ := entity.NewImage(name, imageType, size)
	image_2, _ := entity.NewImage(name, imageType, size)

	file1, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao file1: %v", myError)
	}

	file2, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao criar file2: %v", myError)
	}

	list, _ := entity.NewList("Minha Lista", "Descrição da Lista", "Minha Lista", "Descrição da Lista", chooser.ID)

	list.AddMovies(movies)
	list.CoverImageID = image_1.ID
	list.ProfileImageID = image_2.ID

	input := UpdateListInputDTO{
		ID:               list.ID,
		ChooserID:        chooser.ID,
		Title:            "Nova Lista",
		Description:      "Descrição da nova Lista",
		ProfileImageFile: file1,
		ProfileImageHandler: &multipart.FileHeader{
			Filename: uuid.NewString(),
			Size:     100,
		},
		CoverImageFile: file2,
		CoverImageHandler: &multipart.FileHeader{
			Filename: uuid.NewString(),
			Size:     150,
		},
		Movies: []string{movie.ID, movie.ID, movie.ID},
	}

	mockListRepo.EXPECT().GetByID(input.ID).Return(true, *list, nil)
	mockListRepo.EXPECT().Update(gomock.Any()).Return(nil)
	mockChooserRepo.EXPECT().GetByID(chooser.ID).Return(true, *chooser, nil)
	mockMovieRepo.EXPECT().DoTheseMoviesExist(gomock.Any()).Return(true, movies, nil)
	mockImageRepo.EXPECT().Create(gomock.Any()).Return(nil).Times(2)
	mockListMovieRepo.EXPECT().DeactivateAll(gomock.Any()).Return(nil)
	mockListMovieRepo.EXPECT().Create(gomock.Any()).Return(nil)

	output, problem := useCase.Execute(input)

	assert.Len(t, problem.ProblemDetails, 0)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.NotEqual(t, list.Title, output.Title)
	assert.NotEqual(t, list.Description, output.Description)
	assert.NotEmpty(t, output.ProfileImageID)
	assert.NotEmpty(t, output.CoverImageID)
	assert.Equal(t, input.ChooserID, output.ChooserID)
	assert.Len(t, output.Movies, 3)
}
