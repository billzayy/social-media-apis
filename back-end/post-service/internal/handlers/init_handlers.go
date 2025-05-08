package handlers

import (
	"github.com/billzayy/social-media/back-end/post-service/internal/services"
)

type Handlers struct {
	PostHandler     *PostHandler
	InteractHandler *InteractHandler
}

func NewHandlers(sv *services.Services) *Handlers {
	return &Handlers{
		PostHandler:     NewPostHandler(sv.PostService),
		InteractHandler: NewInteractHandler(sv.InteractService),
	}
}
