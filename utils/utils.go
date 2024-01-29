package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) (string, error) {
	b, e := bcrypt.GenerateFromPassword([]byte(p), 14)
	return string(b), e
}

func CheckPasswordHash(p, hashedP string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedP), []byte(p))
	return err == nil
}
