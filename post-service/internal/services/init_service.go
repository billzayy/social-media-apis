package services

import "github.com/billzayy/social-media/post-service/internal/db/repositories"

type Services struct {
	PostService     *PostService
	InteractService *InteractService
}

func NewServices(pr *repositories.PostRepository, ir *repositories.InteractRepository) *Services {
	return &Services{
		PostService:     NewPostService(pr),
		InteractService: NewInteractService(ir),
	}
}
