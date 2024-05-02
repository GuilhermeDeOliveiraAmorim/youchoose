package usecase

import (
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type RegisterMovieVoteInputDTO struct {
	ChooserID     string `json:"chooser_id"`
	ListID        string `json:"list_id"`
	FirstMovieID  string `json:"first_movie_id"`
	SecondMovieID string `json:"second_movie_id"`
	ChosenMovieID string `json:"chosen_movie_id"`
}

type RegisterMovieVoteOutputDTO struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"success"`
}

type RegisterMovieVoteUseCase struct {
	ChooserRepository  repositoryinterface.ChooserRepositoryInterface
	ListRepository     repositoryinterface.ListRepositoryInterface
	VotationRepository repositoryinterface.VotationRepositoryInterface
	MovieRepository    repositoryinterface.MovieRepositoryInterface
}

func NewRegisterMovieVoteUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ListRepository repositoryinterface.ListRepositoryInterface,
	VotationRepository repositoryinterface.VotationRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
) *RegisterMovieVoteUseCase {
	return &RegisterMovieVoteUseCase{
		ChooserRepository:  ChooserRepository,
		ListRepository:     ListRepository,
		VotationRepository: VotationRepository,
		MovieRepository:    MovieRepository,
	}
}

func (rv *RegisterMovieVoteUseCase) Execute(input RegisterMovieVoteInputDTO) (RegisterMovieVoteOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(rv.ChooserRepository, input.ChooserID, "RegisterMovieVoteUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return RegisterMovieVoteOutputDTO{}, chooserValidatorProblems
	}

	_, listValidatorProblems := listValidator(rv.ListRepository, input.ListID, "RegisterMovieVoteUseCase")
	if len(listValidatorProblems.ProblemDetails) > 0 {
		return RegisterMovieVoteOutputDTO{}, listValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	var moviesIDs []string

	moviesIDs = append(moviesIDs, input.FirstMovieID)
	moviesIDs = append(moviesIDs, input.SecondMovieID)
	moviesIDs = append(moviesIDs, input.ChosenMovieID)

	doesTheMoviesExist, movies, doesTheMoviesExistError := rv.MovieRepository.DoTheseMoviesExist(moviesIDs)
	if doesTheMoviesExistError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao resgatar lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   doesTheMoviesExistError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, doesTheMoviesExistError.Error(), "RegisterMovieVoteUseCase", "Use Cases", util.TypeInternalServerError)

		return RegisterMovieVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheMoviesExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Um ou mais filmes não encontrados",
			Status:   http.StatusConflict,
			Detail:   "Um ou mais ids dos filmes não retornou resultado",
			Instance: util.RFC409,
		})

		return RegisterMovieVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheVotationExist, doesTheVotationExistError := rv.VotationRepository.VotationAlreadyExists(input.ChooserID, input.ListID, input.FirstMovieID, input.SecondMovieID, input.ChosenMovieID)
	if doesTheVotationExistError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao resgatar lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   doesTheVotationExistError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, doesTheVotationExistError.Error(), "RegisterMovieVoteUseCase", "Use Cases", util.TypeInternalServerError)

		return RegisterMovieVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if doesTheVotationExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeConflict,
			Title:    "Votação já existe",
			Status:   http.StatusConflict,
			Detail:   "A votação já foi registrada anteriormente",
			Instance: util.RFC409,
		})

		return RegisterMovieVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	newVotation, newVotationError := entity.NewVotation(input.ChooserID, input.ListID, input.FirstMovieID, input.SecondMovieID, input.ChosenMovieID)
	if newVotationError != nil {
		problemsDetails = append(problemsDetails, newVotationError...)

		return RegisterMovieVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	movies[2].IncrementVotes()

	movieUpdatedError := rv.MovieRepository.Update(&movies[2])
	if movieUpdatedError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao atualizar filme de ID " + movies[2].ID,
			Status:   http.StatusInternalServerError,
			Detail:   movieUpdatedError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, movieUpdatedError.Error(), "RegisterMovieVoteUseCase", "Use Cases", util.TypeInternalServerError)

		return RegisterMovieVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	registerVoteCreationError := rv.VotationRepository.Create(newVotation)
	if registerVoteCreationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao persistir uma votação",
			Status:   http.StatusInternalServerError,
			Detail:   registerVoteCreationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, registerVoteCreationError.Error(), "RegisterMovieVoteUseCase", "Use Cases", util.TypeInternalServerError)

		return RegisterMovieVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := RegisterMovieVoteOutputDTO{
		ID:        newVotation.ID,
		Message:   "Voto registrado com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
