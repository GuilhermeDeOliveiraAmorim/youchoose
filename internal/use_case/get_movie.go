package usecase

import (
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"
)

type GetMovieInputDTO struct {
	ChooserID string `json:"chooser_id"`
	MovieID   string `json:"movie_id"`
}

type GetMovieOutputDTO struct {
	ID          string                  `json:"id"`
	Title       string                  `json:"title"`
	Nationality valueobject.Nationality `json:"nationality"`
	ReleaseYear int                     `json:"release_year"`
	ImageID     string                  `json:"image_id"`
	Votes       int                     `json:"votes"`
	Genres      []entity.Genre          `json:"genres"`
	Directors   []entity.Director       `json:"directors"`
	Actors      []entity.Actor          `json:"actors"`
	Writers     []entity.Writer         `json:"writers"`
}

type GetMovieUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
	MovieRepository   repositoryinterface.MovieRepositoryInterface
}

func NewGetMovieUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
) *GetMovieUseCase {
	return &GetMovieUseCase{
		ChooserRepository: ChooserRepository,
		MovieRepository:   MovieRepository,
	}
}

func (gm *GetMovieUseCase) Execute(input GetMovieInputDTO) (GetMovieOutputDTO, util.ProblemDetailsOutputDTO) {
	problemsDetails := []util.ProblemDetails{}

	doesTheChooserExist, chooser, getChooserError := gm.ChooserRepository.GetByID(input.ChooserID)
	if getChooserError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar chooser de ID " + input.ChooserID,
			Status:   http.StatusInternalServerError,
			Detail:   getChooserError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getChooserError.Error(), "GetMovieUseCase", "Use Cases", "Internal Server Error")

		return GetMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheChooserExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum chooser com o ID " + input.ChooserID + " foi encontrado",
			Instance: util.RFC404,
		})

		return GetMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !chooser.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Chooser não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "O chooser com o ID " + input.ChooserID + " está desativado",
			Instance: util.RFC404,
		})

		return GetMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	doesTheMovieExist, movie, getMovieError := gm.MovieRepository.GetByID(input.MovieID)
	if getMovieError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao resgatar filme de ID " + input.MovieID,
			Status:   http.StatusInternalServerError,
			Detail:   getMovieError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, getMovieError.Error(), "GetMovieUseCase", "Use Cases", "Internal Server Error")

		return GetMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !doesTheMovieExist {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Filme não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "Nenhum filme com o ID " + input.MovieID + " foi encontrado",
			Instance: util.RFC404,
		})

		return GetMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	} else if !movie.Active {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Not Found",
			Title:    "Filme não encontrado",
			Status:   http.StatusNotFound,
			Detail:   "O filme com o ID " + input.MovieID + " está desativado",
			Instance: util.RFC404,
		})

		return GetMovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := GetMovieOutputDTO{
		ID:          movie.ID,
		Title:       movie.Title,
		Nationality: movie.Nationality,
		ReleaseYear: movie.ReleaseYear,
		ImageID:     movie.ImageID,
		Votes:       movie.Votes,
		Genres:      movie.Genres,
		Directors:   movie.Directors,
		Actors:      movie.Actors,
		Writers:     movie.Writers,
	}

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
