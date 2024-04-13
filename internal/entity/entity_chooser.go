package entity

import (
	"context"
	"net/http"

	"youchoose/internal/util"

	valueobject "youchoose/internal/value_object"

	"github.com/google/uuid"
)

type Chooser struct {
	SharedEntity
	Name      string                 `json:"name"`
	Login     *valueobject.Login     `json:"login"`
	Address   *valueobject.Address   `json:"address"`
	BirthDate *valueobject.BirthDate `json:"birth_date"`
	ImageID   string                 `json:"image_id"`
}

func NewChooser(name string, login *valueobject.Login, address *valueobject.Address, birthDate *valueobject.BirthDate, imageID string) (*Chooser, []util.ProblemDetails) {
	validationErrors := ValidateChooser(name, imageID)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	chooser := &Chooser{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		Login:        login,
		Address:      address,
		BirthDate:    birthDate,
		ImageID:      imageID,
	}

	return chooser, nil
}

func ValidateChooser(name string, imageID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Nome do Chooser inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do Chooser não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Nome do Chooser inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do Chooser não pode ter mais do que 100 caracteres.",
			Instance: util.RFC400,
		})
	}

	if uuid.Validate(imageID) != nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "ID de imagem do Chooser inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O ID de imagem do Chooser não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	return validationErrors
}

func (c *Chooser) ChangeLogin(ctx context.Context, newLogin *valueobject.Login) []util.ProblemDetails {
	select {
	case <-ctx.Done():
		var validationErrors []util.ProblemDetails
		return append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Erro ao alterar o login do Chooser",
			Status:   http.StatusBadRequest,
			Detail:   ctx.Err().Error(),
			Instance: util.RFC400,
		})
	default:
	}

	c.Login = newLogin
	return nil
}

func (c *Chooser) ChangeAddress(ctx context.Context, newAddress *valueobject.Address) []util.ProblemDetails {
	select {
	case <-ctx.Done():
		var validationErrors []util.ProblemDetails
		return append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Erro ao alterar o endereço do Chooser",
			Status:   http.StatusBadRequest,
			Detail:   ctx.Err().Error(),
			Instance: util.RFC400,
		})
	default:
	}

	c.Address = newAddress
	return nil
}

func (c *Chooser) ChangeBirthDate(ctx context.Context, newBirthDate *valueobject.BirthDate) []util.ProblemDetails {
	select {
	case <-ctx.Done():
		var validationErrors []util.ProblemDetails
		return append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Erro ao alterar a data de aniversário do Chooser",
			Status:   http.StatusBadRequest,
			Detail:   ctx.Err().Error(),
			Instance: util.RFC400,
		})
	default:
	}

	c.BirthDate = newBirthDate
	return nil
}

func (c *Chooser) ChangeImageID(ctx context.Context, newImageID string) []util.ProblemDetails {
	select {
	case <-ctx.Done():
		var validationErrors []util.ProblemDetails
		return append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Erro ao alterar o ID de imagem do Chooser",
			Status:   http.StatusBadRequest,
			Detail:   ctx.Err().Error(),
			Instance: util.RFC400,
		})
	default:
	}

	c.ImageID = newImageID
	return nil
}

func (c *Chooser) ChangeName(ctx context.Context, newName string) []util.ProblemDetails {
	select {
	case <-ctx.Done():
		var validationErrors []util.ProblemDetails
		return append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    "Erro ao alterar o nome do Chooser",
			Status:   http.StatusBadRequest,
			Detail:   ctx.Err().Error(),
			Instance: util.RFC400,
		})
	default:
	}

	c.Name = newName
	return nil
}
