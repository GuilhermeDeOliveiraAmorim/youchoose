package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type DeactivateChooserInputDTO struct {
	ChooserID             string `json:"chooser_id"`
	ChooserIDToDeactivate string `json:"chooser_id_to_deactivate"`
}

type DeactivateChooserOutputDTO struct {
	ChooserID string `json:"chooser_id"`
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

func (dc *DeactivateChooserUseCase) Execute(input DeactivateChooserInputDTO) (DeactivateChooserOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(dc.ChooserRepository, input.ChooserID, "DeactivateChooserUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return DeactivateChooserOutputDTO{}, chooserValidatorProblems
	}

	chooserToDeactivate, chooserValidatorProblems := chooserValidator(dc.ChooserRepository, input.ChooserIDToDeactivate, "DeactivateChooserUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return DeactivateChooserOutputDTO{}, chooserValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	chooserToDeactivate.Deactivate()

	chooserDeactivateError := dc.ChooserRepository.Deactivate(&chooserToDeactivate)
	if chooserDeactivateError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao desativar um chooser",
			Status:   http.StatusInternalServerError,
			Detail:   chooserDeactivateError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, chooserDeactivateError.Error(), "DeactivateChooserUseCase", "Use Cases", util.TypeInternalServerError)
	}

	output := DeactivateChooserOutputDTO{
		ChooserID: chooserToDeactivate.ID,
		Message:   "Chooser desativado com sucesso",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
