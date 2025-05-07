package models

import "github.com/google/uuid"

type Post struct {
	ID        uuid.UUID `json:"ID"`
	UserId    uuid.UUID `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"createdAt"`
}

type PostResp struct {
	PostId    uuid.UUID   `json:"postId"`
	Author    Users       `json:"userId"`
	Content   string      `json:"content"`
	Media     []PostMedia `json:"media,omitempty"`
	CreatedAt string      `json:"createdAt"`
	Likes     int         `json:"likes"`
	Comments  int         `json:"comments"`
	Shares    int         `json:"shares"`
}

type Users struct {
	UserId    uuid.UUID `json:"ID"`
	FullName  string    `json:"fullName"`
	AvatarURL string    `json:"profilePicture"`
}
