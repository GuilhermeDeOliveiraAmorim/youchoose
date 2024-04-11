package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetChoosersInputDTO struct {
	ChooserID string `json:"chooser_id"`
}

type GetChoosersOutputDTO struct {
	Choosers []ChooserOutputDTO `json:"choosers"`
}

type GetChoosersUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
}

func NewGetChoosersUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
) *GetChoosersUseCase {
	return &GetChoosersUseCase{
		ChooserRepository: ChooserRepository,
	}
}

func (gc *GetChoosersUseCase) Execute(input GetChoosersInputDTO) (GetChoosersOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(gc.ChooserRepository, input.ChooserID, "GetChoosersUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return GetChoosersOutputDTO{}, chooserValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	allChoosers, allChoosersError := gc.ChooserRepository.GetAll()
	if allChoosersError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar todos os choosers",
			Status:   http.StatusInternalServerError,
			Detail:   allChoosersError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, allChoosersError.Error(), "GetChoosersUseCase", "Use Cases", "Internal Server Error")

		return GetChoosersOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if len(allChoosers) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Choosers não encontrados",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum chooser foi encontrado",
			Instance: util.RFC404,
		})

		return GetChoosersOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	allGetChoosersOutputDTO := []ChooserOutputDTO{}

	for _, chooser := range allChoosers {
		outputChooser := NewChooserOutputDTO(chooser)

		allGetChoosersOutputDTO = append(allGetChoosersOutputDTO, outputChooser)
	}

	output := GetChoosersOutputDTO{
		Choosers: allGetChoosersOutputDTO,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
