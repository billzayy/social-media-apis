package models

import "github.com/google/uuid"

type PostLikes struct {
	UserId   uuid.UUID `json:"UserID"`
	PostId   uuid.UUID `json:"PostId"`
	DateLike string    `json:"DateLike"`
}
