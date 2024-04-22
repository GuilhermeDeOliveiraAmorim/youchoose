package usecase

import (
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type MakeListFavoriteInputDTO struct {
	ChooserID string `json:"chooser_id"`
	ListID    string `json:"list_id"`
}

type MakeListFavoriteOutputDTO struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"is_success"`
}

type MakeListFavoriteUseCase struct {
	ListRepository         repositoryinterface.ListRepositoryInterface
	ChooserRepository      repositoryinterface.ChooserRepositoryInterface
	ListFavoriteRepository repositoryinterface.ListFavoriteRepositoryInterface
}

func NewMakeListFavoriteUseCase(
	ListRepository repositoryinterface.ListRepositoryInterface,
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ListFavoriteRepository repositoryinterface.ListFavoriteRepositoryInterface,
) *MakeListFavoriteUseCase {
	return &MakeListFavoriteUseCase{
		ListRepository:         ListRepository,
		ChooserRepository:      ChooserRepository,
		ListFavoriteRepository: ListFavoriteRepository,
	}
}

func (ml *MakeListFavoriteUseCase) Execute(input MakeListFavoriteInputDTO) (MakeListFavoriteOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(ml.ChooserRepository, input.ChooserID, "MakeListFavoriteUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return MakeListFavoriteOutputDTO{}, chooserValidatorProblems
	}

	list, listValidatorProblems := listValidator(ml.ListRepository, input.ListID, "MakeListFavoriteUseCase")
	if len(listValidatorProblems.ProblemDetails) > 0 {
		return MakeListFavoriteOutputDTO{}, listValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	listsFavorites, listsFavoritesError := ml.ListFavoriteRepository.GetAllByListID(list.ID)
	if listsFavoritesError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao desativar uma lista",
			Status:   http.StatusInternalServerError,
			Detail:   listsFavoritesError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listsFavoritesError.Error(), "MakeListFavoriteUseCase", "Use Cases", util.TypeInternalServerError)
	}

	newListFavorite := entity.NewListFavorite(input.ChooserID, input.ListID)

	if len(listsFavorites) > 0 {
		for _, listFavorite := range listsFavorites {
			if listFavorite.Equals(newListFavorite) {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeConflict,
					Title:    "Lista já favorita",
					Status:   http.StatusConflict,
					Detail:   "A lista com o ID " + input.ListID + " já é favorita",
					Instance: util.RFC409,
				})

				return MakeListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}
			}
		}
	}

	listFavoriteError := ml.ListFavoriteRepository.Create(newListFavorite)
	if listFavoriteError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao favoritar uma lista",
			Status:   http.StatusInternalServerError,
			Detail:   listFavoriteError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listFavoriteError.Error(), "MakeListFavoriteUseCase", "Use Cases", util.TypeInternalServerError)

		return MakeListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	list.IncrementVotes()

	listError := ml.ListRepository.Update(&list)
	if listError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao favoritar uma lista",
			Status:   http.StatusInternalServerError,
			Detail:   listError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listError.Error(), "MakeListFavoriteUseCase", "Use Cases", util.TypeInternalServerError)

		return MakeListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := MakeListFavoriteOutputDTO{
		ID:        list.ID,
		Message:   "Lista acrescentada às favoritas",
		IsSuccess: true,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
