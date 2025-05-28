package services

import "github.com/billzayy/social-media/back-end/user-service/internal/db/repositories"

type Services struct {
	UserService *UserService
}

func NewServices(ur *repositories.UserRepository) *Services {
	return &Services{
		UserService: NewUserService(ur),
	}
}
