package middleware

import (
	"errors"
	"fmt"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyAccessToken(tokenString string) (string, error) {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		return "", errors.New("error loading .env file")
	}

	mySigningKey := []byte(os.Getenv("ACCESS_TOKEN_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return "", err
	}

	return token.Claims.(jwt.MapClaims)["userID"].(string), nil
}

func VerifyRefreshToken(tokenString string) (string, error) {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		err := godotenv.Load("../internal/.env")

		if err != nil {
			fmt.Println("Error loading file .env")
			return "", err
		}
	}

	mySigningKey := []byte(os.Getenv("REFRESH_TOKEN_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return "", err
	}

	return token.Claims.(jwt.MapClaims)["userId"].(string), nil
}
