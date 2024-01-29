package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "scretkey"

func GenerateToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(t string) (int64, error) {
	parsedToken, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return -1, errors.New("Could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return -1, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("Invalid claims")
	}

	// email := claims["email"].(string)
	id := claims["id"].(float64)

	return int64(id), nil

}
