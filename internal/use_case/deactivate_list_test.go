package usecase

import (
	"testing"
	"time"
	"youchoose/internal/entity"
	"youchoose/internal/use_case/mock"
	valueobject "youchoose/internal/value_object"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeactivateListUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChooserRepo := mock.NewMockChooserRepositoryInterface(ctrl)
	mockListRepo := mock.NewMockListRepositoryInterface(ctrl)
	deactivateListUseCase := NewDeactivateListUseCase(mockChooserRepo, mockListRepo)

	list, _ := entity.NewList("Nome 1", "Descrição", uuid.NewString(), uuid.NewString(), uuid.NewString())
	time := time.Now()

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.NewString()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)

	mockChooserRepo.EXPECT().GetByID(chooser.ID).Return(true, *chooser, nil)
	mockListRepo.EXPECT().GetByID(list.ID).Return(true, *list, nil)

	mockListRepo.EXPECT().Deactivate(gomock.Any()).Do(func(c *entity.List) {
		assert.False(t, c.Active)
		assert.NotEqual(t, c.DeactivatedAt, time)
		assert.NotEqual(t, c.UpdatedAt, time)
	}).Return(nil)

	input := DeactivateListInputDTO{
		ChooserID: chooser.ID,
		ListID:    list.ID,
	}

	output, problemDetails := deactivateListUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, list.ID, output.ListID)
	assert.Equal(t, "Lista desativada com sucesso", output.Message)
	assert.True(t, output.IsSuccess)
}
