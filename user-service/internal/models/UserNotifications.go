package models

import "github.com/google/uuid"

type UserNotifications struct {
	UserId uuid.UUID `json:"UserId"`
}
