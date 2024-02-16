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
	VotationID    string `json:"currentVotation_id"`
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
	ChooserRepository     repositoryinterface.ChooserRepositoryInterface
	VotationRepository    repositoryinterface.VotationRepositoryInterface
	ListRepository        repositoryinterface.ListRepositoryInterface
	CombinationRepository repositoryinterface.CombinationRepositoryInterface
}

func NewRegisterVoteUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
) *RegisterVoteUseCase {
	return &RegisterVoteUseCase{
		ChooserRepository: ChooserRepository,
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

	doesTheVotationExist, currentVotation, getVotationError := rv.VotationRepository.GetByID(input.VotationID)
	if getVotationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar votação de ID " + input.VotationID,
			Status:   http.StatusInternalServerError,
			Detail:   getVotationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getVotationError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !currentVotation.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Votação não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "Nenhuma votação com o ID " + input.VotationID + " foi encontrada",
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

	getAllCombinationsByVotationID, getAllCombinationsByVotationIDError := rv.CombinationRepository.GetAllByVotationID(input.VotationID)
	if getAllCombinationsByVotationIDError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao carregar informações da lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   getListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getAllCombinationsByVotationIDError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if len(getAllCombinationsByVotationID) == 0 && currentVotation.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Erro ao carregar informações da lista de ID " + input.ListID,
			Status:   http.StatusNotFound,
			Detail:   "Nenhuma combinação de votos foi encontrada para a lista com o ID " + input.ListID,
			Instance: util.RFC404,
		})

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	if !doesTheVotationExist {
		newVotation, newVotationProblems := entity.NewVotation(input.ChooserID, input.ListID)
		if len(newVotationProblems) > 0 {
			return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: newVotationProblems,
			}
		}

		newCombination, newCombinationProblems := entity.NewCombination(newVotation.ID, input.FirstMovieID, input.SecondMovieID, input.ChosenMovieID)
		if len(newCombinationProblems) > 0 {
			return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: newCombinationProblems,
			}
		}

		for _, combination := range getAllCombinationsByVotationID {
			if combination.Equals(newCombination) {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     "Conflict",
					Title:    "Combinação de votos já cadastrada",
					Status:   http.StatusConflict,
					Detail:   "Já existe uma combinação de votos cadastrada para a lista com o ID " + input.ListID,
					Instance: util.RFC409,
				})

				return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}
			}
		}

		newVotation.Vote(*newCombination)
		newVotation.AddCombination(newCombination)

		newVotation.StartVotation()

		createVotationError := rv.VotationRepository.Create(newVotation)
		if createVotationError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Internal Server Error",
				Title:    "Erro ao persistir votação",
				Status:   http.StatusInternalServerError,
				Detail:   createVotationError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, createVotationError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")

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

	newCombination, newCombinationProblems := entity.NewCombination(input.VotationID, input.FirstMovieID, input.SecondMovieID, input.ChosenMovieID)
	if len(newCombinationProblems) > 0 {
		problemsDetails = append(problemsDetails, newCombinationProblems...)

		return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	for _, combination := range getAllCombinationsByVotationID {
		if combination.Equals(newCombination) {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Conflict",
				Title:    "Combinação de votos já cadastrada",
				Status:   http.StatusConflict,
				Detail:   "Já existe uma combinação de votos cadastrada para a lista com o ID " + input.ListID,
				Instance: util.RFC409,
			})

			return RegisterVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}
	}

	currentVotation.AddCombination(newCombination)

	combinationCreationError := rv.CombinationRepository.Create(newCombination)
	if combinationCreationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir um voto",
			Status:   http.StatusInternalServerError,
			Detail:   combinationCreationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, combinationCreationError.Error(), "RegisterVoteUseCase", "Use Cases", "Internal Server Error")
	}

	output := RegisterVoteOutputDTO{
		ID:        input.VotationID,
		Message:   "Voto registrado com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
