package internal

import "net/http"

type Combination struct {
	VotationID    string `json:"votation_id"`
	FirstMovieID  string `json:"first_movie_id"`
	SecondMovieID string `json:"second_movie_id"`
	ChosenMovieID string `json:"chosen_movie_id"`
}

func NewCombination(votationID, firstMovieID, secondMovieID, chosenMovieID string) (*Combination, []ProblemDetails) {
	validationErrors := ValidateCombination(votationID, firstMovieID, secondMovieID, chosenMovieID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	combination := &Combination{
		VotationID:    votationID,
		FirstMovieID:  firstMovieID,
		SecondMovieID: secondMovieID,
		ChosenMovieID: chosenMovieID,
	}

	return combination, nil
}

func ValidateCombination(votationID, firstMovieID, secondMovieID, chosenMovieID string) []ProblemDetails {
	var validationErrors []ProblemDetails

	if votationID == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "ID da votação inválido",
			Status: http.StatusBadRequest,
			Detail: "O ID da votação não pode estar vazio.",
		})
	}

	if firstMovieID == "" || secondMovieID == "" || chosenMovieID == "" {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "IDs dos filmes na combinação inválidos",
			Status: http.StatusBadRequest,
			Detail: "Os IDs dos filmes na combinação não podem estar vazios.",
		})
	}

	return validationErrors
}