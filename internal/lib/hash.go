package lib

import (
	"golang.org/x/crypto/bcrypt"
)

func MakeHash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	return string(bytes), err
}

func CheckHash(s string, h string) bool {
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(s)) == nil
}
