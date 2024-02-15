package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetChoosersInputDTO struct{}

type ChooserOutputDTO struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Day     int    `json:"day"`
	Month   int    `json:"month"`
	Year    int    `json:"year"`
	ImageID string `json:"image_id"`
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

func (cc *GetChoosersUseCase) Execute(input GetChoosersInputDTO) (GetChoosersOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	allChoosers, allChoosersError := cc.ChooserRepository.GetAll()
	if allChoosersError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar todos os choosers",
			Status:   http.StatusInternalServerError,
			Detail:   allChoosersError.Error(),
			Instance: util.RFC500,
		})

		util.NewLoggerError(http.StatusInternalServerError, allChoosersError.Error(), "GetChoosersUseCase", "Use Cases", "Internal Server Error")

		return GetChoosersOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if len(allChoosers) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Choosers não encontrados",
			Status:   http.StatusInternalServerError,
			Detail:   "Nenhum chooser foi encontrado",
			Instance: util.RFC404,
		})

		return GetChoosersOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	allGetChoosersOutputDTO := []ChooserOutputDTO{}

	for _, chooser := range allChoosers {
		allGetChoosersOutputDTO = append(allGetChoosersOutputDTO, ChooserOutputDTO{
			ID:      chooser.ID,
			Name:    chooser.Name,
			City:    chooser.Address.City,
			State:   chooser.Address.State,
			Country: chooser.Address.Country,
			Day:     chooser.BirthDate.Day,
			Month:   chooser.BirthDate.Month,
			Year:    chooser.BirthDate.Year,
			ImageID: chooser.ImageID,
		})
	}

	output := GetChoosersOutputDTO{
		Choosers: allGetChoosersOutputDTO,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
