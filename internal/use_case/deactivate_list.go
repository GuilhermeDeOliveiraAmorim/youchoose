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
	_, chooserValidatorProblems := chooserValidator(dl.ChooserRepository, input.ChooserID, "DeactivateListInputDTO")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return DeactivateListOutputDTO{}, chooserValidatorProblems
	}

	list, listValidatorProblems := listValidator(dl.ListRepository, input.ListID, "DeactivateListInputDTO")
	if len(listValidatorProblems.ProblemDetails) > 0 {
		return DeactivateListOutputDTO{}, listValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

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
