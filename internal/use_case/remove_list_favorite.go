package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type RemoveListFavoriteInputDTO struct {
	ChooserID string `json:"chooser_id"`
	ListID    string `json:"list_id"`
}

type RemoveListFavoriteOutputDTO struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"is_success"`
}

type RemoveListFavoriteUseCase struct {
	ChooserRepository      repositoryinterface.ChooserRepositoryInterface
	ListRepository         repositoryinterface.ListRepositoryInterface
	ListFavoriteRepository repositoryinterface.ListFavoriteRepositoryInterface
}

func NewRemoveListFavoriteUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ListRepository repositoryinterface.ListRepositoryInterface,
	ListFavoriteRepository repositoryinterface.ListFavoriteRepositoryInterface,
) *RemoveListFavoriteUseCase {
	return &RemoveListFavoriteUseCase{
		ChooserRepository:      ChooserRepository,
		ListRepository:         ListRepository,
		ListFavoriteRepository: ListFavoriteRepository,
	}
}

func (rl *RemoveListFavoriteUseCase) Execute(input RemoveListFavoriteInputDTO) (RemoveListFavoriteOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(rl.ChooserRepository, input.ChooserID, "RemoveListFavoriteUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return RemoveListFavoriteOutputDTO{}, chooserValidatorProblems
	}

	list, listValidatorProblems := listValidator(rl.ListRepository, input.ListID, "RemoveListFavoriteUseCase")
	if len(listValidatorProblems.ProblemDetails) > 0 {
		return RemoveListFavoriteOutputDTO{}, listValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	doesTheListFavoriteExist, listFavorite, getListFavoriteError := rl.ListFavoriteRepository.GetByChooserIDAndListID(input.ChooserID, input.ListID)
	if getListFavoriteError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao remover lista das favoritas",
			Status:   http.StatusInternalServerError,
			Detail:   getListFavoriteError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListFavoriteError.Error(), "RemoveListFavoriteUseCase", "Use Cases", util.TypeInternalServerError)

		return RemoveListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheListFavoriteExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Erro ao remover lista das favoritas",
			Status:   http.StatusNotFound,
			Detail:   "Não foi possível remover a lista das favoritas",
			Instance: util.RFC404,
		})

		return RemoveListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !listFavorite.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeNotFound,
			Title:    "Lista já removida",
			Status:   http.StatusNotFound,
			Detail:   "A lista já está removida das favoritas",
			Instance: util.RFC404,
		})

		return RemoveListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	listFavorite.Deactivate()

	listFavoriteError := rl.ListFavoriteRepository.Deactivate(&listFavorite)
	if listFavoriteError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao remover lista das favoritas",
			Status:   http.StatusInternalServerError,
			Detail:   listFavoriteError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listFavoriteError.Error(), "RemoveListFavoriteUseCase", "Use Cases", util.TypeInternalServerError)

		return RemoveListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	list.DecrementVotes()

	listError := rl.ListRepository.Update(&list)
	if listError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao remover lista das favoritas",
			Status:   http.StatusInternalServerError,
			Detail:   listError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listError.Error(), "RemoveListFavoriteUseCase", "Use Cases", util.TypeInternalServerError)

		return RemoveListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := RemoveListFavoriteOutputDTO{
		ID:        listFavorite.ID,
		Message:   "Lista removida das favoritas",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
