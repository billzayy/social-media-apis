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

type UserResponse struct {
	User      GetUserData `json:"User"`
	Token     string      `json:"Token"`
	Type      string      `json:"Bearer"`
	ExpiresIn int64       `json:"ExpiresIn"`
}

type GetUserData struct {
	Id             string `json:"ID"`
	FullName       string `json:"FullName"`
	Email          string `json:"Email"`
	ProfilePicture string `json:"ProfilePicture"`
}

type RefreshTokenResp struct {
	User      GetUserData `json:"User"`
	Token     string      `json:"Token"`
	Type      string      `json:"Bearer"`
	ExpiresIn int64       `json:"ExpiresIn"`
}
