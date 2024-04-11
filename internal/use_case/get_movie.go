package usecase

import (
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetMovieInputDTO struct {
	ChooserID string `json:"chooser_id"`
	MovieID   string `json:"movie_id"`
}

type GetMovieUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
	MovieRepository   repositoryinterface.MovieRepositoryInterface
}

func NewGetMovieUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
) *GetMovieUseCase {
	return &GetMovieUseCase{
		ChooserRepository: ChooserRepository,
		MovieRepository:   MovieRepository,
	}
}

func (gm *GetMovieUseCase) Execute(input GetMovieInputDTO) (MovieOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(gm.ChooserRepository, input.ChooserID, "GetMovieUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return MovieOutputDTO{}, chooserValidatorProblems
	}

	movie, movieValidatorProblems := movieValidator(gm.MovieRepository, input.MovieID, "GetMovieUseCase")
	if len(movieValidatorProblems.ProblemDetails) > 0 {
		return MovieOutputDTO{}, movieValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	output := NewMovieOutputDTO(movie)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
