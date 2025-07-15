package models

import "time"

type ReqUpdateUser struct {
	ID             string    `json:"ID"`
	UserName       string    `json:"UserName"`
	FullName       string    `json:"FullName"`
	Email          string    `json:"Email"`
	ProfilePicture string    `json:"ProfilePicture,omitempty"`
	Description    string    `json:"Description,omitempty"`
	Location       string    `json:"Location,omitempty"`
	Website        []string  `json:"Website"`
	BirthDate      time.Time `json:"BirthDate,omitempty"`
	Theme          string    `json:"Theme,omitempty"`
	Language       string    `json:"Language,omitempty"`
	Country        string    `json:"Country,omitempty"`
}

type ReqUpdatePassword struct {
	ID          string `json:"ID"`
	OldPassword string `json:"OldPassword"`
	NewPassword string `json:"NewPassword"`
}

type ResponseUser struct {
	ID             string    `json:"ID"`
	UserName       string    `json:"UserName"`
	FullName       string    `json:"FullName"`
	Email          string    `json:"Email"`
	UserSince      time.Time `json:"UserSince"`
	LastLogin      time.Time `json:"LastLogin"`
	Active         bool      `json:"Active"`
	ProfilePicture string    `json:"ProfilePicture,omitempty"`
	Description    string    `json:"Description,omitempty"`
	Location       string    `json:"Location,omitempty"`
	Website        []string  `json:"Website"`
	BirthDate      time.Time `json:"BirthDate,omitempty"`
	Theme          string    `json:"Theme,omitempty"`
	Language       string    `json:"Language,omitempty"`
	Country        string    `json:"Country,omitempty"`
}

type SwaggerUserResp struct {
	StatusCode int          `json:"StatusCode"`
	Data       ResponseUser `json:"Data"`
	Message    string       `json:"Message"`
}
