package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func RefreshToken(userId string) (http.Cookie, error) {
	expirationTime := time.Now().Add(168 * time.Hour)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   userId,
	})

	refreshString, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_KEY")))

	if err != nil {
		return http.Cookie{}, nil
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    refreshString,
		Expires:  expirationTime,
		HttpOnly: true, // If set true, the Front-End Js can not get the cookie
		Path:     "/auth",
		SameSite: 3,
	}

	return cookie, nil
}
