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

type GenreDTO struct {
	GenreID      string                `json:"genre_id"`
	Name         string                `json:"name"`
	ImageFile    multipart.File        `json:"genre_image_file"`
	ImageHandler *multipart.FileHeader `json:"genre_image_handler"`
}

type DirectorDTO struct {
	DirectorID   string                `json:"director_id"`
	Name         string                `json:"name"`
	Day          int                   `json:"day"`
	Month        int                   `json:"month"`
	Year         int                   `json:"year"`
	CountryName  string                `json:"country_name"`
	Flag         string                `json:"flag"`
	ImageFile    multipart.File        `json:"director_image_file"`
	ImageHandler *multipart.FileHeader `json:"director_image_handler"`
}

type ActorDTO struct {
	ActorID      string                `json:"actor_id"`
	Name         string                `json:"name"`
	Day          int                   `json:"day"`
	Month        int                   `json:"month"`
	Year         int                   `json:"year"`
	CountryName  string                `json:"country_name"`
	Flag         string                `json:"flag"`
	ImageFile    multipart.File        `json:"actor_image_file"`
	ImageHandler *multipart.FileHeader `json:"actor_image_handler"`
}

type WriterDTO struct {
	WriterID     string                `json:"writer_id"`
	Name         string                `json:"name"`
	Day          int                   `json:"day"`
	Month        int                   `json:"month"`
	Year         int                   `json:"year"`
	CountryName  string                `json:"country_name"`
	Flag         string                `json:"flag"`
	ImageFile    multipart.File        `json:"writer_image_file"`
	ImageHandler *multipart.FileHeader `json:"writer_image_handler"`
}

type CreateMovieInputDTO struct {
	ChooserID    string                `json:"chooser_id"`
	Title        string                `json:"title"`
	CountryName  string                `json:"country_name"`
	Flag         string                `json:"flag"`
	ReleaseYear  int                   `json:"release_year"`
	ImageFile    multipart.File        `json:"movie_image_file"`
	ImageHandler *multipart.FileHeader `json:"movie_image_handler"`
	Genres       []GenreDTO            `json:"genres"`
	Directors    []DirectorDTO         `json:"directors"`
	Actors       []ActorDTO            `json:"actors"`
	Writers      []WriterDTO           `json:"writers"`
}

type CreateMovieUseCase struct {
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

func NewCreateMovieUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	MovieRepository repositoryinterface.MovieRepositoryInterface,
	ImageRepository repositoryinterface.ImageRepositoryInterface,
	GenreRepository repositoryinterface.GenreRepositoryInterface,
	DirectorRepository repositoryinterface.DirectorRepositoryInterface,
	ActorRepository repositoryinterface.ActorRepositoryInterface,
	WriterRepository repositoryinterface.WriterRepositoryInterface,
	MovieGenreRepository repositoryinterface.MovieGenreRepositoryInterface,
	MovieDirectorRepository repositoryinterface.MovieDirectorRepositoryInterface,
	MovieActorRepository repositoryinterface.MovieActorRepositoryInterface,
	MovieWriterRepository repositoryinterface.MovieWriterRepositoryInterface,
) *CreateMovieUseCase {
	return &CreateMovieUseCase{
		ChooserRepository:       ChooserRepository,
		MovieRepository:         MovieRepository,
		ImageRepository:         ImageRepository,
		GenreRepository:         GenreRepository,
		DirectorRepository:      DirectorRepository,
		ActorRepository:         ActorRepository,
		WriterRepository:        WriterRepository,
		MovieGenreRepository:    MovieGenreRepository,
		MovieDirectorRepository: MovieDirectorRepository,
		MovieActorRepository:    MovieActorRepository,
		MovieWriterRepository:   MovieWriterRepository,
	}
}

