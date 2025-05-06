package middleware

import (
	"errors"
	"fmt"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string, key string) (string, error) {
	secretKey := []byte(os.Getenv(key)) // Load once at startup ideally

	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token is signed with HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid or expired token")
	}

	return claims.Subject, nil
}
