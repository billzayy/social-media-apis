package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/billzayy/social-media/back-end/authen-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/authen-service/internal/middleware"
	"github.com/billzayy/social-media/back-end/authen-service/internal/models"
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
		return false, err
	}

	if checkedEmail {
		return false, fmt.Errorf("failed to register account")
	}

	err = as.AuthRepository.AddUser(request.UserName, fullName, request.Email, hashedPass)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (as *AuthService) LoginService(userName string, password string) (models.UserToken, http.Cookie, error) {
	ctx := context.Background()

	hashed, err := as.AuthRepository.GetHashedPassword(userName)

	if err != nil {
		return models.UserToken{}, http.Cookie{}, err
	}

	valid, err := middleware.ValidatePassword(password, hashed)

	if err != nil {
		return models.UserToken{}, http.Cookie{}, err
	}

	if !valid {
		return models.UserToken{}, http.Cookie{}, fmt.Errorf("invalid")
	}

	userId, err := as.AuthRepository.GetUserId(userName, hashed)

	if err != nil {
		return models.UserToken{}, http.Cookie{}, err
	}

	if uuid.MustParse(userId) == uuid.Nil {
		return models.UserToken{}, http.Cookie{}, fmt.Errorf("not found account")
	}

	accessToken, cookie, err := middleware.GenerateTokens(userId)

	if err != nil {
		return models.UserToken{}, http.Cookie{}, err
	}

	err = as.AuthRepository.SaveUserRedis(ctx, userId, accessToken)

	if err != nil {
		return models.UserToken{}, http.Cookie{}, err
	}

	return models.UserToken{UserId: uuid.MustParse(userId), Token: accessToken, Type: "Bearer"}, cookie, nil
}

func (as *AuthService) RefreshTokenService(cookieToken string) (string, string, http.Cookie, error) {
	// Check the refresh token is valid or not
	userId, err := middleware.VerifyToken(cookieToken, "REFRESH_TOKEN_KEY")

	if err != nil {
		return "", "", http.Cookie{}, err
	}

	// Create new access & refresh token
	newToken, cookie, err := middleware.GenerateTokens(userId)

	if err != nil {
		return "", "", http.Cookie{}, err
	}

	return userId, newToken, cookie, nil
}
