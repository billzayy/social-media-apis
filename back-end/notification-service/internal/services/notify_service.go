package services

import (
	"github.com/billzayy/social-media/back-end/notification-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/notification-service/internal/models"
	"github.com/billzayy/social-media/back-end/notification-service/internal/utils"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitService struct {
	Channel *amqp.Channel
}

type NotifyService struct {
	NotifyRepository *repositories.NotifyRepository
	RabbitMQ         *RabbitService
}

func NewUserService(nR *repositories.NotifyRepository, rS *RabbitService) *NotifyService {
	return &NotifyService{
		NotifyRepository: nR,
		RabbitMQ:         rS,
	}
}

func (nS *NotifyService) SendNotifyService(req models.ReqSendNotify) error {
	resp, err := nS.NotifyRepository.GetNotifySettingByUserId(req.ReceiverId)

	if err != nil {
		return err
	}

	//* Check User is DnD? (DnD: Do Not Disturb)
	if resp.PushNotifications == false {
		return nil
	}

	if resp.EmailNotifications == true {
		// Send to send email logic
	}

	// Add notification on Queue
	queue := nS.RabbitMQ

	err = utils.PublishMessage(queue.Channel, req.ReceiverId.String(), req)

	if err != nil {
		return nil
	}

	// Save to Databases
	err = nS.NotifyRepository.SaveDataNotify(req)

	if err != nil {
		return err
	}

	err = utils.SubscribeMessage(queue.Channel, req.ReceiverId.String())

	if err != nil {
		return err
	}

	return nil
}

func (nS *NotifyService) SendNotifyServiceGrpc(req models.ReqSendNotify) error {
	resp, err := nS.NotifyRepository.GetNotifySettingByUserId(req.ReceiverId)

	if err != nil {
		return err
	}

	//* Check User is DnD? (DnD: Do Not Disturb)
	if resp.PushNotifications == false {
		return nil
	}

	if resp.EmailNotifications == true {
		// Send to send email logic
	}

	// Save to Databases
	err = nS.NotifyRepository.SaveDataNotify(req)

	if err != nil {
		return err
	}

	return nil
}

func (nS *NotifyService) GetNotifyService(id string) ([]models.Notifications, error) {
	receiverId := uuid.MustParse(id)

	data, err := nS.NotifyRepository.GetNotificationsByReceiverId(receiverId)

	if err != nil {
		return []models.Notifications{}, err
	}

	return data, nil
}

func (nS *NotifyService) GetUnreadNotifyService(id string) (int, error) {
	receiverId := uuid.MustParse(id)

	data, err := nS.NotifyRepository.CountUnreadNotifications(receiverId)

	if err != nil {
		return -1, err
	}

	return data, nil
}

func (nS *NotifyService) UpdateReadService(id string) error {
	notifyId := uuid.MustParse(id)

	err := nS.NotifyRepository.UpdateReadStatus(notifyId)

	if err != nil {
		return err
	}

	return nil
}
