package usecase

import (
	"testing"
	"time"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	"youchoose/internal/util"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeactivateListUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockListRepositoryInterface(ctrl)
	deactivateListUseCase := NewDeactivateListUseCase(mockRepository)

	newList, _ := entity.NewList("Nome 1", "Descrição", uuid.NewString(), uuid.NewString(), uuid.NewString())
	time := time.Now()

	mockRepository.EXPECT().GetByID(newList.ID).Return(true, *newList, nil)

	mockRepository.EXPECT().Deactivate(gomock.Any()).Do(func(c *entity.List) {
		assert.False(t, c.Active)
		assert.NotEqual(t, c.DeactivatedAt, time)
		assert.NotEqual(t, c.UpdatedAt, time)
	}).Return(nil)

	input := DeactivateListInputDTO{ID: newList.ID}
	output, problemDetails := deactivateListUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, newList.ID, output.ID)
	assert.Equal(t, "Lista desativada com sucesso", output.Message)
	assert.True(t, output.IsSuccess)

	mockRepository.EXPECT().GetByID(newList.ID).Return(false, entity.List{}, nil)

	input = DeactivateListInputDTO{ID: newList.ID}
	output, problemDetails = deactivateListUseCase.Execute(input)

	assert.NotEmpty(t, problemDetails.ProblemDetails)
	assert.Equal(t, util.RFC404, problemDetails.ProblemDetails[0].Instance)
}
