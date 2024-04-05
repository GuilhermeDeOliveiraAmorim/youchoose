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
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, chooser, getChooserError := rv.ChooserRepository.GetByID(input.ChooserID)
	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar chooser de ID " + input.ChooserID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheListExist, list, getListError := rv.ListRepository.GetByID(input.ListID)
	if getListError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   getListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheListExist || !list.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Lista não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "Nenhuma lista com o ID " + input.ListID + " foi encontrada",
			Instance: util.RFC404,
		})

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	var moviesIDs []string

	moviesIDs = append(moviesIDs, input.FirstMovieID)
	moviesIDs = append(moviesIDs, input.SecondMovieID)
	moviesIDs = append(moviesIDs, input.ChosenMovieID)

	doesTheMoviesExist, movies, doesTheMoviesExistError := rv.MovieRepository.DoTheseMoviesExist(moviesIDs)
	if doesTheMoviesExistError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   doesTheMoviesExistError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, doesTheMoviesExistError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheMoviesExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Validation Error",
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
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   doesTheVotationExistError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, doesTheVotationExistError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if doesTheVotationExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Already Exists",
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
			Type:     "Internal Server Error",
			Title:    "Erro ao atualizar filme de ID " + movies[2].ID,
			Status:   http.StatusInternalServerError,
			Detail:   movieUpdatedError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, movieUpdatedError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	registerVoteCreationError := rv.VotationRepository.Create(newVotation)
	if registerVoteCreationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir uma votação",
			Status:   http.StatusInternalServerError,
			Detail:   registerVoteCreationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, registerVoteCreationError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")

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
