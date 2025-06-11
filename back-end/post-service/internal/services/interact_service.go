package services

import (
	"context"
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
	ctx := context.Background()

	var fail error

	err := iS.InteractRepository.AddLike(userId, postId)

	if err != nil {
		return err
	}

	go func() {
		idx, data, err := iS.InteractRepository.GetIndexPostRedis(ctx, postId.String())

		if err != nil {
			fail = err
			return
		}

		data.Likes++

		err = iS.InteractRepository.UpdateInteractRedis(ctx, idx, data)

		if err != nil {
			fail = err
			return
		}

		fail = nil
	}()

	if fail != nil {
		return fail
	}

	return nil
}

func (iS *InteractService) RemoveLikeService(userId uuid.UUID, postId uuid.UUID) error {
	ctx := context.Background()
	var fail error
	rowsAffected, err := iS.InteractRepository.DeleteLike(userId, postId)

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("not found id to remove")
	}

	go func() {
		idx, data, err := iS.InteractRepository.GetIndexPostRedis(ctx, postId.String())

		if err != nil {
			fail = err
			return
		}

		if data.Likes > 0 {
			data.Likes--
		} else {
			fail = fmt.Errorf("can't remove likes")
			return
		}

		err = iS.InteractRepository.UpdateInteractRedis(ctx, idx, data)

		if err != nil {
			fail = err
			return
		}

		fail = nil
	}()

	if fail != nil {
		return fail
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
