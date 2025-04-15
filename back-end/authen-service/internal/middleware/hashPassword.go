package middleware

import (
	"golang.org/x/crypto/bcrypt"
)

const salt = bcrypt.DefaultCost

func HashPassword(password string) (string, error) {
	convertPass := []byte(password)

	hashPassword, err := bcrypt.GenerateFromPassword(convertPass, salt)

	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func ValidatePassword(password string, hashedPassword string) (bool, error) {
	convertPass := []byte(password)
	convertHashed := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(convertHashed, convertPass)

	if err != nil {
		return false, err
	}

	return true, nil
}
