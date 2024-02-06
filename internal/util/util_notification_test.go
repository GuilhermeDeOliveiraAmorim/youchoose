package util

import (
	"testing"
)

func TestNewNotificationValid(t *testing.T) {
	notification, err := NewNotification("info", "Operação concluída com sucesso")
	if err != nil {
		t.Errorf("Erro inesperado ao criar a notificação: %v", err)
	}

	if notification.Key != "info" {
		t.Errorf("Esperava-se que a chave fosse 'info', mas é '%s'", notification.Key)
	}
	if notification.Value != "Operação concluída com sucesso" {
		t.Errorf("Esperava-se que o valor fosse 'Operação concluída com sucesso', mas é '%s'", notification.Value)
	}
}

func TestNewNotificationInvalidKey(t *testing.T) {
	_, err := NewNotification("", "Operação concluída com sucesso")
	if err == nil {
		t.Error("Esperava-se um erro ao criar a notificação com chave vazia, mas nenhum erro foi retornado")
	} else {
		expectedErrorMsg := "a chave da notificação deve ter entre 1 e 50 caracteres"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Esperava-se o erro '%s', mas o erro foi '%s'", expectedErrorMsg, err.Error())
		}
	}
}

func TestNewNotificationInvalidValue(t *testing.T) {
	_, err := NewNotification("info", "Lorem ipsum, consectetur adipiscing consectetur adipiscing elit dolor sit amet, elit dolor sit amet, consectetur adipiscing elit dolor sit amet, consectetur adipiscing elit. Sed do eiusmod, consectetur adipiscing elit tempor incididunt ut labore et dolore magna aliqua.")
	if err == nil {
		t.Error("Esperava-se um erro ao criar a notificação com valor muito longo, mas nenhum erro foi retornado")
	} else {
		expectedErrorMsg := "o valor da notificação não pode ter mais do que 255 caracteres"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Esperava-se o erro '%s', mas o erro foi '%s'", expectedErrorMsg, err.Error())
		}
	}
}
