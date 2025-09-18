package models

import "github.com/google/uuid"

type UserSettings struct {
	UserId   uuid.UUID `json:"UserId"`
	Theme    string    `json:"Theme"`
	Language string    `json:"Language"`
	Country  string    `json:"Country"`
}
