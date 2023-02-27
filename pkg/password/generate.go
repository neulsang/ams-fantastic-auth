package password

import (
	"golang.org/x/crypto/bcrypt"
)

func Generate(password string) (string, error) {
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return "", hashErr
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(password string, password1 string) error {
	compareErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(password1))
	if compareErr != nil {
		return compareErr
	}
	return nil
}
