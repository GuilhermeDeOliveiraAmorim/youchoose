package entity

import (
	"net/http"
	"time"

	"youchoose/internal/util"
)

type List struct {
	SharedEntity
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	ChooserID      string  `json:"chooser_id"`
	ProfileImageID string  `json:"profile_image_id"`
	CoverImageID   string  `json:"cover_image_id"`
	Movies         []Movie `json:"movies"`
	Votes          int     `json:"votes"`
}

func NewList(title, description, profileImageID, coverImageID, chooserID string) (*List, []util.ProblemDetails) {
	validationErrors := ValidateList(title, description, chooserID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &List{
		SharedEntity:   *NewSharedEntity(),
		ProfileImageID: profileImageID,
		CoverImageID:   coverImageID,
		Title:          title,
		Description:    description,
		Movies:         []Movie{},
		ChooserID:      chooserID,
		Votes:          0,
	}, nil
}

func (l *List) IncrementVotes() {
	l.UpdatedAt = time.Now()

	l.Votes++
}

func (l *List) ChangeProfileImageID(profileImageID string) {
	l.UpdatedAt = time.Now()

	l.ProfileImageID = profileImageID
}

func (l *List) ChangeCoverImageID(coverImageID string) {
	l.UpdatedAt = time.Now()

	l.CoverImageID = coverImageID
}

func (l *List) ChangeTitle(title string) {
	l.UpdatedAt = time.Now()

	l.Title = title
}

func (l *List) ChangeDescription(description string) {
	l.UpdatedAt = time.Now()

	l.Description = description
}

func (l *List) AddMovies(movies []Movie) {
	l.UpdatedAt = time.Now()

	l.Movies = append(l.Movies, movies...)
}

func (l *List) RemoveMovies(moviesToRemove []Movie) {
	var updatedMovies []Movie

	for _, existingMovie := range l.Movies {
		found := false
		for _, movieToRemove := range moviesToRemove {
			if existingMovie.ID == movieToRemove.ID {
				found = true
				break
			}
		}

		if !found {
			l.UpdatedAt = time.Now()
			updatedMovies = append(updatedMovies, existingMovie)
		}
	}

	if len(updatedMovies) > 0 {
		l.UpdatedAt = time.Now()
		l.Movies = updatedMovies
	}
}

func (l *List) GetAvailableMoviesCombinations() [][]Movie {
	var combinations [][]Movie

	for i, movie1 := range l.Movies {
		for j, movie2 := range l.Movies {
			if i < j {
				combinations = append(combinations, []Movie{movie1, movie2})
			}
		}
	}

	return combinations
}

func (l *List) UpdateMovies(newMovies []Movie) ([]Movie, []Movie) {
	l.UpdatedAt = time.Now()

	currentMoviesMap := make(map[string]Movie)
	newMoviesMap := make(map[string]Movie)

	for _, movie := range l.Movies {
		currentMoviesMap[movie.ID] = movie
	}

	for _, newMovie := range newMovies {
		newMoviesMap[newMovie.ID] = newMovie
	}

	var moviesToDelete []Movie
	var moviesToAdd []Movie

	for _, movie := range l.Movies {
		if _, ok := newMoviesMap[movie.ID]; !ok {
			moviesToDelete = append(moviesToDelete, movie)
		}
	}

	for _, newMovie := range newMovies {
		if _, ok := currentMoviesMap[newMovie.ID]; !ok {
			moviesToAdd = append(moviesToAdd, newMovie)
		}
	}

	return moviesToDelete, moviesToAdd
}

func ValidateList(title, description, chooserID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if title == "" || len(title) > 100 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Título da lista inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O título da lista não pode estar vazio e deve ter no máximo 100 caracteres.",
			Instance: util.RFC400,
		})
	}

	if description == "" || len(description) > 150 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Descrição da lista inválida",
			Status:   http.StatusBadRequest,
			Detail:   "A descrição da lista não pode estar vazia e deve ter no máximo 150 caracteres.",
			Instance: util.RFC400,
		})
	}

	return validationErrors
}
