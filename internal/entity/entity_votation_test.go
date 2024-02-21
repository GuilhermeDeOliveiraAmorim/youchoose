package entity

import (
	"net/http"
	"testing"
	"youchoose/internal/util"

	"github.com/stretchr/testify/assert"
)

func TestNewVotation_ValidInput(t *testing.T) {
	chooserID := "chooser_id"
	listID := "list_id"
	firstMovieID := "first_movie_id"
	secondMovieID := "second_movie_id"
	chosenMovieID := "chosen_movie_id"

	votation, validationErrors := NewVotation(chooserID, listID, firstMovieID, secondMovieID, chosenMovieID)

	assert.Nil(t, validationErrors)
	assert.NotNil(t, votation)
	assert.Equal(t, chooserID, votation.ChooserID)
	assert.Equal(t, listID, votation.ListID)
	assert.Equal(t, firstMovieID, votation.FirstMovieID)
	assert.Equal(t, secondMovieID, votation.SecondMovieID)
	assert.Equal(t, chosenMovieID, votation.ChosenMovieID)
}

func TestNewVotation_InvalidInput(t *testing.T) {
	chooserID := ""
	listID := "list_id"
	firstMovieID := "first_movie_id"
	secondMovieID := "second_movie_id"
	chosenMovieID := "chosen_movie_id"

	votation, validationErrors := NewVotation(chooserID, listID, firstMovieID, secondMovieID, chosenMovieID)

	assert.Nil(t, votation)
	assert.NotNil(t, validationErrors)
	assert.Len(t, validationErrors, 1)
	assert.Equal(t, "Validation Error", validationErrors[0].Type)
	assert.Equal(t, "Existe um ou mais IDs inválidos", validationErrors[0].Title)
	assert.Equal(t, http.StatusBadRequest, validationErrors[0].Status)
	assert.Equal(t, "Para registrar uma votação os IDs não podem estar vazios.", validationErrors[0].Detail)
	assert.Equal(t, util.RFC400, validationErrors[0].Instance)
}
