package valueobject

import (
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

	if len(err) == 0 {
		t.Error("Erro esperado não foi retornado para e-mail inválido.")
	} else if len(err) != 1 || err[0].Status != http.StatusBadRequest {
		t.Errorf("Erro inesperado ao validar e-mail inválido: %v", err)
	}
}

func TestNewLoginValid(t *testing.T) {
	email := "user@example.com"
	password := "Abc123@"

	login, err := NewLogin(email, password)

	if len(err) > 0 {
		t.Errorf("Erro inesperado ao criar login válido: %v", err)
	}

	if login == nil {
		t.Error("Falha ao criar login válido. O objeto de login está nulo.")
	}
}

func TestNewLoginInvalidPassword(t *testing.T) {
	email := "user@example.com"
	password := "weak"

	_, err := NewLogin(email, password)

	if len(err) == 0 {
		t.Error("Login com senha fraca foi criado.")
	}

	if len(err) == 0 {
		t.Error("Erro esperado não foi retornado para senha fraca.")
	} else if len(err) != 1 || err[0].Status != http.StatusBadRequest {
		t.Errorf("Erro inesperado ao validar senha fraca: %v", err)
	}
}

func TestEncryptAndDecryptEmail(t *testing.T) {
	email := "user@example.com"
	password := "Abc123@"
	login, _ := NewLogin(email, password)

	encryptedEmail, err := login.EncryptEmail(email)
	if err != nil {
		t.Errorf("Erro ao encriptar e-mail: %v", err)
	}

	decryptedEmail := login.DecryptEmail(encryptedEmail, email)
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

	encryptedPassword, err := login.EncryptPassword(password)
	if err != nil {
		t.Errorf("Erro ao encriptar senha: %v", err)
	}

	decryptedPassword := login.DecryptPassword(encryptedPassword, password)
	if err != nil {
		t.Errorf("Erro ao decriptar senha: %v", err)
	}

	if decryptedPassword != password {
		t.Errorf("Senha decriptada diferente do esperado. Esperado: %s, Obtido: %s", password, decryptedPassword)
	}
}

func TestLogin_Equals(t *testing.T) {
	login1, _ := NewLogin("meuemail@bol.com.br", "As1@a7")
	login2, _ := NewLogin("meuemail@bol.com.br", "As1@a7")

	if got := login1.Equals(login2); got != true {
		t.Errorf("Login.Equals() = %v, want %v", got, login1.Equals(login2))
	}
}
