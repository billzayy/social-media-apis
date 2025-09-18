package models

import "github.com/google/uuid"

type Users struct {
	ID             uuid.UUID `json:"ID"`
	UserName       string    `json:"UserName"`
	FullName       string    `json:"FullName"`
	Email          string    `json:"Email"`
	Password       string    `json:"Password"`
	UserSince      string    `json:"UserSince"`
	LastLogin      string    `json:"LastLogin"`
	Active         bool      `json:"Active"`
	ProfilePicture string    `json:"ProfilePicture"`
}
