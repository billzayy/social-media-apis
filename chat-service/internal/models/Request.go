package models

import "github.com/google/uuid"

type ReqPrivate struct {
	SenderId uuid.UUID `json:"SenderId"`
}

type ReqGroup struct {
	Id uuid.UUID `json:"Id"`
}
