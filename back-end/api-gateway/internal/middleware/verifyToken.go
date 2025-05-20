package middleware

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyAccessToken(tokenString string) (string, error) {
	// Get secret key from environment variable
	secret := os.Getenv("ACCESS_TOKEN_KEY")

	if secret == "" {
		return "", fmt.Errorf("secret key not set in environment: %s", "ACCESS_TOKEN_KEY")
	}

	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	// Extract claims safely
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("sub claim missing or invalid")
	}

	return userID, nil
}

func VerifyRefreshToken(tokenString string) (string, error) {
	// Get secret key from environment variable
	secret := os.Getenv("REFRESH_TOKEN_KEY")

	if secret == "" {
		return "", fmt.Errorf("secret key not set in environment: %s", "ACCESS_TOKEN_KEY")
	}

	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	// Extract claims safely
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("sub claim missing or invalid")
	}

	return userID, nil
}
