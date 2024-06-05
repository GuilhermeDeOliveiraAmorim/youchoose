package usecase

import (
	"time"
	"youchoose/internal/entity"
	valueobject "youchoose/internal/value_object"
)

type ChooserOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	Day       int       `json:"day"`
	Month     int       `json:"month"`
	Year      int       `json:"year"`
	ImageID   string    `json:"image_id"`
}

func NewChooserOutputDTO(chooser entity.Chooser) ChooserOutputDTO {
	output := ChooserOutputDTO{
		ID:        chooser.ID,
		CreatedAt: chooser.CreatedAt,
		Name:      chooser.Name,
		City:      chooser.Address.City,
		State:     chooser.Address.State,
		Country:   chooser.Address.Country,
		Day:       chooser.BirthDate.Day,
		Month:     chooser.BirthDate.Month,
		Year:      chooser.BirthDate.Year,
		ImageID:   chooser.ImageID,
	}

	return output
}

type GenreOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	ImageID   string    `json:"image_id"`
}

type DirectorOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	ImageID   string    `json:"image_id"`
}

type ActorOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	ImageID   string    `json:"image_id"`
}

type WriterOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	ImageID   string    `json:"image_id"`
}

type MovieOutputDTO struct {
	ID          string                  `json:"id"`
	CreatedAt   time.Time               `json:"created_at"`
	Title       string                  `json:"title"`
	Nationality valueobject.Nationality `json:"nationality"`
	ReleaseYear int                     `json:"release_year"`
	ImageID     string                  `json:"image_id"`
	Votes       int                     `json:"votes"`
	Genres      []GenreOutputDTO        `json:"genres"`
	Directors   []DirectorOutputDTO     `json:"directors"`
	Actors      []ActorOutputDTO        `json:"actors"`
	Writers     []WriterOutputDTO       `json:"writers"`
}

func NewMovieOutputDTO(movie entity.Movie) MovieOutputDTO {
	var genresOutputDTO []GenreOutputDTO
	var directorsOutputDTO []DirectorOutputDTO
	var actorsOutputDTO []ActorOutputDTO
	var writersOutputDTO []WriterOutputDTO

	for _, genre := range movie.Genres {
		genresOutputDTO = append(genresOutputDTO, GenreOutputDTO{
			ID:        genre.ID,
			CreatedAt: genre.CreatedAt,
			Name:      genre.Name,
		})
	}

	for _, director := range movie.Directors {
		directorsOutputDTO = append(directorsOutputDTO, DirectorOutputDTO{
			ID:        director.ID,
			CreatedAt: director.CreatedAt,
			Name:      director.Name,
			ImageID:   director.ImageID,
		})
	}

	for _, actor := range movie.Actors {
		actorsOutputDTO = append(actorsOutputDTO, ActorOutputDTO{
			ID:        actor.ID,
			CreatedAt: actor.CreatedAt,
			Name:      actor.Name,
			ImageID:   actor.ImageID,
		})
	}

	for _, writer := range movie.Writers {
		writersOutputDTO = append(writersOutputDTO, WriterOutputDTO{
			ID:        writer.ID,
			CreatedAt: writer.CreatedAt,
			Name:      writer.Name,
			ImageID:   writer.ImageID,
		})
	}

	output := MovieOutputDTO{
		ID:          movie.ID,
		Title:       movie.Title,
		Nationality: movie.Nationality,
		ReleaseYear: movie.ReleaseYear,
		ImageID:     movie.ImageID,
		Votes:       movie.Votes,
		Genres:      genresOutputDTO,
		Directors:   directorsOutputDTO,
		Actors:      actorsOutputDTO,
		Writers:     writersOutputDTO,
	}

	return output
}

type ListOutputDTO struct {
	ListID         string           `json:"list_id"`
	CreatedAt      time.Time        `json:"created_at"`
	Title          string           `json:"title"`
	Description    string           `json:"description"`
	ChooserID      string           `json:"chooser_id"`
	ProfileImageID string           `json:"profile_image_id"`
	CoverImageID   string           `json:"cover_image_id"`
	Votes          int              `json:"votes"`
	Movies         []MovieOutputDTO `json:"movies"`
}

func NewListOutputDTO(list entity.List) ListOutputDTO {
	allGetMoviesOutputDTO := []MovieOutputDTO{}

	for _, movie := range list.Movies {
		outputMovie := NewMovieOutputDTO(movie)

		allGetMoviesOutputDTO = append(allGetMoviesOutputDTO, MovieOutputDTO{
			ID:          movie.ID,
			CreatedAt:   movie.CreatedAt,
			Title:       movie.Title,
			Nationality: movie.Nationality,
			ReleaseYear: movie.ReleaseYear,
			ImageID:     movie.ImageID,
			Votes:       movie.Votes,
			Genres:      outputMovie.Genres,
			Directors:   outputMovie.Directors,
			Actors:      outputMovie.Actors,
			Writers:     outputMovie.Writers,
		})
	}

	output := ListOutputDTO{
		ListID:         list.ID,
		CreatedAt:      list.CreatedAt,
		Title:          list.Title,
		Description:    list.Description,
		ChooserID:      list.ChooserID,
		ProfileImageID: list.ProfileImageID,
		CoverImageID:   list.CoverImageID,
		Votes:          list.Votes,
		Movies:         allGetMoviesOutputDTO,
	}

	return output
}
