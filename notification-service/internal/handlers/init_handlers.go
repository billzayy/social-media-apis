package handlers

import (
	"github.com/billzayy/social-media/notification-service/internal/services"
)

type Handlers struct {
	NotifyHandler *NotifyHandler
}

func NewHandlers(sv *services.Services) *Handlers {
	return &Handlers{
		NotifyHandler: NewNotifyHandler(sv.NotifyService),
	}
}
