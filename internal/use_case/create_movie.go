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

type Genre struct {
	GenreID      string                `json:"genre_id"`
	Name         string                `json:"name"`
	ImageFile    multipart.File        `json:"genre_image_file"`
	ImageHandler *multipart.FileHeader `json:"genre_image_handler"`
}

type Director struct {
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

type Actor struct {
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

type Writer struct {
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
	Genres       []Genre               `json:"genres"`
	Directors    []Director            `json:"directors"`
	Actors       []Actor               `json:"actors"`
	Writers      []Writer              `json:"writers"`
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
			Type:     "Bad Request",
			Title:    "Gêneros não informados",
			Status:   http.StatusBadRequest,
			Detail:   "Nenhum gênero informado",
			Instance: util.RFC400,
		})
	}

	if len(input.Directors) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Diretores não informados",
			Status:   http.StatusBadRequest,
			Detail:   "Nenhum diretor informado",
			Instance: util.RFC400,
		})
	}

	if len(input.Actors) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Atores não informados",
			Status:   http.StatusBadRequest,
			Detail:   "Nenhum ator informado",
			Instance: util.RFC400,
		})
	}

	if len(input.Writers) == 0 {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Bad Request",
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

	var imagesToAdd []entity.Image

	_, movieImageName, movieImageExtension, movieImageSize, profileImageError := service.MoveFile(input.ImageFile, input.ImageHandler)
	if profileImageError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao mover a imagem do filme",
			Status:   http.StatusInternalServerError,
			Detail:   profileImageError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do filme", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

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

	imagesToAdd = append(imagesToAdd, *movieImage)

	var genresToCheck []string
	var genresToAdd []Genre

	for _, genre := range input.Genres {
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
				Type:     "Internal Server Error",
				Title:    "Erro ao resgatar os gêneros pelos ids",
				Status:   http.StatusInternalServerError,
				Detail:   manyGenresError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os gêneros pelos ids", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		} else if !doTheseGenresExist {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Validation Error",
				Title:    "Um ou mais gêneros não encontrados",
				Status:   http.StatusConflict,
				Detail:   "Um ou mais ids dos gêneros não retornou resultado",
				Instance: util.RFC409,
			})

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		if len(existingGenres) > 0 {
			newMovie.AddGenres(existingGenres)
		}
	}

	if len(genresToAdd) > 0 {
		var newGenres []entity.Genre

		for _, genreToAdd := range genresToAdd {
			_, genreToAddName, genreToAddExtension, genreToAddSize, profileImageError := service.MoveFile(genreToAdd.ImageFile, genreToAdd.ImageHandler)
			if profileImageError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     "Internal Server Error",
					Title:    "Erro ao mover a imagem do gênero: " + genreToAdd.Name,
					Status:   http.StatusInternalServerError,
					Detail:   profileImageError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do gênero: "+genreToAdd.Name, "CreateMovieUseCase", "Use Cases", "Internal Server Error")

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

		newMovie.AddGenres(newGenres)
	}

	var directorsToCheck []string
	var directorsToAdd []Director

	for _, director := range input.Directors {
		if director.DirectorID == "" {
			directorsToAdd = append(directorsToAdd, director)
		} else {
			directorsToCheck = append(directorsToCheck, director.DirectorID)
		}
	}

	if len(directorsToCheck) > 0 {
		doTheseDirectorsExist, existingDirectors, manyDirectorsError := cm.DirectorRepository.DoTheseDirectorsExist(genresToCheck)
		if manyDirectorsError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Internal Server Error",
				Title:    "Erro ao resgatar os diretores pelos ids",
				Status:   http.StatusInternalServerError,
				Detail:   manyDirectorsError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os diretores pelos ids", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		} else if !doTheseDirectorsExist {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Validation Error",
				Title:    "Um ou mais diretores não encontrados",
				Status:   http.StatusConflict,
				Detail:   "Um ou mais ids dos diretores não retornou resultado",
				Instance: util.RFC409,
			})

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		if len(existingDirectors) > 0 {
			newMovie.AddDirectors(existingDirectors)
		}
	}

	if len(directorsToAdd) > 0 {
		var newDirectors []entity.Director

		for _, directorToAdd := range directorsToAdd {
			_, directorToAddName, directorToAddExtension, directorToAddSize, profileImageError := service.MoveFile(directorToAdd.ImageFile, directorToAdd.ImageHandler)
			if profileImageError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     "Internal Server Error",
					Title:    "Erro ao mover a imagem do diretor: " + directorToAdd.Name,
					Status:   http.StatusInternalServerError,
					Detail:   profileImageError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do diretor: "+directorToAdd.Name, "CreateMovieUseCase", "Use Cases", "Internal Server Error")

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

		newMovie.AddDirectors(newDirectors)
	}

	//Actor

	var actorsToCheck []string
	var actorsToAdd []Actor

	for _, actor := range input.Actors {
		if actor.ActorID == "" {
			actorsToAdd = append(actorsToAdd, actor)
		} else {
			actorsToCheck = append(actorsToCheck, actor.ActorID)
		}
	}

	if len(actorsToCheck) > 0 {
		doTheseActorsExist, existingActors, manyActorsError := cm.ActorRepository.DoTheseActorsExist(genresToCheck)
		if manyActorsError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Internal Server Error",
				Title:    "Erro ao resgatar os(as) atores(atrizes) pelos ids",
				Status:   http.StatusInternalServerError,
				Detail:   manyActorsError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os(as) atores(atrizes) pelos ids", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		} else if !doTheseActorsExist {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Validation Error",
				Title:    "Um(a) ou mais atores(atrizes) não encontrados(as)",
				Status:   http.StatusConflict,
				Detail:   "Um ou mais ids dos(as) atores(atrizes) não retornou resultado",
				Instance: util.RFC409,
			})

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		if len(existingActors) > 0 {
			newMovie.AddActors(existingActors)
		}
	}

	if len(actorsToAdd) > 0 {
		var newActors []entity.Actor

		for _, actorToAdd := range actorsToAdd {
			_, actorToAddName, actorToAddExtension, actorToAddSize, profileImageError := service.MoveFile(actorToAdd.ImageFile, actorToAdd.ImageHandler)
			if profileImageError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     "Internal Server Error",
					Title:    "Erro ao mover a imagem do(a) ator(atriz): " + actorToAdd.Name,
					Status:   http.StatusInternalServerError,
					Detail:   profileImageError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do(a) ator(atriz): "+actorToAdd.Name, "CreateMovieUseCase", "Use Cases", "Internal Server Error")

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

		newMovie.AddActors(newActors)
	}

	//Writer

	var writersToCheck []string
	var writersToAdd []Writer

	for _, writer := range input.Writers {
		if writer.WriterID == "" {
			writersToAdd = append(writersToAdd, writer)
		} else {
			writersToCheck = append(writersToCheck, writer.WriterID)
		}
	}

	if len(writersToCheck) > 0 {
		doTheseWritersExist, existingWriters, manyWritersError := cm.WriterRepository.DoTheseWritersExist(genresToCheck)
		if manyWritersError != nil {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Internal Server Error",
				Title:    "Erro ao resgatar os(as) escritores(escritoras) pelos ids",
				Status:   http.StatusInternalServerError,
				Detail:   manyWritersError.Error(),
				Instance: util.RFC503,
			})

			util.NewLoggerError(http.StatusInternalServerError, "Erro ao resgatar os(as) escritores(escritoras) pelos ids", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		} else if !doTheseWritersExist {
			problemsDetails = append(problemsDetails, util.ProblemDetails{
				Type:     "Validation Error",
				Title:    "Um(a) ou mais escritores(escritoras) não encontrados(as)",
				Status:   http.StatusConflict,
				Detail:   "Um ou mais ids dos(as) escritores(escritoras) não retornou resultado",
				Instance: util.RFC409,
			})

			return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
				ProblemDetails: problemsDetails,
			}
		}

		if len(existingWriters) > 0 {
			newMovie.AddWriters(existingWriters)
		}
	}

	if len(writersToAdd) > 0 {
		var newWriters []entity.Writer

		for _, writerToAdd := range writersToAdd {
			_, writerToAddName, writerToAddExtension, writerToAddSize, profileImageError := service.MoveFile(writerToAdd.ImageFile, writerToAdd.ImageHandler)
			if profileImageError != nil {
				problemsDetails = append(problemsDetails, util.ProblemDetails{
					Type:     "Internal Server Error",
					Title:    "Erro ao mover a imagem do(a) escritor(escritora): " + writerToAdd.Name,
					Status:   http.StatusInternalServerError,
					Detail:   profileImageError.Error(),
					Instance: util.RFC503,
				})

				util.NewLoggerError(http.StatusInternalServerError, "Erro ao mover a imagem do(a) escritor(escritora): "+writerToAdd.Name, "CreateMovieUseCase", "Use Cases", "Internal Server Error")

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

		newMovie.AddWriters(newWriters)
	}

	if len(problemsDetails) > 0 {
		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	var movieActors []entity.MovieActor
	var movieDirectors []entity.MovieDirector
	var movieGenres []entity.MovieGenre
	var movieWriters []entity.MovieWriter

	for _, actor := range newMovie.Actors {
		newMovieActor, newMovieActorError := entity.NewMovieActor(newMovie.ID, actor.ID)
		if newMovieActorError != nil {
			problemsDetails = append(problemsDetails, newMovieActorError...)
		}

		movieActors = append(movieActors, *newMovieActor)
	}

	for _, director := range newMovie.Directors {
		newMovieDirector, newMovieDirectorError := entity.NewMovieDirector(newMovie.ID, director.ID)
		if newMovieDirectorError != nil {
			problemsDetails = append(problemsDetails, newMovieDirectorError...)
		}

		movieDirectors = append(movieDirectors, *newMovieDirector)
	}

	for _, genre := range newMovie.Genres {
		newMovieGenre, newMovieGenreError := entity.NewMovieGenre(newMovie.ID, genre.ID)
		if newMovieGenreError != nil {
			problemsDetails = append(problemsDetails, newMovieGenreError...)
		}

		movieGenres = append(movieGenres, *newMovieGenre)
	}

	for _, writer := range newMovie.Writers {
		newMovieWriter, newMovieWriterError := entity.NewMovieWriter(newMovie.ID, writer.ID)
		if newMovieWriterError != nil {
			problemsDetails = append(problemsDetails, newMovieWriterError...)
		}

		movieWriters = append(movieWriters, *newMovieWriter)
	}

	if len(problemsDetails) > 0 {
		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	imagesToAddError := cm.ImageRepository.CreateMany(&imagesToAdd)
	if imagesToAddError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao criar as imagens",
			Status:   http.StatusInternalServerError,
			Detail:   imagesToAddError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar as imagens", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	movieActorsError := cm.MovieActorRepository.CreateMany(&movieActors)
	if movieActorsError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao criar os(as) actores(atrizes)",
			Status:   http.StatusInternalServerError,
			Detail:   movieActorsError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar os(as) actores(atrizes)", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	movieDirectorsError := cm.MovieDirectorRepository.CreateMany(&movieDirectors)
	if movieDirectorsError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao criar os(as) diretores(as)",
			Status:   http.StatusInternalServerError,
			Detail:   movieDirectorsError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar os(as) diretores(as)", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	movieGenresError := cm.MovieGenreRepository.CreateMany(&movieGenres)
	if movieGenresError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao criar os gêneros",
			Status:   http.StatusInternalServerError,
			Detail:   movieGenresError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar os gêneros", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	movieWritersError := cm.MovieWriterRepository.CreateMany(&movieWriters)
	if movieWritersError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao criar os(as) escritores(as)",
			Status:   http.StatusInternalServerError,
			Detail:   movieWritersError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar os(as) escritores(as)", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	newMovieError := cm.MovieRepository.Create(newMovie)
	if newMovieError != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao criar o filme",
			Status:   http.StatusInternalServerError,
			Detail:   newMovieError.Error(),
			Instance: util.RFC503,
		})

		util.NewLoggerError(http.StatusInternalServerError, "Erro ao criar o filme", "CreateMovieUseCase", "Use Cases", "Internal Server Error")

		return MovieOutputDTO{}, util.ProblemDetailsOutputDTO{
			ProblemDetails: problemsDetails,
		}
	}

	output := NewMovieOutputDTO(*newMovie)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
