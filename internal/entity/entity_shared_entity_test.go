package entity

import (
	"testing"
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
