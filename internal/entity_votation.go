package internal

import (
	"time"
)

const (
	INPROGRESS = "In progress"
	FINISHED   = "Finished"
)

type Votation struct {
	SharedEntity
	ChooserID      string        `json:"chooser_id"`
	ListID         string        `json:"list_id"`
	Status         string        `json:"status"`
	Combinations   []Combination `json:"combinations"`
	StartTime      time.Time     `json:"start_time"`
	EndTime        time.Time     `json:"end_time"`
	ChosenCombination Combination   `json:"chosen_combination"`
}

func NewVotation(chooserID, listID string, combinations []Combination) *Votation {
	return &Votation{
		SharedEntity:   *NewSharedEntity(),
		ChooserID:      chooserID,
		ListID:         listID,
		Status:         INPROGRESS,
		Combinations:   combinations,
		StartTime:      time.Now(),
		EndTime:        time.Now(),
		ChosenCombination: Combination{},
	}
}

func (v *Votation) StartVotation() {
	v.StartTime = time.Now()
	v.Status = INPROGRESS
}

func (v *Votation) EndVotation() {
	v.EndTime = time.Now()
	v.Status = FINISHED
}

func (v *Votation) IsVotationInProgress() bool {
	return v.Status == INPROGRESS
}

func (v *Votation) Vote(combination Combination) {
	v.ChosenCombination = combination
}

func (v *Votation) GetAvailableCombinations() []Combination {
	return v.Combinations
}

func (v *Votation) GetVotedCombinations() []Combination {
	var votedCombinations []Combination

	for _, combination := range v.Combinations {
		if combination.ChosenMovieID != "" {
			votedCombinations = append(votedCombinations, combination)
		}
	}

	return votedCombinations
}

func (v *Votation) GetUnvotedCombinations() []Combination {
	var unvotedCombinations []Combination

	for _, combination := range v.Combinations {
		if combination.ChosenMovieID == "" {
			unvotedCombinations = append(unvotedCombinations, combination)
		}
	}

	return unvotedCombinations
}