package usecase

import (
	"mime/multipart"
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/service"
	"youchoose/internal/util"
)

type UpdateListInputDTO struct {
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	ProfileImageFile    multipart.File        `json:"profile_image_file"`
	ProfileImageHandler *multipart.FileHeader `json:"profile_image_handler"`
	CoverImageFile      multipart.File        `json:"cover_image_file"`
	CoverImageHandler   *multipart.FileHeader `json:"cover_image_handler"`
	Description         string                `json:"description"`
	Movies              []string              `json:"movies"`
	ChooserID           string                `json:"chooser_id"`
}

type UpdateListUseCase struct {
	ListRepository      repositoryinterface.ListRepositoryInterface
	ChooserRepository   repositoryinterface.ChooserRepositoryInterface
	ImageRepository     repositoryinterface.ImageRepositoryInterface
	MovieRepository     repositoryinterface.MovieRepositoryInterface
	ListMovieRepository repositoryinterface.ListMovieRepositoryInterface
}

func NewUpdateListUseCase(
	ListRepository repositoryinterface.ListRepositoryInterface,
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ImageRepository repositoryinterface.ImageRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
	ListMovieRepository repositoryinterface.ListMovieRepositoryInterface,
) *UpdateListUseCase {
	return &UpdateListUseCase{
		ListRepository:      ListRepository,
		ChooserRepository:   ChooserRepository,
		ImageRepository:     ImageRepository,
		MovieRepository:     MovieRepository,
		ListMovieRepository: ListMovieRepository,
	}
}

func (ul *UpdateListUseCase) Execute(input UpdateListInputDTO) (ListOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(ul.ChooserRepository, input.ChooserID, "UpdateListUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return ListOutputDTO{}, chooserValidatorProblems
	}

	list, listValidatorProblems := listValidator(ul.ListRepository, input.ID, "UpdateListUseCase")
	if len(listValidatorProblems.ProblemDetails) > 0 {
		return ListOutputDTO{}, listValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	doTheseMoviesExist, moviesForUpdate, manyMoviesError := ul.MovieRepository.DoTheseMoviesExist(input.Movies)
	if manyMoviesError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao resgatar os filmes pelos ids",
			Status:   http.StatusInternalServerError,
			Detail:   manyMoviesError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os filmes pelos ids", "CreateListUseCase", "Use Cases", util.TypeInternalServerError)

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doTheseMoviesExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Um ou mais filmes não encontrados",
			Status:   http.StatusConflict,
			Detail:   "Um ou mais ids dos filmes não retornou resultado",
			Instance: util.RFC409,
		})

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	validateListProblems := entity.ValidateList(input.Title, input.Description, input.ChooserID)
	if len(validateListProblems) > 0 {
		problemsDetails = append(problemsDetails, validateListProblems...)
	}

	if input.Title != list.Title {
		list.ChangeTitle(input.Title)
	}

	if input.Description != list.Description {
		list.ChangeDescription(input.Description)
	}

	if input.ProfileImageFile != nil {
		_, profileImageName, profileImageExtension, profileImageSize, profileImageError := service.MoveFile(input.ProfileImageFile, input.ProfileImageHandler)
		if profileImageError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao mover a imagem de profile da lista",
				Status:   http.StatusInternalServerError,
				Detail:   profileImageError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem de profile da lista", "UpdateListUseCase", "Use Cases", util.TypeInternalServerError)

			return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		newProfileImageName, newProfileImageNameProblems := entity.NewImage(profileImageName, profileImageExtension, profileImageSize)
		if len(newProfileImageNameProblems) > 0 {
			problemsDetails = append(problemsDetails, newProfileImageNameProblems...)
		}

		if len(problemsDetails) > 0 {
			return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		profileImageCreationError := ul.ImageRepository.Create(newProfileImageName)
		if profileImageCreationError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao persistir a imagem de profile",
				Status:   http.StatusInternalServerError,
				Detail:   profileImageCreationError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, profileImageCreationError.Error(), "UpdateListUseCase", "Use Cases", util.TypeInternalServerError)

			return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		list.ChangeProfileImageID(newProfileImageName.ID)
	}

	if input.CoverImageFile != nil {
		_, coverImageName, coverImageExtension, coverImageSize, coverImageError := service.MoveFile(input.CoverImageFile, input.CoverImageHandler)
		if coverImageError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao mover a imagem de capa da lista",
				Status:   http.StatusInternalServerError,
				Detail:   coverImageError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem de capa da lista", "UpdateListUseCase", "Use Cases", util.TypeInternalServerError)

			return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		newCoverImageName, newCoverImageNameProblems := entity.NewImage(coverImageName, coverImageExtension, coverImageSize)
		if len(newCoverImageNameProblems) > 0 {
			problemsDetails = append(problemsDetails, newCoverImageNameProblems...)
		}

		if len(problemsDetails) > 0 {
			return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		coverImageCreationError := ul.ImageRepository.Create(newCoverImageName)
		if coverImageCreationError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao persistir a imagem de capa",
				Status:   http.StatusInternalServerError,
				Detail:   coverImageCreationError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, coverImageCreationError.Error(), "UpdateListUseCase", "Use Cases", util.TypeInternalServerError)

			return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		list.ChangeCoverImageID(newCoverImageName.ID)
	}

	moviesToDelete, moviesToAdd := list.UpdateMovies(moviesForUpdate)

	list.RemoveMovies(moviesToDelete)
	list.AddMovies(moviesToAdd)

	listMoviesToDeactivateError := ul.ListMovieRepository.DeactivateAll(&moviesToDelete)
	if listMoviesToDeactivateError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao remover filmes da lista",
			Status:   http.StatusInternalServerError,
			Detail:   listMoviesToDeactivateError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listMoviesToDeactivateError.Error(), "UpdateListUseCase", "Use Cases", util.TypeInternalServerError)
	}

	var listMoviesToAdd []entity.ListMovie

	for _, movieToAdd := range moviesToAdd {
		newListMovie, newListMovieError := entity.NewListMovie(list.ID, movieToAdd.ID, input.ID)
		if newListMovieError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeValidationError,
				Title:    "Um ou mais filmes não encontrados",
				Status:   http.StatusBadRequest,
				Detail:   "Não foi possível adicioar o filme de ID " + movieToAdd.ID + " à lista de ID " + list.ID,
				Instance: util.RFC404,
			})

			return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		listMoviesToAdd = append(listMoviesToAdd, *newListMovie)
	}

	listMoviesToAddError := ul.ListMovieRepository.Create(&listMoviesToAdd)
	if listMoviesToAddError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao persistir filmes na lista",
			Status:   http.StatusInternalServerError,
			Detail:   listMoviesToAddError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listMoviesToAddError.Error(), "UpdateListUseCase", "Use Cases", util.TypeInternalServerError)
	}

	listUpdatedError := ul.ListRepository.Update(&list)
	if listUpdatedError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao persistir uma lista",
			Status:   http.StatusInternalServerError,
			Detail:   listUpdatedError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listUpdatedError.Error(), "UpdateListUseCase", "Use Cases", util.TypeInternalServerError)
	}

	output := NewListOutputDTO(list)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
