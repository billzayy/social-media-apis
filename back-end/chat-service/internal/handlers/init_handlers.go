package handlers

import (
	"github.com/billzayy/social-media/back-end/chat-service/internal/services"
)

type Handlers struct {
	ChatHandler *ChatHandler
}

func NewHandlers(sv *services.Services) *Handlers {
	return &Handlers{
		ChatHandler: NewChatHandler(sv.ChatService),
	}
}
