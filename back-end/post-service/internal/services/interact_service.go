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

func (iS *InteractService) CheckLikeOnPostService(userId uuid.UUID, postId uuid.UUID) (bool, error) {
	data, err := iS.InteractRepository.GetLikeOnPost(userId, postId)

	if err != nil {
		return false, err
	}

	if data == 0 {
		return false, nil
	}

	return true, nil
}

func (iS *InteractService) AddLikeService(userId uuid.UUID, postId uuid.UUID) error {
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

func (iS *InteractService) AddCommentService(reqComm models.CommentRequest) error {
	err := iS.InteractRepository.AddComment(reqComm)

	if err != nil {
		return err
	}

	return nil
}

func (iS *InteractService) DeleteCommentService(id uuid.UUID) error {
	rows, err := iS.InteractRepository.DeleteComment(id)

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}
