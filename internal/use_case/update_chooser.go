package usecase

import (
	"context"
	"mime/multipart"
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/service"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"
)

type UpdateChooserInputDTO struct {
	ChooserID    string                `json:"chooser_id"`
	Name         string                `json:"name"`
	City         string                `json:"city"`
	State        string                `json:"state"`
	Country      string                `json:"country"`
	Day          int                   `json:"day"`
	Month        int                   `json:"month"`
	Year         int                   `json:"year"`
	ImageID      string                `json:"image_id"`
	ImageFile    multipart.File        `json:"chooser_image_file"`
	ImageHandler *multipart.FileHeader `json:"chooser_image_handler"`
}

type UpdateChooserUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
	ImageRepository   repositoryinterface.ImageRepositoryInterface
}

func NewUpdateChooserUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ImageRepository repositoryinterface.ImageRepositoryInterface,
) *UpdateChooserUseCase {
	return &UpdateChooserUseCase{
		ChooserRepository: ChooserRepository,
		ImageRepository:   ImageRepository,
	}
}

func (uc *UpdateChooserUseCase) Execute(input UpdateChooserInputDTO) (ChooserOutputDTO, util.ProblemDetailsOutputDTO) {
	chooser, chooserValidatorProblems := chooserValidator(uc.ChooserRepository, input.ChooserID, "UpdateChooserUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return ChooserOutputDTO{}, chooserValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	ctx := context.Background()

	newAddress, newAddressProblems := valueobject.NewAddress(input.City, input.State, input.Country)
	if len(newAddressProblems) > 0 {
		problemsDetails = append(problemsDetails, newAddressProblems...)
	}

	newBirthdate, newBirthdateProblems := valueobject.NewBirthDate(input.Day, input.Month, input.Year)
	if len(newBirthdateProblems) > 0 {
		problemsDetails = append(problemsDetails, newBirthdateProblems...)
	}

	if (input.ImageID == "" && (input.ImageFile == nil || input.ImageHandler == nil)) || (input.ImageID != "" && (input.ImageID != chooser.ImageID)) {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Imagem não informada",
			Status:   http.StatusBadRequest,
			Detail:   "O chooser deve ter uma imagem",
			Instance: util.RFC400,
		})
	}

	if len(problemsDetails) > 0 {
		return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	var imagesToAdd []entity.Image

	if input.ImageID == "" {
		_, newChooserImageProblemName, newChooserImageProblemExtension, newChooserImageProblemSize, newChooserImageProblemError := service.MoveFile(input.ImageFile, input.ImageHandler)
		if newChooserImageProblemError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao mover a imagem do chooser",
				Status:   http.StatusInternalServerError,
				Detail:   newChooserImageProblemError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do chooser", "UpdateChooserUseCase", "Use Cases", util.TypeInternalServerError)

			return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		newChooserImage, newChooserImageProblem := entity.NewImage(newChooserImageProblemName, newChooserImageProblemExtension, newChooserImageProblemSize)
		if len(newChooserImageProblem) > 0 {
			return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: newChooserImageProblem,
			}
		}

		imagesToAdd = append(imagesToAdd, *newChooserImage)
	} else {
		validateChooserProblems := entity.ValidateChooser(input.Name, input.ImageID)
		if len(validateChooserProblems) > 0 {
			problemsDetails = append(problemsDetails, validateChooserProblems...)
		}
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

	if len(imagesToAdd) > 0 {
		imagesToAddError := uc.ImageRepository.Create(&imagesToAdd[0])
		if imagesToAddError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao atualizar imagem",
				Status:   http.StatusInternalServerError,
				Detail:   imagesToAddError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao atualizar a imagem do chooser", "UpdateChooserUseCase", "Use Cases", util.TypeInternalServerError)

			return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		doesTheImageExist, imageToDeactivate, getImageError := uc.ImageRepository.GetByID(chooser.ImageID)
		if getImageError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao buscar imagem",
				Status:   http.StatusInternalServerError,
				Detail:   getImageError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao buscar a imagem antiga do chooser", "UpdateChooserUseCase", "Use Cases", util.TypeInternalServerError)

			return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		} else if !doesTheImageExist {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeNotFound,
				Title:    util.SharedErrorTitleNotFound,
				Status:   http.StatusNotFound,
				Detail:   util.ImageErrorDetailNotFound,
				Instance: util.RFC404,
			})
		} else if !chooser.Active {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeNotFound,
				Title:    util.SharedErrorTitleNotFound,
				Status:   http.StatusNotFound,
				Detail:   util.ImageErrorDetailDeactivate,
				Instance: util.RFC404,
			})
		}

		imageToDeactivate.Deactivate()

		imageDeactivatedError := uc.ImageRepository.Deactivate(&imageToDeactivate)
		if imageDeactivatedError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao desativar imagem",
				Status:   http.StatusInternalServerError,
				Detail:   imageDeactivatedError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao desativar a imagem antiga do chooser", "UpdateChooserUseCase", "Use Cases", util.TypeInternalServerError)

			return ChooserOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		chooser.ChangeImage(ctx, imagesToAdd[0].ID)
	}

	chooserUpdatedError := uc.ChooserRepository.Update(&chooser)
	if chooserUpdatedError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao atualizar chooser",
			Status:   http.StatusInternalServerError,
			Detail:   chooserUpdatedError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, chooserUpdatedError.Error(), "UpdateChooserUseCase", "Use Cases", util.TypeInternalServerError)
	}

	output := NewChooserOutputDTO(chooser)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
