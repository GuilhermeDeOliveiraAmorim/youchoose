package valueobject

import (
	"net/http"
	"regexp"
	"strings"

	"youchoose/internal/util"

	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLogin(email, password string) (*Login, []util.ProblemDetails) {
	validationErrors := ValidateLogin(email, password)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Login{
		Email:    email,
		Password: password,
	}, nil
}

func ValidateLogin(email, password string) []util.ProblemDetails {
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

	return validationErrors
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

func hashString(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareAndDecrypt(hashedData string, data string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(data))
	return err == nil
}

func (lo *Login) EncryptEmail(email string) (string, error) {
	hashedEmail, err := hashString(email)
	if err != nil {
		return "", err
	}
	return hashedEmail, nil
}

func (lo *Login) EncryptPassword(password string) (string, error) {
	hashedPassword, err := hashString(password)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}

func (lo *Login) DecryptEmail(hashedEmail string, email string) string {
	if compareAndDecrypt(hashedEmail, email) {
		return email
	} else {
		return ""
	}
}

func (lo *Login) DecryptPassword(hashedPassword string, password string) string {
	if compareAndDecrypt(hashedPassword, password) {
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
