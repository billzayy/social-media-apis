package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GenerateTokens(userId string) (string, http.Cookie, error) {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		err := godotenv.Load("../internal/.env")

		if err != nil {
			fmt.Println("Error loading file .env")
			return "", http.Cookie{}, err
		}
	}

	accessToken, err := AccessToken(userId, []byte(os.Getenv("ACCESS_TOKEN_KEY")))

	if err != nil {
		fmt.Println(err)
		return "", http.Cookie{}, err
	}

	cookie, err := RefreshToken(userId)

	if err != nil {
		return "", http.Cookie{}, err
	}

	return accessToken, cookie, nil
}
