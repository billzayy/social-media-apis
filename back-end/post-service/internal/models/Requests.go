package models

import "github.com/google/uuid"

type AddPostRequest struct {
	UserId    uuid.UUID `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"createdAt"`
}
