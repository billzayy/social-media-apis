package services

import (
	"github.com/billzayy/social-media/back-end/chat-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/chat-service/internal/utils"
)

type Services struct {
	ChatService *ChatService
}

func NewServices(ur *repositories.ChatRepository, rS *RabbitService, wS *utils.WebSocketServer) *Services {
	return &Services{
		ChatService: NewChatService(ur, rS, wS),
	}
}
