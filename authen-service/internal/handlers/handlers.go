package handlers

import (
	"github.com/billzayy/social-media/authen-service/internal/services"
)

type Handlers struct {
	AuthHandler *AuthHandler
}

func NewHandlers(us *services.AuthService) *Handlers {
	return &Handlers{
		AuthHandler: NewAuthHandler(us),
	}
}
