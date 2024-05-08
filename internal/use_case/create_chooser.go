package usecase

import (
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"
)

type CreateChooserInputDTO struct {
	ChooserID string `json:"chooser_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	ImageID   string `json:"image_id"`
}

type CreateChooserUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
}

func NewCreateChooserUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
) *CreateChooserUseCase {
	return &CreateChooserUseCase{
		ChooserRepository: ChooserRepository,
	}
}

func (cc *CreateChooserUseCase) Execute(input CreateChooserInputDTO) (ChooserOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(cc.ChooserRepository, input.ChooserID, "CreateChooserUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return ChooserOutputDTO{}, chooserValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	allChoosers, allChoosersProblem := cc.ChooserRepository.GetAll()
	if allChoosersProblem != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao resgatar todos os choosers",
			Status:   http.StatusInternalServerError,
			Detail:   allChoosersProblem.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar todos os choosers", "CreateChooserUseCase", "Use Cases", util.TypeInternalServerError)

		return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	for _, allChooser := range allChoosers {
		if allChooser.Login.DecryptEmail(allChooser.Login.Email, input.Email) == input.Email {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeValidationError,
				Title:    "E-mail já está em uso",
				Status:   http.StatusConflict,
				Detail:   "O e-mail fornecido já está sendo utilizado por outro chooser.",
				Instance: util.RFC409,
			})

			return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}
	}

	newLogin, newLoginProblems := valueobject.NewLogin(input.Email, input.Password)
	if len(newLoginProblems) > 0 {
		problemsDetails = append(problemsDetails, newLoginProblems...)
	}

	newAddress, newAddressProblems := valueobject.NewAddress(input.City, input.State, input.Country)
	if len(newAddressProblems) > 0 {
		problemsDetails = append(problemsDetails, newAddressProblems...)
	}

	newBirthdate, newBirthdateProblems := valueobject.NewBirthDate(input.Day, input.Month, input.Year)
	if len(newBirthdateProblems) > 0 {
		problemsDetails = append(problemsDetails, newBirthdateProblems...)
	}

	newChooser, newChooserProblems := entity.NewChooser(input.Name, newLogin, newAddress, newBirthdate, input.ImageID)
	if len(newChooserProblems) > 0 {
		problemsDetails = append(problemsDetails, newChooserProblems...)
	}

	encryptPassword, encryptPasswordProblems := newChooser.Login.EncryptPassword(newLogin.Password)
	if encryptPasswordProblems != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.LoginErrorTitleEncryptPassword,
			Status:   http.StatusConflict,
			Detail:   util.LoginErrorDetailInvalidPassword,
			Instance: util.RFC409,
		})
	}

	encryptEmail, encryptEmailProblems := newChooser.Login.EncryptEmail(newLogin.Email)
	if encryptEmailProblems != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.LoginErrorTitleEncryptEmail,
			Status:   http.StatusConflict,
			Detail:   util.LoginErrorDetailInvalidEmail,
			Instance: util.RFC409,
		})
	}

	if len(problemsDetails) > 0 {
		return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	newChooser.Login.ChangeEmail(encryptEmail)
	newChooser.Login.ChangePassword(encryptPassword)

	chooserCreationError := cc.ChooserRepository.Create(newChooser)
	if chooserCreationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao persistir um chooser",
			Status:   http.StatusInternalServerError,
			Detail:   chooserCreationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, chooserCreationError.Error(), "CreateChooserUseCase", "Use Cases", util.TypeInternalServerError)

		return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := NewChooserOutputDTO(*newChooser)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
