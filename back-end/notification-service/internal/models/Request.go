package models

import (
	"github.com/google/uuid"
)

type ReqSendNotify struct {
	SenderId   uuid.UUID        `json:"SenderId"`
	ReceiverId uuid.UUID        `json:"ReceiverId"`
	Messages   string           `json:"Messages"`
	Type       NotificationType `json:"Type"`
	Url        string           `json:"Url"`
}
