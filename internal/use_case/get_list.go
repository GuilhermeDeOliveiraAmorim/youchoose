package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetListInputDTO struct {
	ID string `json:"id"`
}

type GetListOutputDTO struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	ChooserID      string `json:"chooser_id"`
	ProfileImageID string `json:"profile_image_id"`
	CoverImageID   string `json:"cover_image_id"`
	Votes          int    `json:"votes"`
}

type GetListUseCase struct {
	ListRepository repositoryinterface.ListRepositoryInterface
}

func NewGetListUseCase(
	ListRepository repositoryinterface.ListRepositoryInterface,
) *GetListUseCase {
	return &GetListUseCase{
		ListRepository: ListRepository,
	}
}

func (gl *GetListUseCase) Execute(input GetListInputDTO) (GetListOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheListExist, list, getListError := gl.ListRepository.GetByID(input.ID)
	if getListError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar lista de ID " + input.ID,
			Status:   http.StatusInternalServerError,
			Detail:   getListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListError.Error(), "GetListUseCase", "Use Cases", "Internal Server Error")

		return GetListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheListExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Lista não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "Nenhuma lista com o ID " + input.ID + " foi encontrada",
			Instance: util.RFC404,
		})

		return GetListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := GetListOutputDTO{
		ID:             list.ID,
		Title:          list.Title,
		Description:    list.Description,
		ChooserID:      list.ChooserID,
		ProfileImageID: list.ProfileImageID,
		CoverImageID:   list.CoverImageID,
		Votes:          list.Votes,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
