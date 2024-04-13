package usecase

import (
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type RegisterVoteInputDTO struct {
	ChooserID     string `json:"chooser_id"`
	ListID        string `json:"list_id"`
	FirstMovieID  string `json:"first_movie_id"`
	SecondMovieID string `json:"second_movie_id"`
	ChosenMovieID string `json:"chosen_movie_id"`
}

type RegisterVoteOutputDTO struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"success"`
}

type RegisterVoteUseCase struct {
	ChooserRepository  repositoryinterface.ChooserRepositoryInterface
	ListRepository     repositoryinterface.ListRepositoryInterface
	VotationRepository repositoryinterface.VotationRepositoryInterface
	MovieRepository    repositoryinterface.MovieRepositoryInterface
}

func NewRegisterVoteUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ListRepository repositoryinterface.ListRepositoryInterface,
	VotationRepository repositoryinterface.VotationRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
) *RegisterVoteUseCase {
	return &RegisterVoteUseCase{
		ChooserRepository:  ChooserRepository,
		ListRepository:     ListRepository,
		VotationRepository: VotationRepository,
		MovieRepository:    MovieRepository,
	}
}

func (rv *RegisterVoteUseCase) Execute(input RegisterVoteInputDTO) (RegisterVoteOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(rv.ChooserRepository, input.ChooserID, "RegisterVoteUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return RegisterVoteOutputDTO{}, chooserValidatorProblems
	}

	_, listValidatorProblems := listValidator(rv.ListRepository, input.ListID, "RegisterVoteUseCase")
	if len(listValidatorProblems.ProblemDetails) > 0 {
		return RegisterVoteOutputDTO{}, listValidatorProblems
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

		util.NewLoggerError(http.StatusInternalServerError, doesTheMoviesExistError.Error(), "RegisterVoteUseCase", "Use Cases", util.TypeInternalServerError)

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		util.NewLoggerError(http.StatusInternalServerError, doesTheVotationExistError.Error(), "RegisterVoteUseCase", "Use Cases", util.TypeInternalServerError)

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	newVotation, newVotationError := entity.NewVotation(input.ChooserID, input.ListID, input.FirstMovieID, input.SecondMovieID, input.ChosenMovieID)
	if newVotationError != nil {
		problemsDetails = append(problemsDetails, newVotationError...)

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		util.NewLoggerError(http.StatusInternalServerError, movieUpdatedError.Error(), "RegisterVoteUseCase", "Use Cases", util.TypeInternalServerError)

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		util.NewLoggerError(http.StatusInternalServerError, registerVoteCreationError.Error(), "RegisterVoteUseCase", "Use Cases", util.TypeInternalServerError)

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := RegisterVoteOutputDTO{
		ID:        newVotation.ID,
		Message:   "Voto registrado com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