func (cm *CreateMovieUseCase) Execute(input CreateMovieInputDTO) (MovieOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(cm.ChooserRepository, input.ChooserID, "CreateMovieUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return MovieOutputDTO{}, chooserValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

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

	nationality, nationalityProblem := valueobject.NewNationality(input.CountryName, input.Flag)
	if len(nationalityProblem) > 0 {
		problemsDetails = append(problemsDetails, nationalityProblem...)
	}

	if len(problemsDetails) > 0 {
		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	_, movieImageName, movieImageExtension, movieImageSize, profileImageError := service.MoveFile(input.ImageFile, input.ImageHandler)
	if profileImageError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao mover a imagem do filme",
			Status:   http.StatusInternalServerError,
			Detail:   profileImageError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do filme", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	movieImage, movieImageProblem := entity.NewImage(movieImageName, movieImageExtension, movieImageSize)
	if len(movieImageProblem) > 0 {
		problemsDetails = append(problemsDetails, movieImageProblem...)
	}

	newMovie, movieProblem := entity.NewMovie(input.Title, *nationality, input.ReleaseYear, movieImage.ID)
	if len(movieProblem) > 0 {
		problemsDetails = append(problemsDetails, movieProblem...)
	}

	if len(problemsDetails) > 0 {
		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	movieImageToAddError := cm.ImageRepository.Create(movieImage)
	if movieImageToAddError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao criar imagem",
			Status:   http.StatusInternalServerError,
			Detail:   movieImageToAddError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar a imagem do filme", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	newMovieError := cm.MovieRepository.Create(newMovie)
	if newMovieError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao criar o filme",
			Status:   http.StatusInternalServerError,
			Detail:   newMovieError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar o filme", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	_, newMovie = TreatmentGenres(input.Genres, cm, newMovie)
	_, newMovie = TreatmentActors(input.Actors, cm, newMovie)
	_, newMovie = TreatmentDirectors(input.Directors, cm, newMovie)
	_, newMovie = TreatmentWriters(input.Writers, cm, newMovie)

	output := NewMovieOutputDTO(*newMovie)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}

func TreatmentGenres(inputGenres []GenreDTO, cm *CreateMovieUseCase, newMovie *entity.Movie) (util.ProblemDetailsOutputDTO, *entity.Movie) {
	var genresToCheck []string
	var genresToAdd []GenreDTO
	var problemsDetails []util.ProblemDetails
	var imagesToAdd []entity.Image

	for _, genre := range inputGenres {
		if genre.GenreID == "" {
			genresToAdd = append(genresToAdd, genre)
		} else {
			genresToCheck = append(genresToCheck, genre.GenreID)
		}
	}

	if len(genresToCheck) > 0 {
		doTheseGenresExist, existingGenres, manyGenresError := cm.GenreRepository.DoTheseGenresExist(genresToCheck)
		if manyGenresError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao resgatar os gêneros pelos ids",
				Status:   http.StatusInternalServerError,
				Detail:   manyGenresError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os gêneros pelos ids", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		} else if !doTheseGenresExist {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeValidationError,
				Title:    "Um ou mais gêneros não encontrados",
				Status:   http.StatusConflict,
				Detail:   "Um ou mais ids dos gêneros não retornou resultado",
				Instance: util.RFC409,
			})

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		}

		if len(existingGenres) > 0 {
			newMovie.AddGenres(existingGenres)
		}
	}

	var newGenres []entity.Genre

	if len(genresToAdd) > 0 {
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

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do gênero: "+genreToAdd.Name, "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}

			genreToAddImage, genreToAddImageProblem := entity.NewImage(genreToAddName, genreToAddExtension, genreToAddSize)
			if len(genreToAddImageProblem) > 0 {
				problemsDetails = append(problemsDetails, genreToAddImageProblem...)
			}

			newGenre, newGenreProblem := entity.NewGenre(genreToAdd.Name, genreToAddImage.ID)
			if len(newGenreProblem) > 0 {
				problemsDetails = append(problemsDetails, newGenreProblem...)
			}

			if len(problemsDetails) > 0 {
				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}

			imagesToAdd = append(imagesToAdd, *genreToAddImage)
			newGenres = append(newGenres, *newGenre)
		}

		if len(imagesToAdd) > 0 {
			imagesToAddError := cm.ImageRepository.CreateMany(&imagesToAdd)
			if imagesToAddError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao criar imagens",
					Status:   http.StatusInternalServerError,
					Detail:   imagesToAddError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar as imagens dos gêneros", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}
		}

		if len(newGenres) > 0 {
			genresToAddError := cm.GenreRepository.CreateMany(&newGenres)
			if genresToAddError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao criar os gêneros",
					Status:   http.StatusInternalServerError,
					Detail:   genresToAddError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar os gêneros", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}

			newMovie.AddGenres(newGenres)
		}
	}

	var movieGenres []entity.MovieGenre

	for _, genre := range newMovie.Genres {
		newMovieGenre, newMovieGenreError := entity.NewMovieGenre(newMovie.ID, genre.ID)
		if newMovieGenreError != nil {
			problemsDetails = append(problemsDetails, newMovieGenreError...)
		}

		movieGenres = append(movieGenres, *newMovieGenre)
	}

	movieGenresError := cm.MovieGenreRepository.CreateMany(&movieGenres)
	if movieGenresError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     util.TypeInternalServerError,
			Title:    "Erro ao criar os gêneros",
			Status:   http.StatusInternalServerError,
			Detail:   movieGenresError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar os gêneros", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

		return util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}, nil
	}

	return util.ProblemDetailsOutputDTO{}, newMovie
}

