package factory

import (
	repository "youchoose/internal/infra/repository"
	"youchoose/internal/service"
	usecase "youchoose/internal/use_case"

	"gorm.io/gorm"
)

type MovieFactory struct {
	CreateMovie         *usecase.CreateMovieUseCase
	CreateMovieWithIMDB *service.CreateMovieWithIMDBIdService
}

func NewMovieFactory(db *gorm.DB) *MovieFactory {
	chooserRepository := repository.NewChooserRepository(db)
	movieRepository := repository.NewMovieRepository(db)
	imageRepository := repository.NewImageRepository(db)
	genreRepository := repository.NewGenreRepository(db)
	directorRepository := repository.NewDirectorRepository(db)
	actorRepository := repository.NewActorRepository(db)
	writerRepository := repository.NewWriterRepository(db)
	movieGenreRepository := repository.NewMovieGenreRepository(db)
	movieDirectorRepository := repository.NewMovieDirectorRepository(db)
	movieActorRepository := repository.NewMovieActorRepository(db)
	movieWriterRepository := repository.NewMovieWriterRepository(db)
	imdbRepository := repository.NewIMDBRepository(db)

	createMovie := usecase.NewCreateMovieUseCase(
		chooserRepository,
		movieRepository,
		imageRepository,
		genreRepository,
		directorRepository,
		actorRepository,
		writerRepository,
		movieGenreRepository,
		movieDirectorRepository,
		movieActorRepository,
		movieWriterRepository,
	)

	createMovieWithIMDB := service.NewCreateMovieWithIMDBIdService(
		imdbRepository,
	)

	return &MovieFactory{
		CreateMovie:         createMovie,
		CreateMovieWithIMDB: createMovieWithIMDB,
	}
}
