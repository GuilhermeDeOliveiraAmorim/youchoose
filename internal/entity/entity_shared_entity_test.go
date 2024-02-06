package entity

import (
	"testing"

	"github.com/GuilhermeDeOliveiraAmorim/youchoose/internal/util"
)

func TestNewSharedEntity(t *testing.T) {
	sharedEntity := NewSharedEntity()

	if sharedEntity.ID == "" {
		t.Error("O ID da entidade não foi inicializado corretamente")
	}
	if !sharedEntity.Active {
		t.Error("A entidade deveria ser ativa por padrão")
	}
	if sharedEntity.CreatedAt.IsZero() {
		t.Error("O campo CreatedAt não foi inicializado corretamente")
	}
	if sharedEntity.UpdatedAt.IsZero() {
		t.Error("O campo UpdatedAt não foi inicializado corretamente")
	}
	if sharedEntity.DeactivatedAt.IsZero() {
		t.Error("O campo DeactivatedAt não foi inicializado corretamente")
	}
}

func TestActivate(t *testing.T) {
	sharedEntity := NewSharedEntity()
	sharedEntity.Activate()

	if !sharedEntity.Active {
		t.Error("A entidade deveria estar ativa após a ativação")
	}
	if sharedEntity.UpdatedAt.IsZero() {
		t.Error("O campo UpdatedAt não foi atualizado corretamente durante a ativação")
	}
}

func TestDeactivate(t *testing.T) {
	sharedEntity := NewSharedEntity()
	sharedEntity.Deactivate()

	if sharedEntity.Active {
		t.Error("A entidade deveria estar inativa após a desativação")
	}
	if sharedEntity.UpdatedAt.IsZero() {
		t.Error("O campo UpdatedAt não foi atualizado corretamente durante a desativação")
	}
	if sharedEntity.DeactivatedAt.IsZero() {
		t.Error("O campo DeactivatedAt não foi atualizado corretamente durante a desativação")
	}
}

func TestAddNotification(t *testing.T) {
	sharedEntity := NewSharedEntity()
	sharedEntity.AddNotification("info", "Operação concluída com sucesso")

	if len(sharedEntity.Notifications) != 1 {
		t.Error("A notificação não foi adicionada corretamente")
	}
}

func TestAddError(t *testing.T) {
	sharedEntity := NewSharedEntity()
	sharedEntity.AddError(util.ProblemDetails{
		Type:   "ValidationError",
		Title:  "Erro de validação",
		Status: 400,
		Detail: "Campos obrigatórios não preenchidos",
	})

	if len(sharedEntity.Errors) != 1 {
		t.Error("O erro não foi adicionado corretamente")
	}
}

func TestClearNotifications(t *testing.T) {
	sharedEntity := NewSharedEntity()
	sharedEntity.AddNotification("info", "Operação concluída com sucesso")
	sharedEntity.ClearNotifications()

	if len(sharedEntity.Notifications) != 0 {
		t.Error("As notificações não foram removidas corretamente")
	}
}

func TestClearErrors(t *testing.T) {
	sharedEntity := NewSharedEntity()
	sharedEntity.AddError(util.ProblemDetails{
		Type:   "ValidationError",
		Title:  "Erro de validação",
		Status: 400,
		Detail: "Campos obrigatórios não preenchidos",
	})
	sharedEntity.ClearErrors()

	if len(sharedEntity.Errors) != 0 {
		t.Error("Os erros não foram removidos corretamente")
	}
}
