package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func AccessToken(userId string, secretKey []byte) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	accessString, err := accessToken.SignedString(secretKey)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	return accessString, nil
}
