package usecase_test

import (
	"testing"
	"youchoose/internal/entity"
	usecase "youchoose/internal/use_case"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetMoviesUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)

	getMovies := usecase.NewGetMoviesUseCase(
		mockChooserRepo,
		mockMovieRepo,
	)

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	movie_1, _ := entity.NewMovie("Inception", *nationality, 2010, uuid.NewString())
	movie_2, _ := entity.NewMovie("Inception", *nationality, 2015, uuid.NewString())
	movie_3, _ := entity.NewMovie("Inception", *nationality, 2020, uuid.NewString())

	var movies []entity.Movie

	movies = append(movies, *movie_1)
	movies = append(movies, *movie_2)
	movies = append(movies, *movie_3)

	login, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")

	birthDate, _ := valueobject.NewBirthDate(1, 1, 1990)

	chooser, _ := entity.NewChooser("First Chooser", login, address, birthDate, uuid.New().String())

	mockChooserRepo.EXPECT().GetByID(chooser.ID).Return(true, *chooser, nil)
	mockMovieRepo.EXPECT().GetAll().Return(movies, nil)

	input := usecase.GetMoviesInputDTO{
		ChooserID: chooser.ID,
	}

	var moviesOutputDTO []usecase.MovieOutputDTO

	for _, movie := range movies {
		moviesOutputDTO = append(moviesOutputDTO, usecase.MovieOutputDTO{
			ID:          movie.ID,
			CreatedAt:   movie.CreatedAt,
			Title:       movie.Title,
			Nationality: movie.Nationality,
			ReleaseYear: movie.ReleaseYear,
			ImageID:     movie.ImageID,
			Votes:       movie.Votes,
		})
	}

	outputExpected := usecase.GetMoviesOutputDTO{
		Movies: moviesOutputDTO,
	}

	output, problemDetails := getMovies.Execute(input)

	assert.Equal(t, output, outputExpected)
	assert.Empty(t, problemDetails.ProblemDetails)
}
