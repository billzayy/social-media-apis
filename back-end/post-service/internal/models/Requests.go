package models

import "github.com/google/uuid"

type AddPostRequest struct {
	UserId    uuid.UUID `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"createdAt"`
}

type LikeRequest struct {
	UserId uuid.UUID `json:"userId"`
	PostId uuid.UUID `json:"postId"`
}

type CommentRequest struct {
	UserId  uuid.UUID `json:"userId"`
	PostId  uuid.UUID `json:"postId"`
	Comment string    `json:"comment"`
}