func TreatmentActors(inputActors []ActorDTO, cm *CreateMovieUseCase, newMovie *entity.Movie) (util.ProblemDetailsOutputDTO, *entity.Movie) {
	var actorsToCheck []string
	var actorsToAdd []ActorDTO
	var problemsDetails []util.ProblemDetails
	var imagesToAdd []entity.Image

	for _, actor := range inputActors {
		if actor.ActorID == "" {
			actorsToAdd = append(actorsToAdd, actor)
		} else {
			actorsToCheck = append(actorsToCheck, actor.ActorID)
		}
	}

	if len(actorsToCheck) > 0 {
		doTheseActorsExist, existingActors, manyActorsError := cm.ActorRepository.DoTheseActorsExist(actorsToCheck)
		if manyActorsError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao resgatar o(a)s atores(atrizes) pelos ids",
				Status:   http.StatusInternalServerError,
				Detail:   manyActorsError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar do(a)s atores(atrizes) pelos ids", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		} else if !doTheseActorsExist {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeValidationError,
				Title:    "Um ou mais gêneros não encontrados",
				Status:   http.StatusConflict,
				Detail:   "Um ou mais ids do(a)s atores(atrizes) não retornou resultado",
				Instance: util.RFC409,
			})

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		}

		if len(existingActors) > 0 {
			newMovie.AddActors(existingActors)
		}
	}

	var newActors []entity.Actor

	if len(actorsToAdd) > 0 {
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

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do(a) ator(atriz): "+actorToAdd.Name, "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
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

			if len(problemsDetails) > 0 {
				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}

			imagesToAdd = append(imagesToAdd, *actorToAddImage)
			newActors = append(newActors, *newActor)
		}

		if len(imagesToAdd) > 0 {
			imagesToAddError := cm.ImageRepository.CreateMany(&imagesToAdd)
			if imagesToAddError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao criar imagens",
					Status:   http.StatusInternalServerError,
					Detail:   imagesToAddError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar as imagens do(a)s atores(atrizes)", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}
		}

		if len(newActors) > 0 {
			actorsToAddError := cm.ActorRepository.CreateMany(&newActors)
			if actorsToAddError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao criar o(a)s atores(atrizes)",
					Status:   http.StatusInternalServerError,
					Detail:   actorsToAddError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar o(a)s atores(atrizes)", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}

			newMovie.AddActors(newActors)
		}
	}

	var movieActors []entity.MovieActor

	for _, actor := range newMovie.Actors {
		newMovieActor, newMovieActorError := entity.NewMovieActor(newMovie.ID, actor.ID)
		if newMovieActorError != nil {
			problemsDetails = append(problemsDetails, newMovieActorError...)
		}

		movieActors = append(movieActors, *newMovieActor)
	}

	if len(movieActors) > 0 {
		movieActorsToAddError := cm.MovieActorRepository.CreateMany(&movieActors)
		if movieActorsToAddError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao criar o(a)s atores(atrizes) do filme",
				Status:   http.StatusInternalServerError,
				Detail:   movieActorsToAddError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar o(a)s atores(atrizes) do filme", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		}
	}

	return util.ProblemDetailsOutputDTO{}, newMovie
}

