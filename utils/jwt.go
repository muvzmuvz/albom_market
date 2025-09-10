package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtSecret = []byte("8b6b9104-599d-4ef9-9807-23c5e62e77e0") // В продакшн хранить в env

func GenerateJWT(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // токен на 72 часа
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
