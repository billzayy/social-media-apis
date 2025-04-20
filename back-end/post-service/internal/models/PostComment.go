package models

import "github.com/google/uuid"

type PostComment struct {
	ID       uuid.UUID `json:"ID"`
	UserId   uuid.UUID `json:"UserId"`
	Comment  string    `json:"Comment"`
	SendDate string    `json:"SendDate"`
	PostId   uuid.UUID `json:"PostId"`
}
