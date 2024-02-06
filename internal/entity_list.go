package internal

import (
	"net/http"
)

type List struct {
	SharedEntity
	ProfileImageID          string       `json:"profile_image_id"`
	CoverImageID            string       `json:"cover_image_id"`
	Title                   string       `json:"title"`
	Description             string       `json:"description"`
	Movies                  []Movie      `json:"movies"`
	ChooserID               string       `json:"chooser_id"`
	Votes                   int          `json:"votes"`
}

func NewList(profileImageID, coverImageID, title, description, chooserID string) (*List, []ProblemDetails) {
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
	l.Votes++
}

func (l *List) ChangeProfileImageID(profileImageID string) {
	l.ProfileImageID = profileImageID
}

func (l *List) ChangeCoverImageID(coverImageID string) {
	l.CoverImageID = coverImageID
}

func (l *List) ChangeTitle(title string) {
	l.Title = title
}

func (l *List) AddMovies(movies []Movie) {
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
			updatedMovies = append(updatedMovies, existingMovie)
		}
	}
	
	if len(updatedMovies) > 0 {
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

func ValidateList(title, description, chooserID string) []ProblemDetails {
	var validationErrors []ProblemDetails

	if title == "" || len(title) > 100 {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Título da lista inválido",
			Status: http.StatusBadRequest,
			Detail: "O título da lista não pode estar vazio e deve ter no máximo 100 caracteres.",
			Instance: RFC400,
		})
	}

	if description == "" || len(description) > 150 {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Descrição da lista inválida",
			Status: http.StatusBadRequest,
			Detail: "A descrição da lista não pode estar vazia e deve ter no máximo 150 caracteres.",
			Instance: RFC400,
		})
	}

	return validationErrors
}