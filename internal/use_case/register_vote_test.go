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
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)

	registerVoteUC := NewRegisterVoteUseCase(mockChooserRepo, mockListRepo, mockVotationRepo, mockMovieRepo)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	list, _ := entity.NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	tomHardy, _ := entity.NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")

	ellenPage, _ := entity.NewActor("Ellen Page", birthDate, nationality, "ellen_page_image")

	nolan, _ := entity.NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")

	nolanWriter, _ := entity.NewWriter("Christopher Nolan", birthDate, nationality, "nolan_image")

	sciFi, _ := entity.NewGenre("Sci-Fi", "image_id_genre")

	movie_1, _ := entity.NewMovie("Inception", *nationality, []entity.Genre{*sciFi}, []entity.Director{*nolan}, []entity.Actor{*tomHardy, *ellenPage}, []entity.Writer{*nolanWriter}, 2010, "image123")

	movie_2, _ := entity.NewMovie("Inception", *nationality, []entity.Genre{*sciFi}, []entity.Director{*nolan}, []entity.Actor{*tomHardy, *ellenPage}, []entity.Writer{*nolanWriter}, 2010, "image123")

	movie_3, _ := entity.NewMovie("Inception", *nationality, []entity.Genre{*sciFi}, []entity.Director{*nolan}, []entity.Actor{*tomHardy, *ellenPage}, []entity.Writer{*nolanWriter}, 2010, "image123")

	var movieIDs []string
	var movies []entity.Movie

	movieIDs = append(movieIDs, movie_1.ID)
	movieIDs = append(movieIDs, movie_2.ID)
	movieIDs = append(movieIDs, movie_3.ID)

	movies = append(movies, *movie_1)
	movies = append(movies, *movie_2)
	movies = append(movies, *movie_3)

	input := RegisterVoteInputDTO{
		ChooserID:     chooser.ID,
		ListID:        list.ID,
		FirstMovieID:  movie_1.ID,
		SecondMovieID: movie_2.ID,
		ChosenMovieID: movie_3.ID,
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *chooser, nil)
	mockListRepo.EXPECT().GetByID(input.ListID).Return(true, *list, nil)
	mockMovieRepo.EXPECT().DoTheseMoviesExist(movieIDs).Return(true, movies, nil)
	mockVotationRepo.EXPECT().VotationAlreadyExists(input.ChooserID, input.ListID, input.FirstMovieID, input.SecondMovieID, input.ChosenMovieID).Return(false, nil)
	mockMovieRepo.EXPECT().Update(gomock.Any()).Return(nil)
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
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)

	registerVoteUC := NewRegisterVoteUseCase(mockChooserRepo, mockListRepo, mockVotationRepo, mockMovieRepo)

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
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)

	registerVoteUC := NewRegisterVoteUseCase(mockChooserRepo, mockListRepo, mockVotationRepo, mockMovieRepo)

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
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)

	registerVoteUC := NewRegisterVoteUseCase(mockChooserRepo, mockListRepo, mockVotationRepo, mockMovieRepo)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	list, _ := entity.NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	nationality, _ := valueobject.NewNationality("United States", "🇺🇸")

	tomHardy, _ := entity.NewActor("Tom Hardy", birthDate, nationality, "tom_hardy_image")

	ellenPage, _ := entity.NewActor("Ellen Page", birthDate, nationality, "ellen_page_image")

	nolan, _ := entity.NewDirector("Christopher Nolan", birthDate, nationality, "nolan_image")

	nolanWriter, _ := entity.NewWriter("Christopher Nolan", birthDate, nationality, "nolan_image")

	sciFi, _ := entity.NewGenre("Sci-Fi", "image_id_genre")

	movie_1, _ := entity.NewMovie("Inception", *nationality, []entity.Genre{*sciFi}, []entity.Director{*nolan}, []entity.Actor{*tomHardy, *ellenPage}, []entity.Writer{*nolanWriter}, 2010, "image123")

	movie_2, _ := entity.NewMovie("Inception", *nationality, []entity.Genre{*sciFi}, []entity.Director{*nolan}, []entity.Actor{*tomHardy, *ellenPage}, []entity.Writer{*nolanWriter}, 2010, "image123")

	movie_3, _ := entity.NewMovie("Inception", *nationality, []entity.Genre{*sciFi}, []entity.Director{*nolan}, []entity.Actor{*tomHardy, *ellenPage}, []entity.Writer{*nolanWriter}, 2010, "image123")

	var movieIDs []string
	var movies []entity.Movie

	movieIDs = append(movieIDs, movie_1.ID)
	movieIDs = append(movieIDs, movie_2.ID)
	movieIDs = append(movieIDs, movie_3.ID)

	movies = append(movies, *movie_1)
	movies = append(movies, *movie_2)
	movies = append(movies, *movie_3)

	input := RegisterVoteInputDTO{
		ChooserID:     chooser.ID,
		ListID:        list.ID,
		FirstMovieID:  movie_1.ID,
		SecondMovieID: movie_2.ID,
		ChosenMovieID: movie_3.ID,
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *chooser, nil)
	mockListRepo.EXPECT().GetByID(input.ListID).Return(true, *list, nil)
	mockMovieRepo.EXPECT().DoTheseMoviesExist(movieIDs).Return(true, movies, nil)
	mockVotationRepo.EXPECT().VotationAlreadyExists(input.ChooserID, input.ListID, input.FirstMovieID, input.SecondMovieID, input.ChosenMovieID).Return(true, nil)

	output, problemDetails := registerVoteUC.Execute(input)

	assert.Empty(t, output)
	assert.NotEmpty(t, problemDetails.ProblemDetails)
	assert.Equal(t, http.StatusConflict, problemDetails.ProblemDetails[0].Status)
}
