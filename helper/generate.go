package helper

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateNewUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}

func GenerateHashPassword(password string) string {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	PanicIfError(err)
	return string(res)
}
