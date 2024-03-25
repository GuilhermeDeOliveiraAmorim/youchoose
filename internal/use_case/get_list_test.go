package usecase

import (
	"testing"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetListUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockListRepositoryInterface(ctrl)
	getListUseCase := NewGetListUseCase(mockRepository)

	list1, _ := entity.NewList("Minha Lista 1", "Descrição da Lista 1", uuid.NewString(), uuid.NewString(), uuid.NewString())

	input := GetListInputDTO{
		ID: list1.ID,
	}

	mockRepository.EXPECT().GetByID(input.ID).Return(true, *list1, nil)

	output, problemDetails := getListUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, output.ID, list1.ID)
}
