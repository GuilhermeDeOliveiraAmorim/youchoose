package usecase

import (
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

func chooserValidator(chooserRepository repositoryinterface.ChooserRepositoryInterface, chooserID, useCaseName string) (entity.Chooser, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, chooser, getChooserError := chooserRepository.GetByID(chooserID)

	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao resgatar chooser de ID " + chooserID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), useCaseName, "Use Cases", util.TypeInternalServerError)
	} else if !doesTheChooserExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum chooser com o ID " + chooserID + " foi encontrado",
			Instance: util.RFC404,
		})
	} else if !chooser.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "O chooser com o ID " + chooserID + " está desativado",
			Instance: util.RFC404,
		})
	}

	if len(problemsDetails) > 0 {
		return entity.Chooser{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else {
		return chooser, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}
}

func listValidator(listRepository repositoryinterface.ListRepositoryInterface, listID, useCaseName string) (entity.List, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheListExist, list, getListError := listRepository.GetByID(listID)
	if getListError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao resgatar lista de ID " + listID,
			Status:   http.StatusInternalServerError,
			Detail:   getListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListError.Error(), useCaseName, "Use Cases", util.TypeInternalServerError)
	} else if !doesTheListExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Lista não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "Nenhuma lista com o ID " + listID + " foi encontrada",
			Instance: util.RFC404,
		})
	} else if !list.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Lista não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "A lista com o ID " + listID + " está desativada",
			Instance: util.RFC404,
		})
	}

	if len(problemsDetails) > 0 {
		return entity.List{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else {
		return list, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}
}

func movieValidator(movieRepository repositoryinterface.MovieRepositoryInterface, movieID, useCaseName string) (entity.Movie, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheMovieExist, movie, getMovieError := movieRepository.GetByID(movieID)
	if getMovieError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao resgatar filme de ID " + movieID,
			Status:   http.StatusInternalServerError,
			Detail:   getMovieError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getMovieError.Error(), useCaseName, "Use Cases", util.TypeInternalServerError)
	} else if !doesTheMovieExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Filme não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum filme com o ID " + movieID + " foi encontrado",
			Instance: util.RFC404,
		})
	} else if !movie.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Filme não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "O filme com o ID " + movieID + " está desativado",
			Instance: util.RFC404,
		})
	}

	if len(problemsDetails) > 0 {
		return entity.Movie{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else {
		return movie, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}
}
