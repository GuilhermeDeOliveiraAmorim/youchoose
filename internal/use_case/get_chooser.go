package usecase

import (
	"net/http"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetChooserInputDTO struct {
	ID string `json:"id"`
}

type GetChooserOutputDTO struct {
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

type GetChooserUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
}

func NewGetChooserUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
) *GetChooserUseCase {
	return &GetChooserUseCase{
		ChooserRepository: ChooserRepository,
	}
}

func (cc *GetChooserUseCase) Execute(input GetChooserInputDTO) (GetChooserOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, chooser, getChooserError := cc.ChooserRepository.GetByID(input.ID)
	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar chooser de ID " + input.ID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), "GetChooserUseCase", "Use Cases", "Internal Server Error")

		return GetChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheChooserExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum chooser com o ID " + input.ID + " foi encontrado",
			Instance: util.RFC404,
		})

		return GetChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !chooser.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "O chooser com o ID " + input.ID + " está desativado",
			Instance: util.RFC404,
		})

		return GetChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := GetChooserOutputDTO{
		ID:      chooser.ID,
		Name:    chooser.Name,
		City:    chooser.Address.City,
		State:   chooser.Address.State,
		Country: chooser.Address.Country,
		Day:     chooser.BirthDate.Day,
		Month:   chooser.BirthDate.Month,
		Year:    chooser.BirthDate.Year,
		ImageID: chooser.ImageID,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
