package usecase

import (
	"mime/multipart"
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/service"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"
)

type UpdateMovieInputDTO struct {
	ChooserID    string                `json:"chooser_id"`
	MovieID      string                `json:"movie_id"`
	Title        string                `json:"title"`
	CountryName  string                `json:"country_name"`
	Flag         string                `json:"flag"`
	ReleaseYear  int                   `json:"release_year"`
	ImageID      string                `json:"image_id"`
	ImageFile    multipart.File        `json:"movie_image_file"`
	ImageHandler *multipart.FileHeader `json:"movie_image_handler"`
	Genres       []GenreDTO            `json:"genres"`
	Directors    []DirectorDTO         `json:"directors"`
	Actors       []ActorDTO            `json:"actors"`
	Writers      []WriterDTO           `json:"writers"`
}

type UpdateMovieUseCase struct {
	ActorRepository         repositoryinterface.ActorRepositoryInterface
	ChooserRepository       repositoryinterface.ChooserRepositoryInterface
	DirectorRepository      repositoryinterface.DirectorRepositoryInterface
	GenreRepository         repositoryinterface.GenreRepositoryInterface
	ImageRepository         repositoryinterface.ImageRepositoryInterface
	MovieActorRepository    repositoryinterface.MovieActorRepositoryInterface
	MovieDirectorRepository repositoryinterface.MovieDirectorRepositoryInterface
	MovieGenreRepository    repositoryinterface.MovieGenreRepositoryInterface
	MovieRepository         repositoryinterface.MovieRepositoryInterface
	MovieWriterRepository   repositoryinterface.MovieWriterRepositoryInterface
	WriterRepository        repositoryinterface.WriterRepositoryInterface
}

func NewUpdateMovieUseCase(
	ActorRepository repositoryinterface.ActorRepositoryInterface,
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	DirectorRepository repositoryinterface.DirectorRepositoryInterface,
	GenreRepository repositoryinterface.GenreRepositoryInterface,
	ImageRepository repositoryinterface.ImageRepositoryInterface,
	MovieActorRepository repositoryinterface.MovieActorRepositoryInterface,
	MovieDirectorRepository repositoryinterface.MovieDirectorRepositoryInterface,
	MovieGenreRepository repositoryinterface.MovieGenreRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
	MovieWriterRepository repositoryinterface.MovieWriterRepositoryInterface,
	WriterRepository repositoryinterface.WriterRepositoryInterface,
) *UpdateMovieUseCase {
	return &UpdateMovieUseCase{
		ActorRepository:         ActorRepository,
		ChooserRepository:       ChooserRepository,
		DirectorRepository:      DirectorRepository,
		GenreRepository:         GenreRepository,
		ImageRepository:         ImageRepository,
		MovieActorRepository:    MovieActorRepository,
		MovieDirectorRepository: MovieDirectorRepository,
		MovieGenreRepository:    MovieGenreRepository,
		MovieRepository:         MovieRepository,
		MovieWriterRepository:   MovieWriterRepository,
		WriterRepository:        WriterRepository,
	}
}

