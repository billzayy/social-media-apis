package services

import "github.com/billzayy/social-media/back-end/authen-service/internal/db/repositories"

type Services struct {
	AuthService *AuthService
}

func NewServices(ar *repositories.AuthRepository) *Services {
	return &Services{
		AuthService: NewAuthService(ar),
	}
}
