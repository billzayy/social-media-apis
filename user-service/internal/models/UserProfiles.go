package models

import "github.com/google/uuid"

type UserProfiles struct {
	UserId      uuid.UUID `json:"UserId"`
	Description string    `json:"Description"`
	Location    string    `json:"Location"`
	Website     []string  `json:"Website"`
	BirthDate   string    `json:"BirthDate"`
}
