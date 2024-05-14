package factory

import (
	repository "youchoose/internal/infra/repository"
	usecase "youchoose/internal/use_case"

	"gorm.io/gorm"
)

type MovieFactory struct {
	CreateMovie *usecase.CreateMovieUseCase
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

	return &MovieFactory{
		CreateMovie: createMovie,
	}
}
