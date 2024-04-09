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

func TestGetMovieUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)

	getMovie := usecase.NewGetMovieUseCase(
		mockChooserRepo,
		mockMovieRepo,
	)

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	movie, _ := entity.NewMovie("Inception", *nationality, 2010, "image123")

	login1, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address1, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")

	birthDate1, _ := valueobject.NewBirthDate(1, 1, 1990)

	newChooser, _ := entity.NewChooser("Nome 1", login1, address1, birthDate1, uuid.New().String())

	mockChooserRepo.EXPECT().GetByID(newChooser.ID).Return(true, *newChooser, nil)
	mockMovieRepo.EXPECT().GetByID(movie.ID).Return(true, *movie, nil)

	outputExpected := usecase.GetMovieOutputDTO{
		ID:          movie.ID,
		Title:       movie.Title,
		Nationality: movie.Nationality,
		ReleaseYear: movie.ReleaseYear,
		ImageID:     movie.ImageID,
		Votes:       movie.Votes,
		Genres:      movie.Genres,
		Directors:   movie.Directors,
		Actors:      movie.Actors,
		Writers:     movie.Writers,
	}

	input := usecase.GetMovieInputDTO{
		ChooserID: newChooser.ID,
		MovieID:   movie.ID,
	}

	output, problemDetails := getMovie.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, output, outputExpected)
}
