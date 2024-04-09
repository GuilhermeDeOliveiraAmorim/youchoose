package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type DeactivateListInputDTO struct {
	ChooserID string `json:"chooser_id"`
	ListID    string `json:"list_id"`
}

type DeactivateListOutputDTO struct {
	ListID    string `json:"list_id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"success"`
}

type DeactivateListUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
	ListRepository    repositoryinterface.ListRepositoryInterface
}

func NewDeactivateListUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ListRepository repositoryinterface.ListRepositoryInterface,
) *DeactivateListUseCase {
	return &DeactivateListUseCase{
		ChooserRepository: ChooserRepository,
		ListRepository:    ListRepository,
	}
}

func (dl *DeactivateListUseCase) Execute(input DeactivateListInputDTO) (DeactivateListOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, chooser, getChooserError := dl.ChooserRepository.GetByID(input.ChooserID)
	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar chooser de ID " + input.ChooserID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), "DeactivateListUseCase", "Use Cases", "Internal Server Error")

		return DeactivateListOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return DeactivateListOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return DeactivateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheListExist, list, getListError := dl.ListRepository.GetByID(input.ListID)
	if getListError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao desativar lista de ID " + input.ListID,
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
			Detail:   "Nenhuma lista com o ID " + input.ListID + " foi encontrada",
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
			Detail:   "A lista com o ID " + input.ListID + " já está desativada",
			Instance: util.RFC409,
		})

		return DeactivateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	list.Deactivate()

	listDeactivateError := dl.ListRepository.Deactivate(&list)
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
		ListID:    list.ID,
		Message:   "Lista desativada com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
