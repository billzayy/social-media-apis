package services

import (
	"context"
	"fmt"

	"github.com/billzayy/social-media/back-end/post-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/post-service/internal/models"
	"github.com/billzayy/social-media/back-end/post-service/internal/utils"
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
	ctx := context.Background()
	if req.Content == "" {
		return false, fmt.Errorf("content can not empty")
	}

	err := pS.PostRepository.AddPost(req)

	if err != nil {
		return false, err
	}

	data, err := pS.PostRepository.GetPost()

	if err != nil {
		return false, err
	}

	err = pS.PostRepository.AddPostRedis(ctx, data[0])

	if err != nil {
		return false, err
	}

	return true, nil
}

func (pS *PostService) GetPost() ([]models.PostResp, error) {
	ctx := context.Background()

	res, err := pS.PostRepository.GetPostRedis(ctx)

	if err != nil {
		return []models.PostResp{}, err
	}

	if len(res) == 0 {
		data, err := pS.PostRepository.GetPost()

		if err != nil {
			return []models.PostResp{}, err
		}

		if len(data) == 0 {
			return []models.PostResp{}, nil
		}

		for _, v := range data {
			err := pS.PostRepository.AddPostRedis(ctx, v)

			if err != nil {
				return []models.PostResp{}, err
			}
		}

		return data, err
	}

	return utils.SortPostWithCreatedTime(res), nil
}

func (pS *PostService) DeletePost(id string) error {
	ctx := context.Background()
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

	err = pS.PostRepository.DeletePostByIdRedis(ctx, id)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
