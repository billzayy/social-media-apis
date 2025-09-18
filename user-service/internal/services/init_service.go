package services

import "github.com/billzayy/social-media/user-service/internal/db/repositories"

type Services struct {
	UserService *UserService
}

func NewServices(ur *repositories.UserRepository) *Services {
	return &Services{
		UserService: NewUserService(ur),
	}
}
