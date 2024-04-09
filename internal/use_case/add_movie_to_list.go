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
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, chooser, getChooserError := am.ChooserRepository.GetByID(input.ChooserID)
	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar chooser de ID " + input.ChooserID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), "AddMovieToListUseCase", "Use Cases", "Internal Server Error")

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheChooserExist || !chooser.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum chooser com o ID " + input.ChooserID + " foi encontrado",
			Instance: util.RFC404,
		})

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !chooser.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum chooser com o ID " + input.ChooserID + " foi encontrado",
			Instance: util.RFC404,
		})

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheListExist, list, getListError := am.ListRepository.GetByID(input.ListID)
	if getListError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   getListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListError.Error(), "AddMovieToListUseCase", "Use Cases", "Internal Server Error")

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheListExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Lista não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "Nenhuma lista com o ID " + input.ListID + " foi encontrada",
			Instance: util.RFC404,
		})

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !list.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Lista não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "Nenhuma lista com o ID " + input.ListID + " foi encontrada",
			Instance: util.RFC404,
		})

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheMovieExist, movie, getMovieError := am.MovieRepository.GetByID(input.MovieID)
	if getMovieError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar filme de ID " + input.MovieID,
			Status:   http.StatusInternalServerError,
			Detail:   getMovieError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getMovieError.Error(), "AddMovieToListUseCase", "Use Cases", "Internal Server Error")

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheMovieExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Filme não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum filme com o ID " + input.MovieID + " foi encontrado",
			Instance: util.RFC404,
		})

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !movie.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Filme não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum filme com o ID " + input.MovieID + " foi encontrado",
			Instance: util.RFC404,
		})

		return AddMovieToListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

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
