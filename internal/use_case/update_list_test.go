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
	"github.com/oklog/ulid/v2"
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
	address := &valueobject.Address{City: "Aracaju", State: "Sergipe", Country: "Brasil"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	movie_1, _ := entity.NewMovie("Inception_1", *nationality, 2011, uuid.New().String())
	movie_2, _ := entity.NewMovie("Inception_2", *nationality, 2012, uuid.New().String())
	movie_3, _ := entity.NewMovie("Inception_3", *nationality, 2013, uuid.New().String())

	var movies []entity.Movie

	movies = append(movies, *movie_1)
	movies = append(movies, *movie_2)
	movies = append(movies, *movie_3)
	imageType := "jpeg"
	size := int64(50000)

	image_1, _ := entity.NewImage(ulid.Make().String(), imageType, size)
	image_2, _ := entity.NewImage(ulid.Make().String(), imageType, size)

	file1, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao file1: %v", myError)
	}

	file2, myError := os.Open("/home/guilherme/Workspace/youchoose/image.jpeg")
	if myError != nil {
		t.Errorf("Erro ao criar file2: %v", myError)
	}

	list, _ := entity.NewList("Minha Lista", "Descrição da Lista", image_2.ID, image_1.ID, chooser.ID)

	list.AddMovies(movies)

	input := UpdateListInputDTO{
		ListID:           list.ID,
		ChooserID:        chooser.ID,
		Title:            "Nova Lista",
		Description:      "Descrição da nova Lista",
		ProfileImageFile: file1,
		ProfileImageHandler: &multipart.FileHeader{
			Filename: ulid.Make().String(),
			Size:     100,
		},
		CoverImageFile: file2,
		CoverImageHandler: &multipart.FileHeader{
			Filename: ulid.Make().String(),
			Size:     150,
		},
		Movies: []string{movie_1.ID, movie_2.ID, movie_3.ID},
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *chooser, nil)
	mockListRepo.EXPECT().GetByID(input.ListID).Return(true, *list, nil)
	mockMovieRepo.EXPECT().DoTheseMoviesExist(gomock.Any()).Return(true, movies, nil)
	mockImageRepo.EXPECT().CreateMany(gomock.Any()).Return(nil)
	mockListRepo.EXPECT().Update(gomock.Any()).Return(nil)

	output, problem := useCase.Execute(input)

	assert.Len(t, problem.ProblemDetails, 0)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ChooserID)
	assert.NotEqual(t, list.Title, output.Title)
	assert.NotEqual(t, list.Description, output.Description)
	assert.NotEmpty(t, output.ProfileImageID)
	assert.NotEmpty(t, output.CoverImageID)
	assert.Equal(t, input.ChooserID, output.ChooserID)
	assert.Len(t, output.Movies, 3)
}
