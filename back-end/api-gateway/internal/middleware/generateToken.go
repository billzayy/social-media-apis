package middleware

import (
	"fmt"
	"net/http"
	"os"
)

func GenerateTokens(userId string) (string, http.Cookie, error) {
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
