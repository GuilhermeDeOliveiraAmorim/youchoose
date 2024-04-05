package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetListsInputDTO struct{}

type ListOutputDTO struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	ChooserID      string `json:"chooser_id"`
	ProfileImageID string `json:"profile_image_id"`
	CoverImageID   string `json:"cover_image_id"`
	Votes          int    `json:"votes"`
}

type GetListsOutputDTO struct {
	Lists []ListOutputDTO `json:"lists"`
}

type GetListsUseCase struct {
	ListRepository repositoryinterface.ListRepositoryInterface
}

func NewGetListsUseCase(
	ListRepository repositoryinterface.ListRepositoryInterface,
) *GetListsUseCase {
	return &GetListsUseCase{
		ListRepository: ListRepository,
	}
}

func (gl *GetListsUseCase) Execute(input GetListsInputDTO) (GetListsOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	allLists, allListsError := gl.ListRepository.GetAll()
	if allListsError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar todas as listas",
			Status:   http.StatusInternalServerError,
			Detail:   allListsError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, allListsError.Error(), "GetListsUseCase", "Use Cases", "Internal Server Error")

		return GetListsOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if len(allLists) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
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
		allGetListsOutputDTO = append(allGetListsOutputDTO, ListOutputDTO{
			ID:             list.ID,
			Title:          list.Title,
			Description:    list.Description,
			ChooserID:      list.ChooserID,
			ProfileImageID: list.ProfileImageID,
			CoverImageID:   list.CoverImageID,
			Votes:          list.Votes,
		})
	}

	output := GetListsOutputDTO{
		Lists: allGetListsOutputDTO,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
