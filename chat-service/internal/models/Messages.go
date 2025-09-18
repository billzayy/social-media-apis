package models

import (
	"github.com/google/uuid"
)

type Messages struct {
	ID             uuid.UUID `json:"ID"`
	SenderId       uuid.UUID `json:"SenderId"`
	ReceiverId     uuid.UUID `json:"ReceiverId"`
	ConversationId uuid.UUID `json:"ConversationId"`
	Content        string    `json:"Content"`
	MessageType    string    `json:"MessageType"`
	CreatedAt      string    `json:"CreatedAt"`
}

type ReqSendMessage struct {
	SenderId    uuid.UUID `json:"SenderId"`
	ReceiverId  uuid.UUID `json:"ReceiverId"`
	Content     string    `json:"Content"`
	MessageType string    `json:"MessageType"`
}

type ReceiveSocketMsg struct {
	SenderId       uuid.UUID `json:"SenderId"`
	ConversationId uuid.UUID `json:"ConversationId"`
	Content        string    `json:"Content"`
	MessageType    string    `json:"MessageType"`
	CreatedAt      string    `json:"CreatedAt"`
}

type RespGetMessage struct {
	ID          uuid.UUID `json:"Id"`
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
