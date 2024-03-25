package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type DeactivateListInputDTO struct {
	ID string `json:"id"`
}

type DeactivateListOutputDTO struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"success"`
}

type DeactivateListUseCase struct {
	ListRepository repositoryinterface.ListRepositoryInterface
}

func NewDeactivateListUseCase(
	ListRepository repositoryinterface.ListRepositoryInterface,
) *DeactivateListUseCase {
	return &DeactivateListUseCase{
		ListRepository: ListRepository,
	}
}

func (cc *DeactivateListUseCase) Execute(input DeactivateListInputDTO) (DeactivateListOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheListExist, list, getListError := cc.ListRepository.GetByID(input.ID)
	if getListError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao desativar lista de ID " + input.ID,
			Status:   http.StatusInternalServerError,
			Detail:   getListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListError.Error(), "DeactivateListUseCase", "Use Cases", "Internal Server Error")

		return DeactivateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheListExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Lista não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "Nenhuma lista com o ID " + input.ID + " foi encontrada",
			Instance: util.RFC404,
		})

		return DeactivateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !list.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Conflict",
			Title:    "Lista já está desativada",
			Status:   http.StatusConflict,
			Detail:   "A lista com o ID " + input.ID + " já está desativada",
			Instance: util.RFC409,
		})

		return DeactivateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	listDeactivateError := cc.ListRepository.Deactivate(list.ID)
	if listDeactivateError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao desativar uma lista",
			Status:   http.StatusInternalServerError,
			Detail:   listDeactivateError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listDeactivateError.Error(), "DeactivateListUseCase", "Use Cases", "Internal Server Error")
	}

	output := DeactivateListOutputDTO{
		ID:        list.ID,
		Message:   "Lista desativada com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
