package usecase

import (
	"mime/multipart"
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/service"
	"youchoose/internal/util"
)

type CreateListInputDTO struct {
	Title               string                `json:"title"`
	ProfileImageFile    multipart.File        `json:"profile_image_file"`
	ProfileImageHandler *multipart.FileHeader `json:"profile_image_handler"`
	CoverImageFile      multipart.File        `json:"cover_image_file"`
	CoverImageHandler   *multipart.FileHeader `json:"cover_image_handler"`
	Description         string                `json:"description"`
	Movies              []string              `json:"movies"`
	ChooserID           string                `json:"chooser_id"`
}

type CreateListUseCase struct {
	ListRepository      repositoryinterface.ListRepositoryInterface
	ChooserRepository   repositoryinterface.ChooserRepositoryInterface
	ImageRepository     repositoryinterface.ImageRepositoryInterface
	MovieRepository     repositoryinterface.MovieRepositoryInterface
	ListMovieRepository repositoryinterface.ListMovieRepositoryInterface
}

func NewCreateListUseCase(
	ListRepository repositoryinterface.ListRepositoryInterface,
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ImageRepository repositoryinterface.ImageRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
	ListMovieRepository repositoryinterface.ListMovieRepositoryInterface,
) *CreateListUseCase {
	return &CreateListUseCase{
		ListRepository:      ListRepository,
		ChooserRepository:   ChooserRepository,
		ImageRepository:     ImageRepository,
		MovieRepository:     MovieRepository,
		ListMovieRepository: ListMovieRepository,
	}
}

func (cl *CreateListUseCase) Execute(input CreateListInputDTO) (ListOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(cl.ChooserRepository, input.ChooserID, "CreateListUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return ListOutputDTO{}, chooserValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	doTheseMoviesExist, _, manyMoviesError := cl.MovieRepository.DoTheseMoviesExist(input.Movies)
	if manyMoviesError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar os filmes pelos ids",
			Status:   http.StatusInternalServerError,
			Detail:   manyMoviesError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os filmes pelos ids", "CreateListUseCase", "Use Cases", "Internal Server Error")

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
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

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	_, profileImageName, profileImageExtension, profileImageSize, profileImageError := service.MoveFile(input.ProfileImageFile, input.ProfileImageHandler)
	if profileImageError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao mover a imagem de profile da lista",
			Status:   http.StatusInternalServerError,
			Detail:   profileImageError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem de profile da lista", "CreateListUseCase", "Use Cases", "Internal Server Error")

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	_, coverImageName, coverImageExtension, coverImageSize, coverImageError := service.MoveFile(input.CoverImageFile, input.CoverImageHandler)
	if coverImageError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao mover a imagem de capa da lista",
			Status:   http.StatusInternalServerError,
			Detail:   coverImageError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem de capa da lista", "CreateListUseCase", "Use Cases", "Internal Server Error")

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	newProfileImageName, newProfileImageNameProblems := entity.NewImage(profileImageName, profileImageExtension, profileImageSize)
	if len(newProfileImageNameProblems) > 0 {
		problemsDetails = append(problemsDetails, newProfileImageNameProblems...)
	}

	newCoverImageName, newCoverImageNameProblems := entity.NewImage(coverImageName, coverImageExtension, coverImageSize)
	if len(newCoverImageNameProblems) > 0 {
		problemsDetails = append(problemsDetails, newCoverImageNameProblems...)
	}

	newList, newListProblems := entity.NewList(input.Title, input.Description, newProfileImageName.ID, newCoverImageName.ID, input.ChooserID)
	if len(newListProblems) > 0 {
		problemsDetails = append(problemsDetails, newListProblems...)
	}

	var listMovies []entity.ListMovie

	for _, movieID := range input.Movies {
		newListMovie, newListMovieProblems := entity.NewListMovie(newList.ID, movieID, input.ChooserID)
		if len(newListMovieProblems) > 0 {
			problemsDetails = append(problemsDetails, newListMovieProblems...)
		}

		listMovies = append(listMovies, *newListMovie)
	}

	if len(problemsDetails) > 0 {
		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	profileImageCreationError := cl.ImageRepository.Create(newProfileImageName)
	if profileImageCreationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir a imagem de profile",
			Status:   http.StatusInternalServerError,
			Detail:   profileImageCreationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, profileImageCreationError.Error(), "CreateListUseCase", "Use Cases", "Internal Server Error")

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	coverImageCreationError := cl.ImageRepository.Create(newCoverImageName)
	if coverImageCreationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir a imagem de capa",
			Status:   http.StatusInternalServerError,
			Detail:   coverImageCreationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, coverImageCreationError.Error(), "CreateListUseCase", "Use Cases", "Internal Server Error")

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	listCreationError := cl.ListRepository.Create(newList)
	if listCreationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir uma lista",
			Status:   http.StatusInternalServerError,
			Detail:   listCreationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listCreationError.Error(), "CreateListUseCase", "Use Cases", "Internal Server Error")

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	listMoviesCreationError := cl.ListMovieRepository.Create(&listMovies)
	if listMoviesCreationError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir relações entre lista e filmes",
			Status:   http.StatusInternalServerError,
			Detail:   listMoviesCreationError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, listMoviesCreationError.Error(), "CreateListUseCase", "Use Cases", "Internal Server Error")

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	addedMoviesList, addedMoviesListError := cl.ListRepository.GetAllMoviesByListID(newList.ID)
	if addedMoviesListError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao retornar os filmes da lista de id " + newList.ID,
			Status:   http.StatusInternalServerError,
			Detail:   addedMoviesListError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, addedMoviesListError.Error(), "CreateListUseCase", "Use Cases", "Internal Server Error")

		return ListOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	newList.AddMovies(addedMoviesList)

	output := NewListOutputDTO(*newList)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
