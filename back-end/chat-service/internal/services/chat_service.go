package services

import (
	"fmt"
	"sync"

	"github.com/billzayy/social-media/back-end/chat-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/chat-service/internal/models"
	"github.com/billzayy/social-media/back-end/chat-service/internal/utils"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitService struct {
	Channel *amqp.Channel
}

type ChatService struct {
	ChatRepository *repositories.ChatRepository
	RabbitMQ       *RabbitService
	Websocket      *utils.WebSocketServer
}

func NewChatService(cR *repositories.ChatRepository, rS *RabbitService, ws *utils.WebSocketServer) *ChatService {
	return &ChatService{
		ChatRepository: cR,
		RabbitMQ:       rS,
		Websocket:      ws,
	}
}

func (cS *ChatService) SendPrivateConversation(req models.ReqSendMessage) (uuid.UUID, error) {
	roomId, err := cS.ChatRepository.CheckPrivateConversationMember(req.SenderId, req.ReceiverId)

	if roomId == uuid.Nil || err != nil {
		id, err := cS.ChatRepository.CreatePrivateConversation(req.SenderId)

		if err != nil {
			return uuid.UUID{}, err
		}

		// Use a WaitGroup to synchronize goroutines
		var wg sync.WaitGroup
		// Use a channel to collect errors from goroutines
		errChan := make(chan error, 2)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := cS.ChatRepository.CreateConversationMember(req.SenderId, id); err != nil {
				errChan <- fmt.Errorf("failed to add sender to conversation: %w", err)
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := cS.ChatRepository.CreateConversationMember(req.ReceiverId, id); err != nil {
				errChan <- fmt.Errorf("failed to add receiver to conversation: %w", err)
			}
		}()

		// Wait for both goroutines to complete
		wg.Wait()
		// Close the error channel to allow ranging over it
		close(errChan)

		for err := range errChan {
			if err != nil {
				return uuid.UUID{}, err // Return the first error encountered
			}
		}

		err = cS.Websocket.BroadcastToRoom(id, req.SenderId, req.ReceiverId, req.Content)

		if err != nil {
			return uuid.UUID{}, err
		}

		err = cS.ChatRepository.SaveChat(req.SenderId, id, req.Content, req.MessageType)

		if err != nil {
			return uuid.UUID{}, err
		}

		return id, nil
	} else {
		err = cS.Websocket.BroadcastToRoom(roomId, req.SenderId, req.ReceiverId, req.Content)

		if err != nil {
			return uuid.UUID{}, err
		}

		err := cS.ChatRepository.SaveChat(req.SenderId, roomId, req.Content, req.MessageType)

		if err != nil {
			return uuid.UUID{}, err
		}

		return roomId, nil
	}
}

func (cS *ChatService) GetMessage(conversationId uuid.UUID, limit, offset int) ([]models.RespGetMessage, error) {
	data, err := cS.ChatRepository.GetPrivateChat(conversationId, limit, offset)

	if err != nil {
		return data, err
	}
	return data, nil
}

func (cS *ChatService) DeleteMessage(messageId uuid.UUID) error {
	err := cS.ChatRepository.DeleteChat(messageId)

	if err != nil {
		return err
	}

	return nil
}
