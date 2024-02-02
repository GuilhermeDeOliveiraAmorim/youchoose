package valueobject

import (
	"context"
	"net/http"
	"testing"
)

func TestNewLoginInvalidEmail(t *testing.T) {
	email := "invalid-email"
	password := "Abc123@"

	login, err := NewLogin(email, password)

	if login != nil {
		t.Error("Login com e-mail inválido foi criado.")
	}

	if err == nil {
		t.Error("Erro esperado não foi retornado para e-mail inválido.")
	} else if len(err) != 1 || err[0].Status != http.StatusBadRequest {
		t.Errorf("Erro inesperado ao validar e-mail inválido: %v", err)
	}
}

func TestNewLoginValid(t *testing.T) {
	email := "user@example.com"
	password := "Abc123@"

	login, err := NewLogin(email, password)

	if err != nil {
		t.Errorf("Erro inesperado ao criar login válido: %v", err)
	}

	if login == nil {
		t.Error("Falha ao criar login válido. O objeto de login está nulo.")
	}
}

func TestNewLoginInvalidPassword(t *testing.T) {
	email := "user@example.com"
	password := "weak"

	login, err := NewLogin(email, password)

	if login != nil {
		t.Error("Login com senha fraca foi criado.")
	}

	if err == nil {
		t.Error("Erro esperado não foi retornado para senha fraca.")
	} else if len(err) != 1 || err[0].Status != http.StatusBadRequest {
		t.Errorf("Erro inesperado ao validar senha fraca: %v", err)
	}
}

func TestEncryptAndDecryptEmail(t *testing.T) {
	email := "user@example.com"
	password := "Abc123@"
	login, _ := NewLogin(email, password)

	encryptedEmail, err := login.EncryptEmail(context.Background())
	if err != nil {
		t.Errorf("Erro ao encriptar e-mail: %v", err)
	}

	decryptedEmail, _, err := login.DecryptEmail(context.Background(), encryptedEmail)
	if err != nil {
		t.Errorf("Erro ao decriptar e-mail: %v", err)
	}

	if decryptedEmail != email {
		t.Errorf("Email decriptado diferente do esperado. Esperado: %s, Obtido: %s", email, decryptedEmail)
	}
}

func TestEncryptAndDecryptPassword(t *testing.T) {
	email := "user@example.com"
	password := "Abc123@"
	login, _ := NewLogin(email, password)

	encryptedPassword, err := login.EncryptPassword(context.Background())
	if err != nil {
		t.Errorf("Erro ao encriptar senha: %v", err)
	}

	decryptedPassword, _, err := login.DecryptPassword(context.Background(), encryptedPassword)
	if err != nil {
		t.Errorf("Erro ao decriptar senha: %v", err)
	}

	if decryptedPassword != password {
		t.Errorf("Senha decriptada diferente do esperado. Esperado: %s, Obtido: %s", password, decryptedPassword)
	}
}