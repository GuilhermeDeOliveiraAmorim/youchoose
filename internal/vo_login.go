package internal

import (
	"context"
	"crypto/rand"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	EmailSalt    []byte
	PassSalt     []byte
}

func NewLogin(email, password string) (*Login, []ProblemDetails) {
	validationErrors, emailSalt, passSalt := ValidateLogin(email, password)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Login{
		Email:    email,
		Password: password,
		EmailSalt: emailSalt,
		PassSalt: passSalt,
	}, nil
}

func ValidateLogin(email, password string) ([]ProblemDetails, []byte, []byte) {
	var validationErrors []ProblemDetails

	if !isValidEmail(email) {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "E-mail inválido",
			Status: http.StatusBadRequest,
			Detail: "Por favor, forneça um endereço de e-mail válido.",
		})
	}

	if !isValidPassword(password) {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Senha inválida",
			Status: http.StatusBadRequest,
			Detail: "A senha deve ter pelo menos 6 caracteres, incluindo pelo menos uma letra maiúscula, uma letra minúscula, um numeral e um caracter especial.",
		})
	}

	emailSalt, err := generateSalt()
	if err != nil {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Erro ao gerar salt",
			Status: http.StatusBadRequest,
			Detail: "Ocorreu um erro ao gerar salt para o e-mail.",
		})
	}

	passSalt, err := generateSalt()
	if err != nil {
		validationErrors = append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Erro ao gerar salt",
			Status: http.StatusBadRequest,
			Detail: "Ocorreu um erro ao gerar salt para a senha.",
		})
	}

	return validationErrors, emailSalt, passSalt
}

func isValidEmail(email string) bool {
	emailPattern := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	match, _ := regexp.MatchString(emailPattern, email)
	return match
}

func isValidPassword(password string) bool {
	return hasMinimumLength(password, 6) &&
		hasUpperCaseLetter(password) &&
		hasLowerCaseLetter(password) &&
		hasDigit(password) &&
		hasSpecialCharacter(password)
}

func hasMinimumLength(password string, length int) bool {
	return len(password) >= length
}

func hasUpperCaseLetter(password string) bool {
	return strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func hasLowerCaseLetter(password string) bool {
	return strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
}

func hasDigit(password string) bool {
	return strings.ContainsAny(password, "0123456789")
}

func hasSpecialCharacter(password string) bool {
	specialCharacters := "@#$%&*"
	return strings.ContainsAny(password, specialCharacters)
}

func (l *Login) EncryptEmail(ctx context.Context) ([]byte, []ProblemDetails) {
	select {
	case <-ctx.Done():
		var validationErrors []ProblemDetails
		return nil, append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Erro ao encriptar e-mail",
			Status: http.StatusBadRequest,
			Detail: ctx.Err().Error(),
		})
	default:
	}

	return hashPassword(ctx, l.Email, l.EmailSalt)
}

func (l *Login) EncryptPassword(ctx context.Context) ([]byte, []ProblemDetails) {
	select {
	case <-ctx.Done():
		var validationErrors []ProblemDetails
		return nil, append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Erro ao encriptar password",
			Status: http.StatusBadRequest,
			Detail: ctx.Err().Error(),
		})
	default:
	}

	return hashPassword(ctx, l.Password, l.PassSalt)
}

func (l *Login) DecryptEmail(ctx context.Context, encryptedEmail []byte) (string, context.Context, []ProblemDetails) {
	select {
	case <-ctx.Done():
		var validationErrors []ProblemDetails
		return "", ctx, append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Erro ao decriptar e-mail",
			Status: http.StatusBadRequest,
			Detail: ctx.Err().Error(),
		})
	default:
	}

	return compareAndDecrypt(ctx, l.Email, encryptedEmail, l.EmailSalt)
}

func (l *Login) DecryptPassword(ctx context.Context, encryptedPassword []byte) (string, context.Context, []ProblemDetails) {
	select {
	case <-ctx.Done():
		var validationErrors []ProblemDetails
		return "", ctx, append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Erro ao decriptar password",
			Status: http.StatusBadRequest,
			Detail: ctx.Err().Error(),
		})
	default:
	}

	return compareAndDecrypt(ctx, l.Password, encryptedPassword, l.PassSalt)
}

func generateSalt() ([]byte, []ProblemDetails) {
	var validationErrors []ProblemDetails
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Erro ao gerar salt",
			Status: http.StatusBadRequest,
			Detail: "Ocorreu um erro ao gerar salt",
		})
	}
	return salt, nil
}

func hashPassword(_ context.Context, input string, salt []byte) ([]byte, []ProblemDetails) {
	var validationErrors []ProblemDetails
	hash, err := bcrypt.GenerateFromPassword([]byte(input+string(salt)), bcrypt.DefaultCost)
	if err != nil {
		return nil, append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Erro ao gerar hash",
			Status: http.StatusBadRequest,
			Detail: "Ocorreu um erro ao gerar hash para a senha.",
		})
	}
	return hash, nil
}

func compareAndDecrypt(ctx context.Context, input string, encrypted []byte, salt []byte) (string, context.Context, []ProblemDetails) {
	var validationErrors []ProblemDetails
	err := bcrypt.CompareHashAndPassword(encrypted, []byte(input+string(salt)))
	if err != nil {
		return "", ctx, append(validationErrors, ProblemDetails{
			Type:   "ValidationError",
			Title:  "Erro ao comparar e decriptar",
			Status: http.StatusBadRequest,
			Detail: "Ocorreu um erro ao comparar senha e hash para decriptar.",
		})
	}
	return input, ctx, nil
}