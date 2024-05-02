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

func TestDeactivateChooserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockChooserRepositoryInterface(ctrl)
	deactivateChooserUseCase := NewDeactivateChooserUseCase(mockRepository)

	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, _ := entity.NewChooser(name, login, address, birthDate, imageID)
	time := time.Now()

	chooserID := chooser.ID
	input := DeactivateChooserInputDTO{ChooserID: chooserID}

	mockRepository.EXPECT().GetByID(gomock.Any()).Return(true, *chooser, nil).Times(2)

	chooser.Deactivate()

	mockRepository.EXPECT().Deactivate(gomock.Any()).Do(func(c *entity.Chooser) {
		assert.False(t, c.Active)
		assert.NotEqual(t, c.DeactivatedAt, time)
		assert.NotEqual(t, c.UpdatedAt, time)
	}).Return(nil)

	output, problemDetails := deactivateChooserUseCase.Execute(input)

	assert.Empty(t, problemDetails.ProblemDetails)
	assert.Equal(t, chooserID, output.ChooserID)
	assert.Equal(t, "Chooser desativado com sucesso", output.Message)
	assert.True(t, output.IsSuccess)
}
