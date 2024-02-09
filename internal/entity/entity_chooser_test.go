package entity

import (
	"context"
	"testing"

	valueobject "youchoose/internal/value_object"

	"github.com/google/uuid"
)

func TestNewChooser(t *testing.T) {
	name := "John Doe"
	login := &valueobject.Login{Email: "john@example.com", Password: "P@ssw0rd"}
	address := &valueobject.Address{City: "City", State: "State", Country: "Country"}
	birthDate := &valueobject.BirthDate{Day: 1, Month: 1, Year: 2000}
	imageID := uuid.New().String()

	chooser, err := NewChooser(name, login, address, birthDate, imageID)

	if err != nil {
		t.Errorf("Unexpected error creating Chooser: %v", err)
	}

	if chooser == nil {
		t.Error("Chooser should not be nil")
	}
}

func TestChangeLogin(t *testing.T) {
	oldLogin := &valueobject.Login{Email: "old@example.com", Password: "OldP@ssw0rd"}
	newLogin := &valueobject.Login{Email: "new@example.com", Password: "NewP@ssw0rd"}
	chooser := &Chooser{Login: oldLogin}

	err := chooser.ChangeLogin(context.Background(), newLogin)

	if err != nil {
		t.Errorf("Unexpected error changing login: %v", err)
	}

	if chooser.Login != newLogin {
		t.Error("Login should be updated")
	}
}

func TestChangeImageID(t *testing.T) {
	oldImageID := "1ad79480-d10b-4d3e-b01e-13ede0c815bc"
	newImageID := "a8eeacd9-3e0d-43a5-8422-dda0c6b46aa6"
	chooser := &Chooser{ImageID: oldImageID}

	err := chooser.ChangeImageID(context.Background(), newImageID)

	if err != nil {
		t.Errorf("Unexpected error changing image id: %v", err)
	}

	if chooser.ImageID != newImageID {
		t.Error("Image id should be updated")
	}
}
