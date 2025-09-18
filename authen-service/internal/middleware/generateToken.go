package middleware

import (
	"net/http"
)

func GenerateTokens(userId string) (string, int64, http.Cookie, error) {
	accessToken, expiresIn, err := AccessToken(userId)

	if err != nil {
		return "", expiresIn, http.Cookie{}, err
	}

	cookie, err := RefreshToken(userId)

	if err != nil {
		return "", expiresIn, http.Cookie{}, err
	}

	return accessToken, expiresIn, cookie, nil
}