func TreatmentDirectors(inputDirectors []DirectorDTO, cm *CreateMovieUseCase, newMovie *entity.Movie) (util.ProblemDetailsOutputDTO, *entity.Movie) {
	var directorsToCheck []string
	var directorsToAdd []DirectorDTO
	var problemsDetails []util.ProblemDetails
	var imagesToAdd []entity.Image

	for _, director := range inputDirectors {
		if director.DirectorID == "" {
			directorsToAdd = append(directorsToAdd, director)
		} else {
			directorsToCheck = append(directorsToCheck, director.DirectorID)
		}
	}

	if len(directorsToCheck) > 0 {
		doTheseDirectorsExist, existingDirectors, manyDirectorsError := cm.DirectorRepository.DoTheseDirectorsExist(directorsToCheck)
		if manyDirectorsError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao resgatar o(a)s diretores(as) pelos ids",
				Status:   http.StatusInternalServerError,
				Detail:   manyDirectorsError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar o(a)s diretores(as) pelos ids", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		} else if !doTheseDirectorsExist {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeValidationError,
				Title:    "Um ou mais diretores(as) não encontrados",
				Status:   http.StatusConflict,
				Detail:   "Um ou mais ids do(a)s diretores(as) não retornou resultado",
				Instance: util.RFC409,
			})

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		}

		if len(existingDirectors) > 0 {
			newMovie.AddDirectors(existingDirectors)
		}
	}

	var newDirectors []entity.Director

	if len(directorsToAdd) > 0 {
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

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do(a) diretor(a): "+directorToAdd.Name, "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
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

			if len(problemsDetails) > 0 {
				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}

			imagesToAdd = append(imagesToAdd, *directorToAddImage)
			newDirectors = append(newDirectors, *newDirector)
		}

		if len(imagesToAdd) > 0 {
			imagesToAddError := cm.ImageRepository.CreateMany(&imagesToAdd)
			if imagesToAddError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao criar imagens",
					Status:   http.StatusInternalServerError,
					Detail:   imagesToAddError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar as imagens dos(as) diretores(as)", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}
		}

		if len(newDirectors) > 0 {
			directorsToAddError := cm.DirectorRepository.CreateMany(&newDirectors)
			if directorsToAddError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao criar o(a)s diretores(as)",
					Status:   http.StatusInternalServerError,
					Detail:   directorsToAddError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar o(a)s diretores(as)", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}

			newMovie.AddDirectors(newDirectors)
		}
	}

	var movieDirectors []entity.MovieDirector

	for _, director := range newMovie.Directors {
		newMovieDirector, newMovieDirectorError := entity.NewMovieDirector(newMovie.ID, director.ID)
		if newMovieDirectorError != nil {
			problemsDetails = append(problemsDetails, newMovieDirectorError...)
		}

		movieDirectors = append(movieDirectors, *newMovieDirector)
	}

	if len(movieDirectors) > 0 {
		movieDirectorsToAddError := cm.MovieDirectorRepository.CreateMany(&movieDirectors)
		if movieDirectorsToAddError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao criar o(a)s diretores(as)",
				Status:   http.StatusInternalServerError,
				Detail:   movieDirectorsToAddError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar o(a)s diretores(as) do filme", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		}
	}

	return util.ProblemDetailsOutputDTO{}, newMovie
}

