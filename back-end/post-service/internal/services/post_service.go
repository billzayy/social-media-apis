package services

import (
	"fmt"

	"github.com/billzayy/social-media/back-end/post-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/post-service/internal/models"
)

type PostService struct {
	PostRepository *repositories.PostRepository
}

func NewPostService(pr *repositories.PostRepository) *PostService {
	return &PostService{
		PostRepository: pr,
	}
}

func (pS *PostService) CreatePost(req models.AddPostRequest) (bool, error) {
	if req.Content == "" {
		return false, fmt.Errorf("content can not empty")
	}

	err := pS.PostRepository.AddPost(req)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (pS *PostService) GetPost() ([]models.PostResp, error) {
	data, err := pS.PostRepository.GetPost()

	if err != nil {
		return data, err
	}

	return data, nil
}

func (pS *PostService) DeletePost(id string) error {
	if id == "" {
		return fmt.Errorf("id is not empty")
	}

	affected, err := pS.PostRepository.DeletePost(id)

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("post %s not found", id)
	}

	return nil
}
