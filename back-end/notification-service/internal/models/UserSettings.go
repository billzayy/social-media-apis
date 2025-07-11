package models

import "github.com/google/uuid"

type themeMode string

const (
	LightMode  themeMode = "light"
	DarkMode   themeMode = "dark"
	SystemMode themeMode = "system"
)

type UserSettings struct {
	UserId             uuid.UUID `json:"UserId"`
	Theme              themeMode `json:"Theme"`
	Language           string    `json:"Language"`
	Country            string    `json:"Country"`
	EmailNotifications bool      `json:"EmailNotifications"`
	PushNotifications  bool      `json:"PushNotifications"`
}

type RespUserNotifySetting struct {
	UserId             uuid.UUID `json:"UserId"`
	EmailNotifications bool      `json:"EmailNotifications"`
	PushNotifications  bool      `json:"PushNotifications"`
}
