package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ResponseUser struct {
	ID             uuid.UUID      `json:"ID"`
	UserName       string         `json:"UserName"`
	FullName       string         `json:"FullName"`
	Email          string         `json:"Email"`
	UserSince      time.Time      `json:"UserSince"`
	LastLogin      time.Time      `json:"LastLogin"`
	Active         bool           `json:"Active"`
	ProfilePicture string         `json:"ProfilePicture,omitempty"`
	Description    string         `json:"Description,omitempty"`
	Location       string         `json:"Location,omitempty"`
	Website        pq.StringArray `json:"Website"`
	BirthDate      time.Time      `json:"BirthDate,omitempty"`
	Theme          string         `json:"Theme,omitempty"`
	Language       string         `json:"Language,omitempty"`
	Country        string         `json:"Country,omitempty"`
}
