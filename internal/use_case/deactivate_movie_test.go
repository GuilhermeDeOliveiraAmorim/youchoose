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

func TestDeactivateMovieUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)

	deactivateMovieUseCase := usecase.NewDeactivateMovieUseCase(mockChooserRepo, mockMovieRepo)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "Aracaju", State: "Sergipe", Country: "Brazil"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	movie, _ := entity.NewMovie("Inception", *nationality, 2010, "image123")

	login1, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address1, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")

	birthDate1, _ := valueobject.NewBirthDate(1, 1, 1990)

	newChooser, _ := entity.NewChooser("Nome 1", login1, address1, birthDate1, uuid.New().String())

	mockChooserRepo.EXPECT().GetByID(newChooser.ID).Return(true, *chooser, nil)

	mockMovieRepo.EXPECT().GetByID(movie.ID).Return(true, *movie, nil)

	mockMovieRepo.EXPECT().Deactivate(gomock.Any()).Do(func(m *entity.Movie) {
		assert.False(t, m.Active)
	}).Return(nil)

	input := usecase.DeactivateMovieInputDTO{
		ChooserID: newChooser.ID,
		MovieID:   movie.ID,
	}

	output, problemDetails := deactivateMovieUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)

	expectedOutput := usecase.DeactivateMovieOutputDTO{
		MovieID:   movie.ID,
		Message:   "Filme desativado com sucesso",
		IsSuccess: true,
	}

	assert.Equal(t, expectedOutput, output)
}