func TreatmentWriters(inputWriters []WriterDTO, cm *CreateMovieUseCase, newMovie *entity.Movie) (util.ProblemDetailsOutputDTO, *entity.Movie) {
	var writersToCheck []string
	var writersToAdd []WriterDTO
	var problemsDetails []util.ProblemDetails
	var imagesToAdd []entity.Image

	for _, writer := range inputWriters {
		if writer.WriterID == "" {
			writersToAdd = append(writersToAdd, writer)
		} else {
			writersToCheck = append(writersToCheck, writer.WriterID)
		}
	}

	if len(writersToCheck) > 0 {
		doTheseWritersExist, existingWriters, manyWritersError := cm.WriterRepository.DoTheseWritersExist(writersToCheck)
		if manyWritersError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao resgatar o(a)s escritores(as) pelos ids",
				Status:   http.StatusInternalServerError,
				Detail:   manyWritersError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os escritores(as) pelos ids", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		} else if !doTheseWritersExist {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeValidationError,
				Title:    "Um(a) ou mais escritores(as) não encontrados",
				Status:   http.StatusConflict,
				Detail:   "Um ou mais ids do(a)s escritores(as) não retornou resultado",
				Instance: util.RFC409,
			})

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		}

		if len(existingWriters) > 0 {
			newMovie.AddWriters(existingWriters)
		}
	}

	var newWriters []entity.Writer

	if len(writersToAdd) > 0 {
		for _, writerToAdd := range writersToAdd {
			_, writerToAddName, writerToAddExtension, writerToAddSize, profileImageError := service.MoveFile(writerToAdd.ImageFile, writerToAdd.ImageHandler)
			if profileImageError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao mover a imagem do(a) escritor(a): " + writerToAdd.Name,
					Status:   http.StatusInternalServerError,
					Detail:   profileImageError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do(a) escritor(a): "+writerToAdd.Name, "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
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

			if len(problemsDetails) > 0 {
				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}

			imagesToAdd = append(imagesToAdd, *writerToAddImage)
			newWriters = append(newWriters, *newWriter)
		}

		if len(imagesToAdd) > 0 {
			imagesToAddError := cm.ImageRepository.CreateMany(&imagesToAdd)
			if imagesToAddError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao criar imagens",
					Status:   http.StatusInternalServerError,
					Detail:   imagesToAddError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar as imagens dos(as) escritores(as)", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}
		}

		if len(newWriters) > 0 {
			writersToAddError := cm.WriterRepository.CreateMany(&newWriters)
			if writersToAddError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     util.TypeInternalServerError,
					Title:    "Erro ao criar o(a)s escritores(as)",
					Status:   http.StatusInternalServerError,
					Detail:   writersToAddError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar o(a)s escritores(as)", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

				return util.ProblemDetailsOutputDTO{
					ProblemDetails: problemsDetails,
				}, nil
			}

			newMovie.AddWriters(newWriters)
		}
	}

	var movieWriters []entity.MovieWriter

	for _, writer := range newMovie.Writers {
		newMovieWriter, newMovieWriterError := entity.NewMovieWriter(newMovie.ID, writer.ID)
		if newMovieWriterError != nil {
			problemsDetails = append(problemsDetails, newMovieWriterError...)
		}

		movieWriters = append(movieWriters, *newMovieWriter)
	}

	if len(movieWriters) > 0 {
		movieWritersToAddError := cm.MovieWriterRepository.CreateMany(&movieWriters)
		if movieWritersToAddError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     util.TypeInternalServerError,
				Title:    "Erro ao criar o(a)s escritores(as)",
				Status:   http.StatusInternalServerError,
				Detail:   movieWritersToAddError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar o(a)s escritores(as) do filme", "CreateMovieUseCase", "Use Cases", util.TypeInternalServerError)

			return util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}, nil
		}
	}

	return util.ProblemDetailsOutputDTO{}, newMovie
}
