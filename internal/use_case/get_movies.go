package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetMoviesInputDTO struct {
	ChooserID string `json:"chooser_id"`
}

type GetMoviesOutputDTO struct {
	Movies []MovieOutputDTO `json:"movies"`
}

type GetMoviesUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
	MovieRepository   repositoryinterface.MovieRepositoryInterface
}

func NewGetMoviesUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
) *GetMoviesUseCase {
	return &GetMoviesUseCase{
		ChooserRepository: ChooserRepository,
		MovieRepository:   MovieRepository,
	}
}

func (gm *GetMoviesUseCase) Execute(input GetMoviesInputDTO) (GetMoviesOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(gm.ChooserRepository, input.ChooserID, "GetMoviesUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return GetMoviesOutputDTO{}, chooserValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	movies, getMoviesError := gm.MovieRepository.GetAll()
	if getMoviesError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao resgatar todos os filmes",
			Status:   http.StatusInternalServerError,
			Detail:   getMoviesError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getMoviesError.Error(), "GetMoviesUseCase", "Use Cases", util.TypeInternalServerError)

		return GetMoviesOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if len(movies) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Filmes não encontrados",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum filme foi encontrado",
			Instance: util.RFC404,
		})

		return GetMoviesOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	allGetMoviesOutputDTO := []MovieOutputDTO{}

	for _, movie := range movies {
		outputMovie := NewMovieOutputDTO(movie)

		allGetMoviesOutputDTO = append(allGetMoviesOutputDTO, outputMovie)
	}

	output := GetMoviesOutputDTO{
		Movies: allGetMoviesOutputDTO,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
