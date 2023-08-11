package password

import (
	bcrypt "golang.org/x/crypto/bcrypt"
)

type Password struct {
	secret []byte
}

func NewPassword() *Password {
	return &Password{}
}

func (p Password) GeneratePassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func (p Password) ComparePassword(hashPassword []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPassword, password)
	if err != nil {
		return false
	}

	return true
}
