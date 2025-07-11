package services

import "github.com/billzayy/social-media/back-end/notification-service/internal/db/repositories"

type Services struct {
	NotifyService *NotifyService
}

func NewServices(ur *repositories.NotifyRepository, rS *RabbitService) *Services {
	return &Services{
		NotifyService: NewUserService(ur, rS),
	}
}
