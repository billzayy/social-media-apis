package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func AccessToken(userId string) (string, error) {
	secretKey := []byte(os.Getenv("ACCESS_TOKEN_KEY"))

	//? Set token expiration
	expirationTime := time.Now().Add(time.Hour * 24)

	//? Create the JWT claims
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   userId,
	}

	//? Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessString, err := token.SignedString(secretKey)

	if err != nil {
		return "", nil
	}

	return accessString, nil
}
