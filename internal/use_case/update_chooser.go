package usecase

import (
	"context"
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"
)

type UpdateChooserInputDTO struct {
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

type UpdateChooserOutputDTO struct {
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

type UpdateChooserUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
}

func NewUpdateChooserUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
) *UpdateChooserUseCase {
	return &UpdateChooserUseCase{
		ChooserRepository: ChooserRepository,
	}
}

func (uc *UpdateChooserUseCase) Execute(input UpdateChooserInputDTO) (UpdateChooserOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, userThatExists, doesTheChooserExistError := uc.ChooserRepository.GetByID(input.ID)
	if doesTheChooserExistError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao buscar um chooser",
			Status:   http.StatusInternalServerError,
			Detail:   doesTheChooserExistError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, doesTheChooserExistError.Error(), "UpdateChooserUseCase", "Use Cases", "Internal Server Error")

		return UpdateChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheChooserExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Recurso não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Não foi possível encontrar o chooser de id " + input.ID,
			Instance: util.RFC404,
		})

		return UpdateChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !userThatExists.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusForbidden,
			Detail:   "Não foi possível encontrar o chooser de id " + input.ID,
			Instance: util.RFC404,
		})

		return UpdateChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	ctx := context.Background()

	validateChooserProblems := entity.ValidateChooser(input.Name, input.ImageID)
	if len(validateChooserProblems) > 0 {
		problemsDetails = append(problemsDetails, validateChooserProblems...)
	}

	newAddress, newAddressProblems := valueobject.NewAddress(input.City, input.State, input.Country)
	if len(newAddressProblems) > 0 {
		problemsDetails = append(problemsDetails, newAddressProblems...)
	}

	newBirthdate, newBirthdateProblems := valueobject.NewBirthDate(input.Day, input.Month, input.Year)
	if len(newBirthdateProblems) > 0 {
		problemsDetails = append(problemsDetails, newBirthdateProblems...)
	}

	if len(problemsDetails) > 0 {
		return UpdateChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	if userThatExists.Name != input.Name {
		userThatExists.ChangeName(ctx, input.Name)
	}

	if !userThatExists.Address.Equals(newAddress) {
		userThatExists.ChangeAddress(ctx, newAddress)
	}

	if !userThatExists.BirthDate.Equals(newBirthdate) {
		userThatExists.ChangeBirthDate(ctx, newBirthdate)
	}

	if userThatExists.ImageID != input.ImageID {
		userThatExists.ChangeImageID(ctx, input.ImageID)
	}

	chooserUpdatedError := uc.ChooserRepository.Update(&userThatExists)
	if chooserUpdatedError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir um chooser",
			Status:   http.StatusInternalServerError,
			Detail:   chooserUpdatedError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, chooserUpdatedError.Error(), "UpdateChooserUseCase", "Use Cases", "Internal Server Error")
	}

	output := UpdateChooserOutputDTO{
		ID:      userThatExists.ID,
		Name:    userThatExists.Name,
		City:    userThatExists.Address.City,
		State:   userThatExists.Address.State,
		Country: userThatExists.Address.Country,
		Day:     userThatExists.BirthDate.Day,
		Month:   userThatExists.BirthDate.Month,
		Year:    userThatExists.BirthDate.Year,
		ImageID: userThatExists.ImageID,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
