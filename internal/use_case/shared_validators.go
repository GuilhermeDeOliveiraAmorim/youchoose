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
			Title:    util.SharedErrorTitleErrorGetResource,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), useCaseName, "Use Cases", util.TypeInternalServerError)
	} else if !doesTheChooserExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    util.SharedErrorTitleNotFound,
			Status:   http.StatusNotFound,
			Detail:   util.ChooserErrorDetailNotFound,
			Instance: util.RFC404,
		})
	} else if !chooser.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    util.SharedErrorTitleNotFound,
			Status:   http.StatusNotFound,
			Detail:   util.ChooserErrorDetailDeactivate,
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
			Title:    util.SharedErrorTitleErrorGetResource,
			Status:   http.StatusInternalServerError,
			Detail:   getListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListError.Error(), useCaseName, "Use Cases", util.TypeInternalServerError)
	} else if !doesTheListExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    util.SharedErrorTitleNotFound,
			Status:   http.StatusNotFound,
			Detail:   util.ListErrorDetailNotFound,
			Instance: util.RFC404,
		})
	} else if !list.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    util.SharedErrorTitleNotFound,
			Status:   http.StatusNotFound,
			Detail:   util.ListErrorDetailDeactivate,
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
			Title:    util.SharedErrorTitleErrorGetResource,
			Status:   http.StatusInternalServerError,
			Detail:   getMovieError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getMovieError.Error(), useCaseName, "Use Cases", util.TypeInternalServerError)
	} else if !doesTheMovieExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    util.SharedErrorTitleNotFound,
			Status:   http.StatusNotFound,
			Detail:   util.MovieErrorDetailNotFound,
			Instance: util.RFC404,
		})
	} else if !movie.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    util.SharedErrorTitleNotFound,
			Status:   http.StatusNotFound,
			Detail:   util.MovieErrorDetailDeactivate,
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
