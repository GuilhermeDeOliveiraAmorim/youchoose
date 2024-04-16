package entity

import (
	"net/http"
	"strconv"
	"time"

	"youchoose/internal/util"

	valueobject "youchoose/internal/value_object"
)

type Movie struct {
	SharedEntity
	Title       string                  `json:"title"`
	Nationality valueobject.Nationality `json:"nationality"`
	ReleaseYear int                     `json:"release_year"`
	ImageID     string                  `json:"image_id"`
	Votes       int                     `json:"votes"`
	Genres      []Genre                 `json:"genres"`
	Directors   []Director              `json:"directors"`
	Actors      []Actor                 `json:"actors"`
	Writers     []Writer                `json:"writers"`
}

func NewMovie(title string, nationality valueobject.Nationality, releaseYear int, imageID string) (*Movie, []util.ProblemDetails) {
	validationErrors := ValidateMovie(title, nationality, releaseYear, imageID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Movie{
		SharedEntity: *NewSharedEntity(),
		Title:        title,
		Nationality:  nationality,
		ReleaseYear:  releaseYear,
		ImageID:      imageID,
		Votes:        0,
	}, nil
}

func ValidateMovie(title string, nationality valueobject.Nationality, releaseYear int, imageID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if title == "" || len(title) > 255 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidName,
			Status:   http.StatusBadRequest,
			Detail:   "O título do filme não pode estar vazio e deve ter no máximo 255 caracteres.",
			Instance: util.RFC400,
		})
	}

	yearNow := time.Now().Year()

	if releaseYear < 1800 || releaseYear > yearNow {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidYear,
			Status:   http.StatusBadRequest,
			Detail:   util.MovieErrorDetailInvalidYear + strconv.Itoa(yearNow),
			Instance: util.RFC400,
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.SharedErrorTitleInvalidImageID,
			Status:   http.StatusBadRequest,
			Detail:   util.MovieErrorDetailEmptyImageID,
			Instance: util.RFC400,
		})
	}

	return validationErrors
}

func (m *Movie) IncrementVotes() {
	m.UpdatedAt = time.Now()
	m.Votes++
}

func (m *Movie) AddActors(newActors []Actor) {
	m.UpdatedAt = time.Now()
	m.Actors = append(m.Actors, newActors...)
}

func (m *Movie) AddWriters(newWriters []Writer) {
	m.UpdatedAt = time.Now()
	m.Writers = append(m.Writers, newWriters...)
}

func (m *Movie) AddDirectors(newDirectors []Director) {
	m.UpdatedAt = time.Now()
	m.Directors = append(m.Directors, newDirectors...)
}

func (m *Movie) AddGenres(newGenres []Genre) {
	m.UpdatedAt = time.Now()
	m.Genres = append(m.Genres, newGenres...)
}

func (m *Movie) RemoveActors(actorsToRemove []Actor) {
	var updatedActors []Actor

	for _, existingActor := range m.Actors {
		found := false
		for _, actorToRemove := range actorsToRemove {
			if existingActor.ID == actorToRemove.ID {
				found = true
				break
			}
		}

		if !found {
			updatedActors = append(updatedActors, existingActor)
		}
	}

	if len(updatedActors) > 0 {
		m.UpdatedAt = time.Now()
		m.Actors = updatedActors
	}
}

func (m *Movie) RemoveWriters(writersToRemove []Writer) {
	var updatedWriters []Writer

	for _, existingWriter := range m.Writers {
		found := false
		for _, writerToRemove := range writersToRemove {
			if existingWriter.ID == writerToRemove.ID {
				found = true
				break
			}
		}

		if !found {
			updatedWriters = append(updatedWriters, existingWriter)
		}
	}

	if len(updatedWriters) > 0 {
		m.UpdatedAt = time.Now()
		m.Writers = updatedWriters
	}
}

func (m *Movie) RemoveDirectors(directorsToRemove []Director) {
	var updatedDirectors []Director

	for _, existingDirector := range m.Directors {
		found := false
		for _, directorToRemove := range directorsToRemove {
			if existingDirector.ID == directorToRemove.ID {
				found = true
				break
			}
		}

		if !found {
			updatedDirectors = append(updatedDirectors, existingDirector)
		}
	}

	if len(updatedDirectors) > 0 {
		m.UpdatedAt = time.Now()
		m.Directors = updatedDirectors
	}
}

func (m *Movie) RemoveGenres(genresToRemove []Genre) {
	var updatedGenres []Genre

	for _, existingGenre := range m.Genres {
		found := false
		for _, genreToRemove := range genresToRemove {
			if existingGenre.ID == genreToRemove.ID {
				found = true
				break
			}
		}

		if !found {
			updatedGenres = append(updatedGenres, existingGenre)
		}
	}

	if len(updatedGenres) > 0 {
		m.UpdatedAt = time.Now()
		m.Genres = updatedGenres
	}
}
