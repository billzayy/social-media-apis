package handlers

import "github.com/billzayy/social-media/back-end/post-service/internal/db/repositories"

type Handlers struct {
	PostHandler *PostHandler
}

func NewHandlers(ur *repositories.PostRepository) *Handlers {
	return &Handlers{
		PostHandler: NewPostHandler(ur),
	}
}
