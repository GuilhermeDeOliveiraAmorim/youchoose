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

type UpdateListOutputDTO struct {
	ID             string         `json:"id"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	ProfileImageID string         `json:"profile_image_id"`
	CoverImageID   string         `json:"cover_image_id"`
	ChooserID      string         `json:"chooser_id"`
	Movies         []entity.Movie `json:"movies"`
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

func (ul *UpdateListUseCase) Execute(input UpdateListInputDTO) (UpdateListOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, chooser, getChooserError := ul.ChooserRepository.GetByID(input.ChooserID)
	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar chooser de ID " + input.ID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), "UpdateListUseCase", "Use Cases", "Internal Server Error")

		return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheListExist, listThatExists, getListError := ul.ListRepository.GetByID(input.ID)
	if getListError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar lista de ID " + input.ID,
			Status:   http.StatusInternalServerError,
			Detail:   getListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getListError.Error(), "UpdateListUseCase", "Use Cases", "Internal Server Error")

		return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !listThatExists.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Lista não encontrada",
			Status:   http.StatusNotFound,
			Detail:   "A lista com o ID " + input.ID + " está desativada",
			Instance: util.RFC404,
		})

		return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doTheseMoviesExist, moviesForUpdate, manyMoviesError := ul.MovieRepository.DoTheseMoviesExist(input.Movies)
	if manyMoviesError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar os filmes pelos ids",
			Status:   http.StatusInternalServerError,
			Detail:   manyMoviesError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os filmes pelos ids", "CreateListUseCase", "Use Cases", "Internal Server Error")

		return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doTheseMoviesExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Um ou mais filmes não encontrados",
			Status:   http.StatusConflict,
			Detail:   "Um ou mais ids dos filmes não retornou resultado",
			Instance: util.RFC409,
		})

		return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	validateListProblems := entity.ValidateList(input.Title, input.Description, input.ChooserID)
	if len(validateListProblems) > 0 {
		problemsDetails = append(problemsDetails, validateListProblems...)
	}

	if input.Title != listThatExists.Title {
		listThatExists.ChangeTitle(input.Title)
	}

	if input.Description != listThatExists.Description {
		listThatExists.ChangeDescription(input.Description)
	}

	if input.ProfileImageFile != nil {
		_, profileImageName, profileImageExtension, profileImageSize, profileImageError := service.MoveFile(input.ProfileImageFile, input.ProfileImageHandler)
		if profileImageError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Internal Server Error",
				Title:    "Erro ao mover a imagem de profile da lista",
				Status:   http.StatusInternalServerError,
				Detail:   profileImageError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem de profile da lista", "UpdateListUseCase", "Use Cases", "Internal Server Error")

			return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		newProfileImageName, newProfileImageNameProblems := entity.NewImage(profileImageName, profileImageExtension, profileImageSize)
		if len(newProfileImageNameProblems) > 0 {
			problemsDetails = append(problemsDetails, newProfileImageNameProblems...)
		}

		if len(problemsDetails) > 0 {
			return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		profileImageCreationError := ul.ImageRepository.Create(newProfileImageName)
		if profileImageCreationError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Internal Server Error",
				Title:    "Erro ao persistir a imagem de profile",
				Status:   http.StatusInternalServerError,
				Detail:   profileImageCreationError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, profileImageCreationError.Error(), "UpdateListUseCase", "Use Cases", "Internal Server Error")

			return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		listThatExists.ChangeProfileImageID(newProfileImageName.ID)
	}

	if input.CoverImageFile != nil {
		_, coverImageName, coverImageExtension, coverImageSize, coverImageError := service.MoveFile(input.CoverImageFile, input.CoverImageHandler)
		if coverImageError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Internal Server Error",
				Title:    "Erro ao mover a imagem de capa da lista",
				Status:   http.StatusInternalServerError,
				Detail:   coverImageError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem de capa da lista", "UpdateListUseCase", "Use Cases", "Internal Server Error")

			return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		newCoverImageName, newCoverImageNameProblems := entity.NewImage(coverImageName, coverImageExtension, coverImageSize)
		if len(newCoverImageNameProblems) > 0 {
			problemsDetails = append(problemsDetails, newCoverImageNameProblems...)
		}

		if len(problemsDetails) > 0 {
			return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		coverImageCreationError := ul.ImageRepository.Create(newCoverImageName)
		if coverImageCreationError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Internal Server Error",
				Title:    "Erro ao persistir a imagem de capa",
				Status:   http.StatusInternalServerError,
				Detail:   coverImageCreationError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, coverImageCreationError.Error(), "UpdateListUseCase", "Use Cases", "Internal Server Error")

			return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		listThatExists.ChangeCoverImageID(newCoverImageName.ID)
	}

	moviesToDelete, moviesToAdd := listThatExists.UpdateMovies(moviesForUpdate)

	listThatExists.RemoveMovies(moviesToDelete)
	listThatExists.AddMovies(moviesToAdd)

	listMoviesToDeactivateError := ul.ListMovieRepository.DeactivateAll(&moviesToDelete)
	if listMoviesToDeactivateError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao remover filmes da lista",
			Status:   http.StatusInternalServerError,
			Detail:   listMoviesToDeactivateError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listMoviesToDeactivateError.Error(), "UpdateListUseCase", "Use Cases", "Internal Server Error")
	}

	var listMoviesToAdd []entity.ListMovie

	for _, movieToAdd := range moviesToAdd {
		newListMovie, newListMovieError := entity.NewListMovie(listThatExists.ID, movieToAdd.ID)
		if newListMovieError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Validation Error",
				Title:    "Um ou mais filmes não encontrados",
				Status:   http.StatusBadRequest,
				Detail:   "Não foi possível adicioar o filme de ID " + movieToAdd.ID + " à lista de ID " + listThatExists.ID,
				Instance: util.RFC404,
			})

			return UpdateListOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		listMoviesToAdd = append(listMoviesToAdd, *newListMovie)
	}

	listMoviesToAddError := ul.ListMovieRepository.Create(&listMoviesToAdd)
	if listMoviesToAddError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir filmes na lista",
			Status:   http.StatusInternalServerError,
			Detail:   listMoviesToAddError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listMoviesToAddError.Error(), "UpdateListUseCase", "Use Cases", "Internal Server Error")
	}

	listUpdatedError := ul.ListRepository.Update(&listThatExists)
	if listUpdatedError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir uma lista",
			Status:   http.StatusInternalServerError,
			Detail:   listUpdatedError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listUpdatedError.Error(), "UpdateListUseCase", "Use Cases", "Internal Server Error")
	}

	output := UpdateListOutputDTO{
		ID:             listThatExists.ID,
		Title:          listThatExists.Title,
		Description:    listThatExists.Description,
		ChooserID:      listThatExists.ChooserID,
		ProfileImageID: listThatExists.ProfileImageID,
		CoverImageID:   listThatExists.CoverImageID,
		Movies:         listThatExists.Movies,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
