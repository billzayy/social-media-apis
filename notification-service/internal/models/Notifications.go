package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationType string

const (
	Like      NotificationType = "like"
	Comment   NotificationType = "comment"
	AddFriend NotificationType = "addfriend"
)

type Notifications struct {
	Id         uuid.UUID        `json:"Id"`
	SenderId   uuid.UUID        `json:"SenderId"`
	ReceiverId uuid.UUID        `json:"ReceiverId"`
	Messages   string           `json:"Messages"`
	Type       NotificationType `json:"Type"`
	Url        string           `json:"Url"`
	IsRead     bool             `json:"IsRead"`
	Date       time.Time        `json:"Date"`
}
