package handlers

import "github.com/billzayy/social-media/back-end/authen-service/internal/db/repositories"

type Handlers struct {
	UserHandler *AuthHandler
}

func NewHandlers(ur *repositories.AuthRepository) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(ur),
	}
}
