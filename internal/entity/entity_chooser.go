package entity

import (
	"context"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/util"
	valueobject "github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/value_object"
)

type Chooser struct {
	SharedEntity
	Name      string     `json:"name"`
	Login     *valueobject.Login     `json:"login"`
	Address   *valueobject.Address   `json:"address"`
	BirthDate *valueobject.BirthDate `json:"birth_date"`
	ImageID   string     `json:"image_id"`
}

func NewChooser(name string, login *valueobject.Login, address *valueobject.Address, birthDate *valueobject.BirthDate, imageID string) (*Chooser, []util.ProblemDetails) {
	validationErrors := ValidateChooser(name, login, address, birthDate, imageID)

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

func ValidateChooser(name string, login *valueobject.Login, address *valueobject.Address, birthDate *valueobject.BirthDate, imageID string) []util.ProblemDetails {
	var validationErrors []util.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "Nome do Chooser inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do Chooser não pode estar vazio.",
			Instance: util.RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "Nome do Chooser inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O nome do Chooser não pode ter mais do que 100 caracteres.",
			Instance: util.RFC400,
		})
	}

	if login == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "Login do Chooser inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O login do Chooser não pode ser nulo.",
			Instance: util.RFC400,
		})
	}

	if address == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "Endereço do Chooser inválido",
			Status:   http.StatusBadRequest,
			Detail:   "O endereço do Chooser não pode ser nulo.",
			Instance: util.RFC400,
		})
	}

	if birthDate == nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "Data de nascimento do Chooser inválida",
			Status:   http.StatusBadRequest,
			Detail:   "A data de nascimento do Chooser não pode ser nula.",
			Instance: util.RFC400,
		})
	}

	if imageID == "" {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
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
			Type:     "ValidationError",
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
			Type:     "ValidationError",
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

func (c *Chooser) ChangeImageID(ctx context.Context, newImageID string) []util.ProblemDetails {
	select {
	case <-ctx.Done():
		var validationErrors []util.ProblemDetails
		return append(validationErrors, util.ProblemDetails{
			Type:     "ValidationError",
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
