package usecase

import (
	"fmt"
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestAddMovieToList_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockListRepo := mock.NewMockListRepositoryInterface(ctrl)
	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockMovieRepo := mock.NewMockMovieRepositoryInterface(ctrl)
	mockListMovieRepo := mock.NewMockListMovieRepositoryInterface(ctrl)

	usecase := &AddMovieToListUseCase{
		ChooserRepository:   mockChooserRepo,
		MovieRepository:     mockMovieRepo,
		ListRepository:      mockListRepo,
		ListMovieRepository: mockListMovieRepo,
	}

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

	list, _ := entity.NewList("profile123", "cover123", "Minha Lista", "Descrição da Lista", "chooser123")

	list.AddMovies([]entity.Movie{*movie})

	mockChooserRepo.EXPECT().GetByID(gomock.Any()).Return(true, *chooser, nil)
	mockListRepo.EXPECT().GetByID(gomock.Any()).Return(true, *list, nil)
	mockMovieRepo.EXPECT().GetByID(gomock.Any()).Return(true, *movie, nil)
	mockListMovieRepo.EXPECT().GetByListIDAndMovieIDAndChooserID(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, entity.ListMovie{}, nil)
	mockListMovieRepo.EXPECT().Create(gomock.Any()).Return(nil)

	input := AddMovieToListInputDTO{
		ChooserID: uuid.NewString(),
		MovieID:   uuid.NewString(),
		ListID:    uuid.NewString(),
	}

	output, problemsDetails := usecase.Execute(input)
	
	if len(problemsDetails.ProblemDetails) > 0 {
		t.Errorf("ProblemDetails: %+v", problemsDetails.ProblemDetails)
	}

	if output.ID == "" {
		t.Error("Expected non-empty ID in output")
	}
}
