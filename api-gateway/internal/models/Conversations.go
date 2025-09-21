package models

import "github.com/google/uuid"

type Conversations struct {
	ID        uuid.UUID `json:"ID"`
	Name      string    `json:"Name"`
	IsGroup   bool      `json:"IsGroup"`
	CreatedBy uuid.UUID `json:"CreatedBy"`
	CreatedAt uuid.UUID `json:"CreatedAt"`
}

type ConversationMembers struct {
	UserId         uuid.UUID `json:"UserId"`
	ConversationId uuid.UUID `json:"ConversationId"`
	JoinedAt       string    `json:"JoinedAt"`
}

type Messages struct {
	ID             uuid.UUID `json:"ID"`
	SenderId       uuid.UUID `json:"SenderId"`
	ReceiverId     uuid.UUID `json:"ReceiverId"`
	ConversationId uuid.UUID `json:"ConversationId"`
	Content        string    `json:"Content"`
	MessageType    string    `json:"MessageType"`
	CreatedAt      string    `json:"CreatedAt"`
}

type ReceiveSocketMsg struct {
	SenderId       uuid.UUID `json:"SenderId"`
	ConversationId uuid.UUID `json:"ConversationId"`
	Content        string    `json:"Content"`
	MessageType    string    `json:"MessageType"`
	CreatedAt      string    `json:"CreatedAt"`
}
