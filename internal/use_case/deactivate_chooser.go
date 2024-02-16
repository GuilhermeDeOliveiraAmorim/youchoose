package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type DeactivateChooserInputDTO struct {
	ID string `json:"id"`
}

type DeactivateChooserOutputDTO struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"success"`
}

type DeactivateChooserUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
}

func NewDeactivateChooserUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
) *DeactivateChooserUseCase {
	return &DeactivateChooserUseCase{
		ChooserRepository: ChooserRepository,
	}
}

func (cc *DeactivateChooserUseCase) Execute(input DeactivateChooserInputDTO) (DeactivateChooserOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, chooser, getChooserError := cc.ChooserRepository.GetByID(input.ID)
	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao desativar chooser de ID " + input.ID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), "DeactivateChooserUseCase", "Use Cases", "Internal Server Error")

		return DeactivateChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheChooserExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum chooser com o ID " + input.ID + " foi encontrado",
			Instance: util.RFC404,
		})

		return DeactivateChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	chooserDeactivateError := cc.ChooserRepository.Deactivate(chooser.ID)
	if chooserDeactivateError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao desativar um chooser",
			Status:   http.StatusInternalServerError,
			Detail:   chooserDeactivateError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, chooserDeactivateError.Error(), "DeactivateChooserUseCase", "Use Cases", "Internal Server Error")
	}

	output := DeactivateChooserOutputDTO{
		ID:        chooser.ID,
		Message:   "Chooser desativado com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