func (up *UpdateMovieUseCase) Execute(input UpdateMovieInputDTO) (MovieOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(up.ChooserRepository, input.ChooserID, "UpdateMovieUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return MovieOutputDTO{}, chooserValidatorProblems
	}

	movie, movieValidatorProblems := movieValidator(up.MovieRepository, input.MovieID, "UpdateMovieUseCase")
	if len(movieValidatorProblems.ProblemDetails) > 0 {
		return MovieOutputDTO{}, movieValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	nationality, nationalityProblem := valueobject.NewNationality(input.CountryName, input.Flag)
	if len(nationalityProblem) > 0 {
		problemsDetails = append(problemsDetails, nationalityProblem...)
	}

	if len(input.Genres) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Gêneros não informados",
			Status:   http.StatusBadRequest,
			Detail:   "Nenhum gênero informado",
			Instance: util.RFC400,
		})
	}

	if len(input.Directors) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Diretores não informados",
			Status:   http.StatusBadRequest,
			Detail:   "Nenhum diretor informado",
			Instance: util.RFC400,
		})
	}

	if len(input.Actors) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Atores não informados",
			Status:   http.StatusBadRequest,
			Detail:   "Nenhum ator informado",
			Instance: util.RFC400,
		})
	}

	if len(input.Writers) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Escritores não informados",
			Status:   http.StatusBadRequest,
			Detail:   "Nenhum escritor informado",
			Instance: util.RFC400,
		})
	}

	if input.ImageID == "" && (input.ImageFile == nil || input.ImageHandler == nil) {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Imagem não informada",
			Status:   http.StatusBadRequest,
			Detail:   "O filme deve ter uma imagem",
			Instance: util.RFC400,
		})

	}

	if len(problemsDetails) > 0 {
		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	var imagesToAdd []entity.Image

	if input.ImageID == "" {
		_, newMovieImageProblemName, newMovieImageProblemExtension, newMovieImageProblemSize, newMovieImageProblemError := service.MoveFile(input.ImageFile, input.ImageHandler)
		if newMovieImageProblemError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao mover a imagem do filme",
				Status:   http.StatusInternalServerError,
				Detail:   newMovieImageProblemError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do filme", "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		newMovieImage, newMovieImageProblem := entity.NewImage(newMovieImageProblemName, newMovieImageProblemExtension, newMovieImageProblemSize)
		if len(newMovieImageProblem) > 0 {
			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: newMovieImageProblem,
			}
		}

		movie.ChangeImage(newMovieImage.ID)

		imagesToAdd = append(imagesToAdd, *newMovieImage)
	}

	validateMovieProblems := entity.ValidateMovie(input.Title, *nationality, input.ReleaseYear, movie.ImageID)
	if len(validateMovieProblems) > 0 {
		problemsDetails = append(problemsDetails, validateMovieProblems...)

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	if input.Title != movie.Title {
		movie.ChangeTitle(input.Title)
	}

	if !movie.Nationality.Equals(nationality) {
		movie.ChangeNationality(*nationality)
	}

	var genresToCheck []string
	var genresToAdd []GenreDTO

	for _, genre := range input.Genres {
		if genre.GenreID == "" {
			genresToAdd = append(genresToAdd, genre)
		} else {
			genresToCheck = append(genresToCheck, genre.GenreID)
		}
	}

	if len(movie.Genres) != len(genresToCheck) {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Gêneros não informados",
			Status:   http.StatusBadRequest,
			Detail:   "O filme deve ter todos os gêneros informados",
			Instance: util.RFC400,
		})

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	if len(genresToCheck) > 0 {
		doTheseGenresAreIncludedInTheMovie, _, doTheseGenresAreIncludedInTheMovieError := up.GenreRepository.DoTheseGenresAreIncludedInTheMovie(movie.ID, genresToCheck)
		if doTheseGenresAreIncludedInTheMovieError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao resgatar os gêneros do filme",
				Status:   http.StatusInternalServerError,
				Detail:   doTheseGenresAreIncludedInTheMovieError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os gêneros do filme", "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		} else if !doTheseGenresAreIncludedInTheMovie {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeBadRequest,
				Title:    "Gêneros não informados",
				Status:   http.StatusBadRequest,
				Detail:   "Algum dos gêneros informados não estão adicionados ao filme",
				Instance: util.RFC400,
			})

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}
	}

	if len(genresToAdd) > 0 {
		var newGenres []entity.Genre

		for _, genreToAdd := range genresToAdd {
			_, genreToAddName, genreToAddExtension, genreToAddSize, profileImageError := service.MoveFile(genreToAdd.ImageFile, genreToAdd.ImageHandler)
			if profileImageError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao mover a imagem do gênero: " + genreToAdd.Name,
					Status:   http.StatusInternalServerError,
					Detail:   profileImageError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do gênero: "+genreToAdd.Name, "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}
			}

			genreToAddImage, genreToAddImageProblem := entity.NewImage(genreToAddName, genreToAddExtension, genreToAddSize)
			if len(genreToAddImageProblem) > 0 {
				problemsDetails = append(problemsDetails, genreToAddImageProblem...)
			}

			newGenre, newGenreProblem := entity.NewGenre(genreToAdd.Name, genreToAddImage.ID)
			if len(newGenreProblem) > 0 {
				problemsDetails = append(problemsDetails, newGenreProblem...)
			}

			imagesToAdd = append(imagesToAdd, *genreToAddImage)
			newGenres = append(newGenres, *newGenre)
		}

		movie.AddGenres(newGenres)
	}

	var directorsToCheck []string
	var directorsToAdd []DirectorDTO

	for _, director := range input.Directors {
		if director.DirectorID == "" {
			directorsToAdd = append(directorsToAdd, director)
		} else {
			directorsToCheck = append(directorsToCheck, director.DirectorID)
		}
	}

	if len(movie.Directors) != len(directorsToCheck) {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Diretores não informados",
			Status:   http.StatusBadRequest,
			Detail:   "O filme deve ter todos(as) os(as) diretores(as) informados",
			Instance: util.RFC400,
		})

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	if len(directorsToCheck) > 0 {
		doTheseDirectorsAreIncludedInTheMovie, _, doTheseDirectorsAreIncludedInTheMovieError := up.DirectorRepository.DoTheseDirectorsAreIncludedInTheMovie(movie.ID, directorsToCheck)
		if doTheseDirectorsAreIncludedInTheMovieError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao resgatar os(as) diretores(as) do filme",
				Status:   http.StatusInternalServerError,
				Detail:   doTheseDirectorsAreIncludedInTheMovieError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os(as) diretores(as) do filme", "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		} else if !doTheseDirectorsAreIncludedInTheMovie {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeBadRequest,
				Title:    "Diretores não informados",
				Status:   http.StatusBadRequest,
				Detail:   "Alguns(mas) dos(as) diretores(as) informados(as) não estão adicionados(as) ao filme",
				Instance: util.RFC400,
			})
		}
	}

	if len(directorsToAdd) > 0 {
		var newDirectors []entity.Director

		for _, directorToAdd := range directorsToAdd {
			_, directorToAddName, directorToAddExtension, directorToAddSize, profileImageError := service.MoveFile(directorToAdd.ImageFile, directorToAdd.ImageHandler)
			if profileImageError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao mover a imagem do(a) diretor(a): " + directorToAdd.Name,
					Status:   http.StatusInternalServerError,
					Detail:   profileImageError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do(a) diretor(a): "+directorToAdd.Name, "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}
			}

			newDirectorBirthday, newDirectorBirthdayProblem := valueobject.NewBirthDate(directorToAdd.Day, directorToAdd.Month, directorToAdd.Year)
			if len(newDirectorBirthdayProblem) > 0 {
				problemsDetails = append(problemsDetails, newDirectorBirthdayProblem...)
			}

			newDirectorNationality, newDirectorNationalityProblem := valueobject.NewNationality(directorToAdd.CountryName, directorToAdd.Flag)
			if len(newDirectorNationalityProblem) > 0 {
				problemsDetails = append(problemsDetails, newDirectorNationalityProblem...)
			}

			directorToAddImage, directorToAddImageProblem := entity.NewImage(directorToAddName, directorToAddExtension, directorToAddSize)
			if len(directorToAddImageProblem) > 0 {
				problemsDetails = append(problemsDetails, directorToAddImageProblem...)
			}

			newDirector, newDirectorProblem := entity.NewDirector(directorToAdd.Name, newDirectorBirthday, newDirectorNationality, directorToAddImage.ID)
			if len(newDirectorProblem) > 0 {
				problemsDetails = append(problemsDetails, newDirectorProblem...)
			}

			imagesToAdd = append(imagesToAdd, *directorToAddImage)
			newDirectors = append(newDirectors, *newDirector)
		}

		movie.AddDirectors(newDirectors)
	}

	var actorsToCheck []string
	var actorsToAdd []ActorDTO

	for _, actor := range input.Actors {
		if actor.ActorID == "" {
			actorsToAdd = append(actorsToAdd, actor)
		} else {
			actorsToCheck = append(actorsToCheck, actor.ActorID)
		}
	}

	if len(movie.Actors) != len(actorsToCheck) {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Atores não informados",
			Status:   http.StatusBadRequest,
			Detail:   "O filme deve ter todos(as) os(as) atores(atrizes) informados(as)",
			Instance: util.RFC400,
		})

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	if len(actorsToCheck) > 0 {
		doTheseActorsAreIncludedInTheMovie, _, doTheseActorsAreIncludedInTheMovieError := up.ActorRepository.DoTheseActorsAreIncludedInTheMovie(movie.ID, actorsToCheck)
		if doTheseActorsAreIncludedInTheMovieError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao resgatar os(as) atores(atrizes) do filme",
				Status:   http.StatusInternalServerError,
				Detail:   doTheseActorsAreIncludedInTheMovieError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os(as) atores(atrizes) do filme", "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		} else if !doTheseActorsAreIncludedInTheMovie {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeBadRequest,
				Title:    "Atores não informados",
				Status:   http.StatusBadRequest,
				Detail:   "Algum dos(as) atores(atrizes) informados(as) não estão adicionados(as) ao filme",
				Instance: util.RFC400,
			})
		}
	}

	if len(actorsToAdd) > 0 {
		var newActors []entity.Actor

		for _, actorToAdd := range actorsToAdd {
			_, actorToAddName, actorToAddExtension, actorToAddSize, profileImageError := service.MoveFile(actorToAdd.ImageFile, actorToAdd.ImageHandler)
			if profileImageError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao mover a imagem do(a) ator(atriz): " + actorToAdd.Name,
					Status:   http.StatusInternalServerError,
					Detail:   profileImageError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do(a) ator(atriz): "+actorToAdd.Name, "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}
			}

			newActorBirthday, newActorBirthdayProblem := valueobject.NewBirthDate(actorToAdd.Day, actorToAdd.Month, actorToAdd.Year)
			if len(newActorBirthdayProblem) > 0 {
				problemsDetails = append(problemsDetails, newActorBirthdayProblem...)
			}

			newActorNationality, newActorNationalityProblem := valueobject.NewNationality(actorToAdd.CountryName, actorToAdd.Flag)
			if len(newActorNationalityProblem) > 0 {
				problemsDetails = append(problemsDetails, newActorNationalityProblem...)
			}

			actorToAddImage, actorToAddImageProblem := entity.NewImage(actorToAddName, actorToAddExtension, actorToAddSize)
			if len(actorToAddImageProblem) > 0 {
				problemsDetails = append(problemsDetails, actorToAddImageProblem...)
			}

			newActor, newActorProblem := entity.NewActor(actorToAdd.Name, newActorBirthday, newActorNationality, actorToAddImage.ID)
			if len(newActorProblem) > 0 {
				problemsDetails = append(problemsDetails, newActorProblem...)
			}

			imagesToAdd = append(imagesToAdd, *actorToAddImage)
			newActors = append(newActors, *newActor)
		}

		movie.AddActors(newActors)
	}

	var writersToCheck []string
	var writersToAdd []WriterDTO

	for _, writer := range input.Writers {
		if writer.WriterID == "" {
			writersToAdd = append(writersToAdd, writer)
		} else {
			writersToCheck = append(writersToCheck, writer.WriterID)
		}
	}

	if len(movie.Writers) != len(writersToCheck) {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeBadRequest,
			Title:    "Escritores não informados",
			Status:   http.StatusBadRequest,
			Detail:   "O filme deve ter todos(as) os(as) escritores(escritoras) informados(as)",
			Instance: util.RFC400,
		})

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	if len(writersToCheck) > 0 {
		doTheseWritersAreIncludedInTheMovie, _, doTheseWritersAreIncludedInTheMovieError := up.WriterRepository.DoTheseWritersAreIncludedInTheMovie(movie.ID, writersToCheck)
		if doTheseWritersAreIncludedInTheMovieError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao resgatar os(as) escritores(escritoras) do filme",
				Status:   http.StatusInternalServerError,
				Detail:   doTheseWritersAreIncludedInTheMovieError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os(as) escritores(escritoras) do filme", "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		} else if !doTheseWritersAreIncludedInTheMovie {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeBadRequest,
				Title:    "Escritores não informados",
				Status:   http.StatusBadRequest,
				Detail:   "Algum dos(as) escritores(escritoras) informados(as) não estão adicionados(as) ao filme",
				Instance: util.RFC400,
			})
		}
	}

	if len(writersToAdd) > 0 {
		var newWriters []entity.Writer

		for _, writerToAdd := range writersToAdd {
			_, writerToAddName, writerToAddExtension, writerToAddSize, profileImageError := service.MoveFile(writerToAdd.ImageFile, writerToAdd.ImageHandler)
			if profileImageError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao mover a imagem do(a) escritor(escritora): " + writerToAdd.Name,
					Status:   http.StatusInternalServerError,
					Detail:   profileImageError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do(a) escritor(escritora): "+writerToAdd.Name, "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}
			}

			newWriterBirthday, newWriterBirthdayProblem := valueobject.NewBirthDate(writerToAdd.Day, writerToAdd.Month, writerToAdd.Year)
			if len(newWriterBirthdayProblem) > 0 {
				problemsDetails = append(problemsDetails, newWriterBirthdayProblem...)
			}

			newWriterNationality, newWriterNationalityProblem := valueobject.NewNationality(writerToAdd.CountryName, writerToAdd.Flag)
			if len(newWriterNationalityProblem) > 0 {
				problemsDetails = append(problemsDetails, newWriterNationalityProblem...)
			}

			writerToAddImage, writerToAddImageProblem := entity.NewImage(writerToAddName, writerToAddExtension, writerToAddSize)
			if len(writerToAddImageProblem) > 0 {
				problemsDetails = append(problemsDetails, writerToAddImageProblem...)
			}

			newWriter, newWriterProblem := entity.NewWriter(writerToAdd.Name, newWriterBirthday, newWriterNationality, writerToAddImage.ID)
			if len(newWriterProblem) > 0 {
				problemsDetails = append(problemsDetails, newWriterProblem...)
			}

			imagesToAdd = append(imagesToAdd, *writerToAddImage)
			newWriters = append(newWriters, *newWriter)
		}

		movie.AddWriters(newWriters)
	}

	if len(problemsDetails) > 0 {
		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	if len(imagesToAdd) > 0 {
		imagesToAddError := up.ImageRepository.CreateMany(&imagesToAdd)
		if imagesToAddError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao criar imagens",
				Status:   http.StatusInternalServerError,
				Detail:   imagesToAddError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar as imagens do filme", "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}
	}

	updatedMovieError := up.MovieRepository.Update(&movie)
	if updatedMovieError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao atualizar o filme",
			Status:   http.StatusInternalServerError,
			Detail:   updatedMovieError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao atualizar o filme", "UpdateMovieUseCase", "Use Cases", util.TypeInternalServerError)

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := NewMovieOutputDTO(movie)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
