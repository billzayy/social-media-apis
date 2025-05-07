package models

import "github.com/google/uuid"

type PostMedia struct {
	PostId uuid.UUID `json:"postId"`
	Type   string    `json:"type"`
	Url    string    `json:"url"`
}
