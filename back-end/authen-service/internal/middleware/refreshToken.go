package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func RefreshToken(userId string, c *gin.Context) (http.Cookie, error) {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		err := godotenv.Load("../internal/.env")

		if err != nil {
			fmt.Println("Error loading file .env")
			return http.Cookie{}, err
		}
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		// "exp":    time.Now().AddDate(0, 1, 0).Unix(),
		"exp": time.Now().Add(time.Hour * 168).Unix(), // Exp 7 days
	})

	refreshString, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_KEY")))

	if err != nil {
		fmt.Println(err)
		return http.Cookie{}, nil
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    refreshString,
		HttpOnly: true, // If set true, the Front-End Js can not get the cookie
		Path:     "/auth",
		SameSite: 3,
	}

	return cookie, nil
}
