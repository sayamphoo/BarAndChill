package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(str string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func CheckHash(hash, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return (err == nil)
}
