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
	ChooserID string `json:"chooser_id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	ImageID   string `json:"image_id"`
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

func (uc *UpdateChooserUseCase) Execute(input UpdateChooserInputDTO) (ChooserOutputDTO, util.ProblemDetailsOutputDTO) {
	chooser, chooserValidatorProblems := chooserValidator(uc.ChooserRepository, input.ChooserID, "UpdateChooserUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return ChooserOutputDTO{}, chooserValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

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
		return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	if chooser.Name != input.Name {
		chooser.ChangeName(ctx, input.Name)
	}

	if !chooser.Address.Equals(newAddress) {
		chooser.ChangeAddress(ctx, newAddress)
	}

	if !chooser.BirthDate.Equals(newBirthdate) {
		chooser.ChangeBirthDate(ctx, newBirthdate)
	}

	if chooser.ImageID != input.ImageID {
		chooser.ChangeImageID(ctx, input.ImageID)
	}

	chooserUpdatedError := uc.ChooserRepository.Update(&chooser)
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

	output := NewChooserOutputDTO(chooser)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
