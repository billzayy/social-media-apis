package services

import "github.com/billzayy/social-media/back-end/post-service/internal/db/repositories"

type Services struct {
	PostService *PostService
}

func NewServices(pr *repositories.PostRepository) *Services {
	return &Services{
		PostService: NewPostService(pr),
	}
}
