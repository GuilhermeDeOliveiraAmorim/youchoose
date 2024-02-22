package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type ResetVoteInputDTO struct {
	VotationID string `json:"votation_id"`
	ChooserID  string `json:"chooser_id"`
}

type ResetVoteOutputDTO struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"success"`
}

type ResetVoteUseCase struct {
	ChooserRepository  repositoryinterface.ChooserRepositoryInterface
	VotationRepository repositoryinterface.VotationRepositoryInterface
}

func NewResetVoteUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	VotationRepository repositoryinterface.VotationRepositoryInterface,
) *ResetVoteUseCase {
	return &ResetVoteUseCase{
		ChooserRepository:  ChooserRepository,
		VotationRepository: VotationRepository,
	}
}

func (rv *ResetVoteUseCase) Execute(input ResetVoteInputDTO) (ResetVoteOutputDTO, util.ProblemDetailsOutputDTO) {
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

		return ResetVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return ResetVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheVotationExist, votation, getVotationError := rv.VotationRepository.GetByID(input.VotationID)
	if getVotationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar voto de ID " + input.VotationID,
			Status:   http.StatusInternalServerError,
			Detail:   getVotationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getVotationError.Error(), "GetVotationUseCase", "Use Cases", "Internal Server Error")

		return ResetVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheVotationExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Voto não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum voto com o ID " + input.VotationID + " foi encontrado",
			Instance: util.RFC404,
		})

		return ResetVoteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	votation.Deactivate()

	votationUpdatedError := rv.VotationRepository.Update(&votation)
	if votationUpdatedError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao cancelar um voto",
			Status:   http.StatusInternalServerError,
			Detail:   votationUpdatedError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, votationUpdatedError.Error(), "UpdateChooserUseCase", "Use Cases", "Internal Server Error")
	}

	output := ResetVoteOutputDTO{
		ID:        votation.ID,
		Message:   "Voto cancelado com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
