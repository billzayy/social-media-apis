package services

import (
	"context"
	"fmt"

	"github.com/billzayy/social-media/user-service/internal/db/repositories"
	"github.com/billzayy/social-media/user-service/internal/middleware"
	"github.com/billzayy/social-media/user-service/internal/models"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(ir *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: ir,
	}
}

func (ur *UserService) GetUserService(userId string) (models.ResponseUser, error) {
	ctx := context.Background()

	cache, err := ur.UserRepository.GetUserRedis(ctx, userId)

	if err != nil {
		return models.ResponseUser{}, err
	}

	if cache.UserName == "" { // It means User data is empty (not existed) on Redis's Data
		data, err := ur.UserRepository.GetUserById(userId)

		if err != nil {
			return models.ResponseUser{}, err
		}

		err = ur.UserRepository.AddUserRedis(ctx, userId, data)

		return data, nil
	}
	return cache, nil
}

func (ur *UserService) UpdateUserService(req models.ReqUpdateUser) error {
	ctx := context.Background()

	err := ur.UserRepository.UpdateInfoUser(req)

	if err != nil {
		return err
	}

	data, err := ur.UserRepository.GetUserById(req.ID.String())

	if err != nil {
		return err
	}

	err = ur.UserRepository.AddUserRedis(ctx, req.ID.String(), data)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserService) UpdatePasswordService(req models.ReqUpdatePassword) error {
	id := req.ID.String()

	hashedPass, err := ur.UserRepository.GetPassword(id)

	if err != nil {
		return fmt.Errorf("failed to get password: %w", err)
	}

	valid, err := middleware.ValidatePassword(req.OldPassword, hashedPass)

	if err != nil {
		return fmt.Errorf("failed to validate password: %w", err)
	}

	if !valid {
		return fmt.Errorf("old password is incorrect")
	}

	newPass, err := middleware.HashPassword(req.NewPassword)

	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	err = ur.UserRepository.UpdatePassword(id, newPass)

	if err != nil {
		return fmt.Errorf("failed to change password: %w", err)
	}

	return nil
}
