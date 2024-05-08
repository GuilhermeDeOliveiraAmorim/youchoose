package valueobject

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"regexp"
	"strings"

	"youchoose/internal/util"

	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	EmailSalt string
	PasswordSalt  string
}

func NewLogin(email, password string) (*Login, []util.ProblemDetails) {
	validationErrors, emailSalt, passSalt := ValidateLogin(email, password)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Login{
		Email:     email,
		Password:  password,
		EmailSalt: emailSalt,
		PasswordSalt:  passSalt,
	}, nil
}

func ValidateLogin(email, password string) ([]util.ProblemDetails, string, string) {
	var validationErrors []util.ProblemDetails

	if !isValidEmail(email) {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.LoginErrorTitleInvalidEmail,
			Status:   http.StatusBadRequest,
			Detail:   util.LoginErrorDetailInvalidEmail,
			Instance: util.RFC400,
		})
	}

	if !isValidPassword(password) {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.LoginErrorTitleInvalidPassword,
			Status:   http.StatusBadRequest,
			Detail:   util.LoginErrorDetailInvalidPassword,
			Instance: util.RFC400,
		})
	}

	emailSalt, err := generateSalt()
	if err != nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.LoginErrorTitleSaltGeneration,
			Status:   http.StatusBadRequest,
			Detail:   util.LoginErrorDetailSaltGeneration,
			Instance: util.RFC400,
		})
	}

	passSalt, err := generateSalt()
	if err != nil {
		validationErrors = append(validationErrors, util.ProblemDetails{
			Type:     util.TypeValidationError,
			Title:    util.LoginErrorTitleSaltGeneration,
			Status:   http.StatusBadRequest,
			Detail:   util.LoginErrorDetailSaltGeneration,
			Instance: util.RFC400,
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

func generateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func hashString(data string, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (lo *Login) EncryptEmail(email string) (string, string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", "", err
	}
	hashedEmail, err := hashString(email, salt)
	if err != nil {
		return "", "", err
	}
	return hashedEmail, salt, nil
}

func (lo *Login) EncryptPassword(password string) (string, string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", "", err
	}
	hashedPassword, err := hashString(password, salt)
	if err != nil {
		return "", "", err
	}
	return hashedPassword, salt, nil
}
func compareAndDecrypt(data string, hashedData string, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(data+salt))
	return err == nil
}

func (lo *Login) DecryptEmail(email string, hashedEmail string, salt string) string {
	if compareAndDecrypt(email, hashedEmail, salt) {
		return email
	} else {
		return ""
	}
}

func (lo *Login) DecryptPassword(password string, hashedPassword string, salt string) string {
	if compareAndDecrypt(password, hashedPassword, salt) {
		return password
	} else {
		return ""
	}
}

func (lo *Login) ChangeEmail(newEmail string) {
	lo.Email = newEmail
}

func (lo *Login) ChangePassword(newPassword string) {
	lo.Password = newPassword
}

func (lo *Login) Equals(other *Login) bool {
	return lo.Email == other.Email && lo.Password == other.Password
}
