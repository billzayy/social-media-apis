package handlers

import (
	"github.com/billzayy/social-media/back-end/authen-service/internal/services"
)

type Handlers struct {
	AuthHandler *AuthHandler
}

func NewHandlers(us *services.AuthService) *Handlers {
	return &Handlers{
		AuthHandler: NewAuthHandler(us),
	}
}
