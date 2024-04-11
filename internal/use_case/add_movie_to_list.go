package usecase

import (
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type AddMovieToListInputDTO struct {
	ChooserID string `json:"chooser_id"`
	MovieID   string `json:"movie_id"`
	ListID    string `json:"list_id"`
}

type AddMovieToListOutputDTO struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"success"`
}

type AddMovieToListUseCase struct {
	ChooserRepository   repositoryinterface.ChooserRepositoryInterface
	MovieRepository     repositoryinterface.MovieRepositoryInterface
	ListRepository      repositoryinterface.ListRepositoryInterface
	ListMovieRepository repositoryinterface.ListMovieRepositoryInterface
}

func (am *AddMovieToListUseCase) Execute(input AddMovieToListInputDTO) (AddMovieToListOutputDTO, util.ProblemDetailsOutputDTO) {

	_, chooserValidatorProblems := chooserValidator(am.ChooserRepository, input.ChooserID, "AddMovieToListUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return AddMovieToListOutputDTO{}, chooserValidatorProblems
	}

	_, listValidatorProblems := listValidator(am.ListRepository, input.ListID, "AddMovieToListUseCase")
	if len(listValidatorProblems.ProblemDetails) > 0 {
		return AddMovieToListOutputDTO{}, listValidatorProblems
	}

	_, movieValidatorProblems := movieValidator(am.MovieRepository, input.MovieID, "AddMovieToListUseCase")
	if len(movieValidatorProblems.ProblemDetails) > 0 {
		return AddMovieToListOutputDTO{}, movieValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	doesTheListMovieExist, listMovie, getListMovieError := am.ListMovieRepository.GetByListIDAndMovieIDAndChooserID(input.ListID, input.MovieID, input.ChooserID)
	if getListMovieError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao adicionar filme de ID " + input.MovieID + " na lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   getListMovieError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListMovieError.Error(), "AddMovieToListUseCase", "Use Cases", "Internal Server Error")

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if doesTheListMovieExist && listMovie.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Filme já cadastrado",
			Status:   http.StatusConflict,
			Detail:   "O filme com o ID " + input.MovieID + " já está adicionado na lista de ID " + input.ListID,
			Instance: util.RFC409,
		})

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	newListMovie, newListMovieError := entity.NewListMovie(input.ListID, input.MovieID, input.ChooserID)
	if newListMovieError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Erro ao adicionar filme",
			Status:   http.StatusInternalServerError,
			Detail:   "Erro ao adicionar filme de ID " + input.MovieID + " na lista de ID " + input.ListID,
			Instance: util.RFC400,
		})

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	var listMovies []entity.ListMovie

	listMovies = append(listMovies, *newListMovie)

	listMoviesCreationError := am.ListMovieRepository.Create(&listMovies)
	if listMoviesCreationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao adicionar filme de ID " + input.MovieID + " na lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   listMoviesCreationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listMoviesCreationError.Error(), "AddMovieToListUseCase", "Use Cases", "Internal Server Error")

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := AddMovieToListOutputDTO{
		ID:        newListMovie.ID,
		Message:   "Filme adicionado com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
