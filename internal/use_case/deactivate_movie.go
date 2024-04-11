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
	_, chooserValidatorProblems := chooserValidator(dm.ChooserRepository, input.ChooserID, "DeactivateMovieInputDTO")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return DeactivateMovieOutputDTO{}, chooserValidatorProblems
	}

	movie, movieValidatorProblems := movieValidator(dm.MovieRepository, input.MovieID, "DeactivateMovieInputDTO")
	if len(movieValidatorProblems.ProblemDetails) > 0 {
		return DeactivateMovieOutputDTO{}, movieValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

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
