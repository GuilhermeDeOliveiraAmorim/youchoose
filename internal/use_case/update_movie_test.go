package usecase_test

import (
	"testing"
	"youchoose/internal/entity"
	usecase "youchoose/internal/use_case"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMovieUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockActorRepository := mock.NewMockActorRepositoryInterface(ctrl)
	mockChooserRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	mockDirectorRepository := mock.NewMockDirectorRepositoryInterface(ctrl)
	mockGenreRepository := mock.NewMockGenreRepositoryInterface(ctrl)
	mockImageRepository := mock.NewMockImageRepositoryInterface(ctrl)
	mockMovieActorRepository := mock.NewMockMovieActorRepositoryInterface(ctrl)
	mockMovieDirectorRepository := mock.NewMockMovieDirectorRepositoryInterface(ctrl)
	mockMovieGenreRepository := mock.NewMockMovieGenreRepositoryInterface(ctrl)
	mockMovieRepository := mock.NewMockMovieRepositoryInterface(ctrl)
	mockMovieWriterRepository := mock.NewMockMovieWriterRepositoryInterface(ctrl)
	mockWriterRepository := mock.NewMockWriterRepositoryInterface(ctrl)

	updateMovieUseCase := usecase.NewUpdateMovieUseCase(
		mockActorRepository,
		mockChooserRepository,
		mockDirectorRepository,
		mockGenreRepository,
		mockImageRepository,
		mockMovieActorRepository,
		mockMovieDirectorRepository,
		mockMovieGenreRepository,
		mockMovieRepository,
		mockMovieWriterRepository,
		mockWriterRepository,
	)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "Aracaju", State: "SE", Country: "Brasil"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	nationality, _ := valueobject.NewNationality("Brasil", "🇧🇷")

	movie, _ := entity.NewMovie("Inception", *nationality, 2010, imageID)

	genreID := uuid.NewString()
	directorID := uuid.NewString()
	actorID := uuid.NewString()
	writerID := uuid.NewString()

	var genres []entity.Genre
	var directors []entity.Director
	var actors []entity.Actor
	var writers []entity.Writer

	genres = append(genres, entity.Genre{
		Name: "Action",
	})

	directors = append(directors, entity.Director{
		Name:        "Action",
		Nationality: nationality,
	})

	actors = append(actors, entity.Actor{
		Name:        "Action",
		Nationality: nationality,
	})

	writers = append(writers, entity.Writer{
		Name:        "Action",
		Nationality: nationality,
	})

	movie.AddGenres(genres)
	movie.AddDirectors(directors)
	movie.AddActors(actors)
	movie.AddWriters(writers)

	input := usecase.UpdateMovieInputDTO{
		ChooserID:   chooser.ID,
		MovieID:     movie.ID,
		Title:       "New Movie Title",
		CountryName: "Albania",
		Flag:        "🇦🇱",
		ReleaseYear: movie.ReleaseYear,
		ImageID:     imageID,
		Genres: []usecase.GenreDTO{
			{
				GenreID: genreID,
			},
		},
		Directors: []usecase.DirectorDTO{
			{
				DirectorID: directorID,
			},
		},
		Actors: []usecase.ActorDTO{
			{
				ActorID: actorID,
			},
		},
		Writers: []usecase.WriterDTO{
			{
				WriterID: writerID,
			},
		},
	}

	mockChooserRepository.EXPECT().GetByID(gomock.Any()).Return(true, *chooser, nil)
	mockMovieRepository.EXPECT().GetByID(gomock.Any()).Return(true, *movie, nil)
	mockGenreRepository.EXPECT().DoTheseGenresAreIncludedInTheMovie(gomock.Any(), gomock.Any()).Return(true, nil, nil)
	mockDirectorRepository.EXPECT().DoTheseDirectorsAreIncludedInTheMovie(gomock.Any(), gomock.Any()).Return(true, nil, nil)
	mockActorRepository.EXPECT().DoTheseActorsAreIncludedInTheMovie(gomock.Any(), gomock.Any()).Return(true, nil, nil)
	mockWriterRepository.EXPECT().DoTheseWritersAreIncludedInTheMovie(gomock.Any(), gomock.Any()).Return(true, nil, nil)
	mockMovieRepository.EXPECT().Update(gomock.Any()).Return(nil)

	output, _ := updateMovieUseCase.Execute(input)

	assert.NotNil(t, output, "Deve haver uma saída")
	assert.Equal(t, input.MovieID, output.ID, "ID do filme deve ser igual ao fornecido na entrada")
	assert.Equal(t, input.Title, output.Title, "Título do filme deve ser igual ao fornecido na entrada")
	assert.Equal(t, input.ReleaseYear, output.ReleaseYear, "Ano de lançamento do filme deve ser igual ao fornecido na entrada")
	assert.Equal(t, input.ImageID, output.ImageID, "ID da imagem do filme deve ser igual ao fornecido na entrada")
	assert.Len(t, output.Genres, 1, "Deve haver exatamente um gênero associado ao filme")
	assert.Len(t, output.Directors, 1, "Deve haver exatamente um diretor associado ao filme")
	assert.Len(t, output.Actors, 1, "Deve haver exatamente um ator associado ao filme")
	assert.Len(t, output.Writers, 1, "Deve haver exatamente um escritor associado ao filme")
}
