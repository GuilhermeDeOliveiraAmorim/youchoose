package internal

import "net/http"

type Movie struct {
	SharedEntity
	Title        string       `json:"title"`
	Nationality  Nationality  `json:"nationality"`
	Genres       []Genre      `json:"genres"`
	Directors    []Director   `json:"directors"`
	Actors       []Actor      `json:"actors"`
	Writers      []Writer     `json:"writers"`
	ReleaseYear  int          `json:"release_year"`
	ImageID      string       `json:"image_id"`
	Votes        int          `json:"votes"`
}

func NewMovie(title string, nationality Nationality, genres []Genre, directors []Director, actors []Actor, writers []Writer, releaseYear int, imageID string) (*Movie, []ProblemDetails) {
	validationErrors := ValidateMovie(title, nationality, genres, directors, actors, writers, releaseYear, imageID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Movie{
		SharedEntity: *NewSharedEntity(),
		Title:        title,
		Nationality:  nationality,
		Genres:       genres,
		Directors:    directors,
		Actors:       actors,
		Writers:      writers,
		ReleaseYear:  releaseYear,
		ImageID:      imageID,
		Votes:        0,
	}, nil
}

func ValidateMovie(title string, nationality Nationality, genres []Genre, directors []Director, actors []Actor, writers []Writer, releaseYear int, imageID string) []ProblemDetails {
	var validationErrors []ProblemDetails

	if title == "" || len(title) > 255 {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Título do filme inválido",
			Status: http.StatusBadRequest,
			Detail: "O título do filme não pode estar vazio e deve ter no máximo 255 caracteres.",
		})
	}

	return validationErrors
}

func (m *Movie) IncrementVotes() {
	m.Votes++
}

func (m *Movie) AddActors(newActors []Actor) {
	m.Actors = append(m.Actors, newActors...)
}

func (m *Movie) AddWriters(newWriters []Writer) {
	m.Writers = append(m.Writers, newWriters...)
}

func (m *Movie) AddDirectors(newDirectors []Director) {
	m.Directors = append(m.Directors, newDirectors...)
}

func (m *Movie) AddGenres(newGenres []Genre) {
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
		m.Genres = updatedGenres
	}
}