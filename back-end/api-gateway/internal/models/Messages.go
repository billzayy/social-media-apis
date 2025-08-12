package models

import (
	"github.com/google/uuid"
)

type SwaggerReqSendMessage struct {
	SenderId    uuid.UUID `json:"SenderId"`
	ReceiverId  uuid.UUID `json:"ReceiverId"`
	Content     string    `json:"Content"`
	MessageType string    `json:"MessageType"`
}

type SwaggerRespGetMessage struct {
	StatusCode int              `json:"StatusCode"`
	Data       []RespGetMessage `json:"Data"`
	Message    string           `json:"Message"`
}

type RespGetMessage struct {
	Id          uuid.UUID `json:"Id"`
	Sender      User      `json:"Sender"`
	Content     string    `json:"Content"`
	MessageType string    `json:"MessageType"`
	CreatedAt   string    `json:"CreatedAt"`
}

type User struct {
	UserId   uuid.UUID `json:"UserId"`
	UserName string    `json:"UserName"`
	Avatar   string    `json:"Avatar,omitempty"`
}
