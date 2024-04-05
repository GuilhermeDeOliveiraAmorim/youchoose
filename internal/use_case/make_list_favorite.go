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
	problemsDetails := []util.ProblemDetails{}

	doesTheListExist, list, getListError := ml.ListRepository.GetByID(input.ListID)
	if getListError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao desativar lista de ID " + input.ListID,
			Status:   http.StatusInternalServerError,
			Detail:   getListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListError.Error(), "MakeListFavoriteUseCase", "Use Cases", "Internal Server Error")

		return MakeListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return MakeListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !list.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Lista não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "A lista com o ID " + input.ListID + " está desativada",
			Instance: util.RFC404,
		})

		return MakeListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheChooserExist, chooser, getChooserError := ml.ChooserRepository.GetByID(input.ChooserID)
	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar chooser de ID " + input.ChooserID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), "GetChooserUseCase", "Use Cases", "Internal Server Error")

		return MakeListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return MakeListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return MakeListFavoriteOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	listsFavorites, listsFavoritesError := ml.ListFavoriteRepository.GetAllByListID(list.ID)
	if listsFavoritesError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao desativar uma lista",
			Status:   http.StatusInternalServerError,
			Detail:   listsFavoritesError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listsFavoritesError.Error(), "MakeListFavoriteUseCase", "Use Cases", "Internal Server Error")
	}

	newListFavorite := entity.NewListFavorite(input.ChooserID, input.ListID)

	for _, listFavorite := range listsFavorites {
		if listFavorite.Equals(newListFavorite) {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Conflict",
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

	listFavoriteError := ml.ListFavoriteRepository.Create(newListFavorite)
	if listFavoriteError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao desativar uma lista",
			Status:   http.StatusInternalServerError,
			Detail:   listFavoriteError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listFavoriteError.Error(), "MakeListFavoriteUseCase", "Use Cases", "Internal Server Error")

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
