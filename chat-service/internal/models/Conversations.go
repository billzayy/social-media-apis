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
