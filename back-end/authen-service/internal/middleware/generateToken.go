package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GenerateTokens(userId string, c *gin.Context) (string, error) {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		err := godotenv.Load("../internal/.env")

		if err != nil {
			fmt.Println("Error loading file .env")
			return "", err
		}
	}

	accessToken, err := AccessToken(userId, []byte(os.Getenv("ACCESS_TOKEN_KEY")))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	cookie, err := RefreshToken(userId, c)

	if err != nil {
		return "", err
	}

	http.SetCookie(c.Writer, &cookie)

	return accessToken, nil
}
