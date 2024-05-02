package usecase

import (
	"errors"
	"net/http"
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestResetVoteUseCase_VotationNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockVotationRepo := mock.NewMockVotationRepositoryInterface(ctrl)

	resetVoteUC := NewResetVoteUseCase(mockChooserRepo, mockVotationRepo)

	login1, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address1, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")

	birthDate1, _ := valueobject.NewBirthDate(1, 1, 1990)

	newChooser, _ := entity.NewChooser("Nome 1", login1, address1, birthDate1, uuid.New().String())

	input := ResetVoteInputDTO{
		ChooserID:  newChooser.ID,
		VotationID: "non_existing_votation_id",
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *newChooser, nil)
	mockVotationRepo.EXPECT().GetByID(input.VotationID).Return(false, entity.Votation{}, nil)

	output, problemDetails := resetVoteUC.Execute(input)

	assert.Empty(t, output.ID)
	assert.Equal(t, "Voto não encontrado", problemDetails.ProblemDetails[0].Title)
	assert.Equal(t, http.StatusNotFound, problemDetails.ProblemDetails[0].Status)
}

func TestResetVoteUseCase_UpdateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockVotationRepo := mock.NewMockVotationRepositoryInterface(ctrl)

	resetVoteUC := NewResetVoteUseCase(mockChooserRepo, mockVotationRepo)

	login1, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address1, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")

	birthDate1, _ := valueobject.NewBirthDate(1, 1, 1990)

	newChooser, _ := entity.NewChooser("Nome 1", login1, address1, birthDate1, uuid.New().String())

	input := ResetVoteInputDTO{
		ChooserID:  newChooser.ID,
		VotationID: "votation_id",
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *newChooser, nil)
	mockVotationRepo.EXPECT().GetByID(input.VotationID).Return(true, entity.Votation{}, nil)
	mockVotationRepo.EXPECT().Update(gomock.Any()).Return(errors.New("update error"))

	output, problemDetails := resetVoteUC.Execute(input)

	assert.Empty(t, output.ID)
	assert.Equal(t, "Erro ao cancelar um voto", problemDetails.ProblemDetails[0].Title)
	assert.Equal(t, http.StatusInternalServerError, problemDetails.ProblemDetails[0].Status)
}

func TestResetVoteUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockVotationRepo := mock.NewMockVotationRepositoryInterface(ctrl)

	resetVoteUC := NewResetVoteUseCase(mockChooserRepo, mockVotationRepo)

	login1, _ := valueobject.NewLogin("email@email.com", "12@#asd89")

	address1, _ := valueobject.NewAddress("Aracaju", "Sergipe", "Brasil")

	birthDate1, _ := valueobject.NewBirthDate(1, 1, 1990)

	newChooser, _ := entity.NewChooser("Nome 1", login1, address1, birthDate1, uuid.New().String())

	newVotation, _ := entity.NewVotation(uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String())

	input := ResetVoteInputDTO{
		ChooserID:  newChooser.ID,
		VotationID: newVotation.ID,
	}

	mockChooserRepo.EXPECT().GetByID(input.ChooserID).Return(true, *newChooser, nil)
	mockVotationRepo.EXPECT().GetByID(input.VotationID).Return(true, *newVotation, nil)
	mockVotationRepo.EXPECT().Update(gomock.Any()).Return(nil)

	output, problemDetails := resetVoteUC.Execute(input)

	assert.Equal(t, newVotation.ID, output.ID)
	assert.Equal(t, "Voto cancelado com sucesso", output.Message)
	assert.True(t, output.IsSuccess)
	assert.Empty(t, problemDetails.ProblemDetails)
}
