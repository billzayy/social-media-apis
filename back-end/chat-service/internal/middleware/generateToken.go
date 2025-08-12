package middleware

import (
	"net/http"
)

func GenerateTokens(userId string) (string, http.Cookie, error) {
	accessToken, err := AccessToken(userId)

	if err != nil {
		return "", http.Cookie{}, err
	}

	cookie, err := RefreshToken(userId)

	if err != nil {
		return "", http.Cookie{}, err
	}

	return accessToken, cookie, nil
}
