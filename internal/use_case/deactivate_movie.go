package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type DeactivateMovieInputDTO struct {
	ChooserID string `json:"chooser_id"`
	MovieID   string `json:"movie_id"`
}

type DeactivateMovieOutputDTO struct {
	MovieID   string `json:"movie_id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"success"`
}

type DeactivateMovieUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
	MovieRepository   repositoryinterface.MovieRepositoryInterface
}

func NewDeactivateMovieUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
) *DeactivateMovieUseCase {
	return &DeactivateMovieUseCase{
		ChooserRepository: ChooserRepository,
		MovieRepository:   MovieRepository,
	}
}

func (dm *DeactivateMovieUseCase) Execute(input DeactivateMovieInputDTO) (DeactivateMovieOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, chooser, getChooserError := dm.ChooserRepository.GetByID(input.ChooserID)
	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar chooser de ID " + input.ChooserID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), "DeactivateMovieUseCase", "Use Cases", "Internal Server Error")

		return DeactivateMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheChooserExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum chooser com o ID " + input.ChooserID + " foi encontrado",
			Instance: util.RFC404,
		})

		return DeactivateMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !chooser.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "O chooser com o ID " + input.ChooserID + " está desativado",
			Instance: util.RFC404,
		})

		return DeactivateMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheMovieExist, movie, getMovieError := dm.MovieRepository.GetByID(input.MovieID)
	if getMovieError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar filme de ID " + input.MovieID,
			Status:   http.StatusInternalServerError,
			Detail:   getMovieError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getMovieError.Error(), "DeactivateMovieUseCase", "Use Cases", "Internal Server Error")

		return DeactivateMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return DeactivateMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !movie.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Filme não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "O filme com o ID " + input.MovieID + " está desativado",
			Instance: util.RFC404,
		})

		return DeactivateMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	movie.Deactivate()

	movieDeactivateError := dm.MovieRepository.Deactivate(&movie)
	if movieDeactivateError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao desativar filme de ID " + input.MovieID,
			Status:   http.StatusInternalServerError,
			Detail:   movieDeactivateError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, movieDeactivateError.Error(), "DeactivateMovieUseCase", "Use Cases", "Internal Server Error")

		return DeactivateMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := DeactivateMovieOutputDTO{
		MovieID:   movie.ID,
		Message:   "Filme desativado com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
