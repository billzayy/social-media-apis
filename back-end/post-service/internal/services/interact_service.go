package services

import (
	"fmt"

	"github.com/billzayy/social-media/back-end/post-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/post-service/internal/models"
	"github.com/google/uuid"
)

type InteractService struct {
	InteractRepository *repositories.InteractRepository
}

func NewInteractService(ir *repositories.InteractRepository) *InteractService {
	return &InteractService{
		InteractRepository: ir,
	}
}

func (iS *InteractService) LikeService(userId string, postId string) error {
	err := iS.InteractRepository.AddLike(userId, postId)

	if err != nil {
		return err
	}

	return nil
}

func (iS *InteractService) RemoveLikeService(userId uuid.UUID, postId uuid.UUID) error {
	rowsAffected, err := iS.InteractRepository.DeleteLike(userId, postId)

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("not found id to remove")
	}

	return nil
}

func (iS *InteractService) CommentService(reqComm models.CommentRequest) error {
	err := iS.InteractRepository.AddComment(reqComm)

	if err != nil {
		return err
	}

	return nil
}

func (iS *InteractService) DeleteCommentService(id string) error {
	rows, err := iS.InteractRepository.DeleteComment(id)

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}
