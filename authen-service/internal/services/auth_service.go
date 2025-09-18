package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/billzayy/social-media/authen-service/internal/db/repositories"
	"github.com/billzayy/social-media/authen-service/internal/middleware"
	"github.com/billzayy/social-media/authen-service/internal/models"
	"github.com/google/uuid"
)

type AuthService struct {
	AuthRepository *repositories.AuthRepository
}

func NewAuthService(ar *repositories.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: ar,
	}
}

func (as *AuthService) RegisterService(request models.RegisterRequest) (bool, error) {
	fullName := request.FirstName + request.Surname

	hashedPass, err := middleware.HashPassword(request.Password)

	if err != nil {
		return false, err
	}

	checkedEmail, err := as.AuthRepository.CheckEmail(request.Email)

	if err != nil {
		return false, fmt.Errorf("failed to register account")
	}

	if checkedEmail {
		return false, fmt.Errorf("account existed!")
	}

	id, err := as.AuthRepository.AddUser(request.UserName, fullName, request.Email, hashedPass)

	if err != nil {
		return false, err
	}

	err = as.AuthRepository.AddUserProfile(id)

	if err != nil {
		return false, err
	}

	err = as.AuthRepository.AddUserSettings(id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (as *AuthService) LoginService(userName string, password string) (models.UserResponse, http.Cookie, error) {
	hashed, err := as.AuthRepository.GetHashedPassword(userName)

	if err != nil {
		return models.UserResponse{}, http.Cookie{}, err
	}

	valid, err := middleware.ValidatePassword(password, hashed)

	if err != nil {
		return models.UserResponse{}, http.Cookie{}, fmt.Errorf("wrong password")
	}

	if !valid {
		return models.UserResponse{}, http.Cookie{}, fmt.Errorf("invalid")
	}

	data, err := as.AuthRepository.GetUser(userName, hashed)

	if err != nil {
		return models.UserResponse{}, http.Cookie{}, err
	}

	if uuid.MustParse(data.ID) == uuid.Nil {
		return models.UserResponse{}, http.Cookie{}, fmt.Errorf("not found account")
	}

	accessToken, expires, cookie, err := middleware.GenerateTokens(data.ID)

	if err != nil {
		return models.UserResponse{}, http.Cookie{}, err
	}

	errChan := make(chan error, 2) // Make a slice of channels

	go func() { // Concurrency
		errChan <- as.AuthRepository.UpdateLoginTime(data.ID)
	}()

	for range 1 {
		if err := <-errChan; err != nil {
			return models.UserResponse{}, http.Cookie{}, err
		}
	}

	return models.UserResponse{User: data, Token: accessToken, Type: "Bearer", ExpiresIn: expires}, cookie, nil
}

func (as *AuthService) RefreshTokenService(cookieToken string) (models.GetUserData, string, int64, http.Cookie, error) {
	ctx := context.Background()

	userId, err := middleware.VerifyToken(cookieToken, "REFRESH_TOKEN_KEY")

	if err != nil {
		return models.GetUserData{}, "", 0, http.Cookie{}, err
	}

	newToken, expires, cookie, err := middleware.GenerateTokens(userId)

	if err != nil {
		return models.GetUserData{}, "", expires, http.Cookie{}, err
	}

	data, err := as.AuthRepository.GetUserById(userId)

	if err != nil {
		return models.GetUserData{}, "", expires, http.Cookie{}, err
	}

	errChan := make(chan error, 1)

	go func() { // Concurrency
		errChan <- as.AuthRepository.SaveUserRedis(ctx, userId, newToken)
	}()

	if err := <-errChan; err != nil {
		return models.GetUserData{}, "", expires, http.Cookie{}, err
	}

	return data, newToken, expires, cookie, nil
}
