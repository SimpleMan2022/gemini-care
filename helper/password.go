package helper

import "golang.org/x/crypto/bcrypt"

type PasswordHelper interface {
	HashPassword(password string) (string, error)
}
type passwordHelper struct{}

func NewPasswordHelper() *passwordHelper {
	return &passwordHelper{}
}

func (ph *passwordHelper) HashPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
}
