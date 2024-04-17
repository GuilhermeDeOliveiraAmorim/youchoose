package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetListsInputDTO struct {
	ChooserID string `json:"chooser_id"`
}

type GetListsOutputDTO struct {
	Lists []ListOutputDTO `json:"lists"`
}

type GetListsUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
	ListRepository    repositoryinterface.ListRepositoryInterface
}

func NewGetListsUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ListRepository repositoryinterface.ListRepositoryInterface,
) *GetListsUseCase {
	return &GetListsUseCase{
		ChooserRepository: ChooserRepository,
		ListRepository:    ListRepository,
	}
}

func (gl *GetListsUseCase) Execute(input GetListsInputDTO) (GetListsOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(gl.ChooserRepository, input.ChooserID, "GetListsUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return GetListsOutputDTO{}, chooserValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	allLists, allListsError := gl.ListRepository.GetAll()
	if allListsError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao resgatar todas as listas",
			Status:   http.StatusInternalServerError,
			Detail:   allListsError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, allListsError.Error(), "GetListsUseCase", "Use Cases", util.TypeInternalServerError)

		return GetListsOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if len(allLists) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Listas não encontradas",
			Status:   http.StatusNotFound,
			Detail:   "Nenhuma lista foi encontrada",
			Instance: util.RFC404,
		})

		return GetListsOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	allGetListsOutputDTO := []ListOutputDTO{}

	for _, list := range allLists {
		outputList := NewListOutputDTO(list)

		allGetListsOutputDTO = append(allGetListsOutputDTO, outputList)
	}

	output := GetListsOutputDTO{
		Lists: allGetListsOutputDTO,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
