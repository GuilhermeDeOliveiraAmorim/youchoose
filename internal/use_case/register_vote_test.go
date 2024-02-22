package usecase

import (
	"net/http"
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRegisterVoteUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockListRepo := mock.NewMockListRepositoryInterface(ctrl)
	mockVotationRepo := mock.NewMockVotationRepositoryInterface(ctrl)

	registerVoteUC := NewRegisterVoteUseCase(mockChooserRepo, mockListRepo, mockVotationRepo)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	list, _ := entity.NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	input := RegisterVoteInputDTO{
		ChooserID:     chooser.ID,
		ListID:        list.ID,
		FirstMovieID:  uuid.New().String(),
		SecondMovieID: uuid.New().String(),
		ChosenMovieID: uuid.New().String(),
	}
	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *chooser, nil)
	mockListRepo.EXPECT().GetByID(input.ListID).Return(true, *list, nil)
	mockVotationRepo.EXPECT().VotationAlreadyExists(input.ChooserID, input.ListID, input.FirstMovieID, input.SecondMovieID, input.ChosenMovieID).Return(false, nil)
	mockVotationRepo.EXPECT().Create(gomock.Any()).Return(nil)

	output, problemDetails := registerVoteUC.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.NotEmpty(t, output)
}

func TestRegisterVoteUseCase_ChooserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockListRepo := mock.NewMockListRepositoryInterface(ctrl)
	mockVotationRepo := mock.NewMockVotationRepositoryInterface(ctrl)

	registerVoteUC := NewRegisterVoteUseCase(mockChooserRepo, mockListRepo, mockVotationRepo)

	input := RegisterVoteInputDTO{
		ChooserID:     "non_existing_chooser_id",
		ListID:        "list_id",
		FirstMovieID:  "first_movie_id",
		SecondMovieID: "second_movie_id",
		ChosenMovieID: "chosen_movie_id",
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(false, entity.Chooser{}, nil)

	output, problemDetails := registerVoteUC.Execute(input)

	assert.Empty(t, output)
	assert.NotEmpty(t, problemDetails.ProblemDetails)
	assert.Equal(t, http.StatusNotFound, problemDetails.ProblemDetails[0].Status)
}

func TestRegisterVoteUseCase_ListNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockListRepo := mock.NewMockListRepositoryInterface(ctrl)
	mockVotationRepo := mock.NewMockVotationRepositoryInterface(ctrl)

	registerVoteUC := NewRegisterVoteUseCase(mockChooserRepo, mockListRepo, mockVotationRepo)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	list, _ := entity.NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	actor, _ := entity.NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")

	genre, _ := entity.NewGenre("Ação", "image_id_genre")

	director, _ := entity.NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")

	movie1, _ := entity.NewMovie("Inception", *nationality, []entity.Genre{*genre}, []entity.Director{*director}, []entity.Actor{*actor}, []entity.Writer{}, 2010, "image123")

	list.AddMovies([]entity.Movie{*movie1})

	input := RegisterVoteInputDTO{
		ChooserID:     chooser.ID,
		ListID:        list.ID,
		FirstMovieID:  uuid.New().String(),
		SecondMovieID: uuid.New().String(),
		ChosenMovieID: uuid.New().String(),
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *chooser, nil)
	mockListRepo.EXPECT().GetByID(input.ListID).Return(false, *list, nil)

	output, problemDetails := registerVoteUC.Execute(input)

	assert.Empty(t, output)
	assert.NotEmpty(t, problemDetails.ProblemDetails)
	assert.Equal(t, http.StatusNotFound, problemDetails.ProblemDetails[0].Status)
}

func TestRegisterVoteUseCase_VotationAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockListRepo := mock.NewMockListRepositoryInterface(ctrl)
	mockVotationRepo := mock.NewMockVotationRepositoryInterface(ctrl)

	registerVoteUC := NewRegisterVoteUseCase(mockChooserRepo, mockListRepo, mockVotationRepo)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	list, _ := entity.NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	actor, _ := entity.NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")

	genre, _ := entity.NewGenre("Ação", "image_id_genre")

	director, _ := entity.NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")

	movie1, _ := entity.NewMovie("Inception", *nationality, []entity.Genre{*genre}, []entity.Director{*director}, []entity.Actor{*actor}, []entity.Writer{}, 2010, "image123")

	list.AddMovies([]entity.Movie{*movie1})

	input := RegisterVoteInputDTO{
		ChooserID:     chooser.ID,
		ListID:        list.ID,
		FirstMovieID:  "first_movie_id",
		SecondMovieID: "second_movie_id",
		ChosenMovieID: "chosen_movie_id",
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *chooser, nil)
	mockListRepo.EXPECT().GetByID(input.ListID).Return(true, *list, nil)
	mockVotationRepo.EXPECT().VotationAlreadyExists(input.ChooserID, input.ListID, input.FirstMovieID, input.SecondMovieID, input.ChosenMovieID).Return(true, nil)

	output, problemDetails := registerVoteUC.Execute(input)

	assert.Empty(t, output)
	assert.NotEmpty(t, problemDetails.ProblemDetails)
	assert.Equal(t, http.StatusConflict, problemDetails.ProblemDetails[0].Status)
}
