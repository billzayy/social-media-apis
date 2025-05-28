package handlers

import (
	"github.com/billzayy/social-media/back-end/user-service/internal/services"
)

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandlers(sv *services.Services) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(sv.UserService),
	}
}
