package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"ID"`
	UserName       string    `json:"UserName"`
	FullName       string    `json:"FullName"`
	Email          string    `json:"Email"`
	Password       string    `json:"Password,omitempty"`
	UserSince      string    `json:"UserSince"`
	LastLogin      string    `json:"LastLogin"`
	Active         bool      `json:"Active"`
	ProfilePicture string    `json:"ProfilePicture,omitempty"`
}

type UserToken struct {
	UserId uuid.UUID `json:"ID"`
	Token  string    `json:"Token"`
	Type   string    `json:"Bearer"`
}
