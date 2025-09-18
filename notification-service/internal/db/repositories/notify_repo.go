package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/billzayy/social-media/notification-service/internal/models"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type UserRepositoryInterface interface {
}

type NotifyRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewNotifyRepository(db *sql.DB, rdb *redis.Client) *NotifyRepository {
	return &NotifyRepository{
		db:  db,
		rdb: rdb,
	}
}

func (nR *NotifyRepository) GetNotificationsByReceiverId(receiverId uuid.UUID) ([]models.Notifications, error) {
	rows, err := nR.db.Query(getNotifyDataByIdQuery, receiverId)

	if err != nil {
		return []models.Notifications{}, err
	}

	defer rows.Close()

	var result []models.Notifications
	var temp models.Notifications

	for rows.Next() {
		err := rows.Scan(
			&temp.Id, &temp.SenderId, &temp.ReceiverId,
			&temp.Messages, &temp.Type, &temp.Url, &temp.IsRead, &temp.Date)

		if err != nil {
			return []models.Notifications{}, err
		}

		result = append(result, temp)
	}

	return result, nil
}

func (nR *NotifyRepository) CountUnreadNotifications(userId uuid.UUID) (int, error) {
	rows, err := nR.db.Query(getUnreadNotify, userId)

	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var result int

	for rows.Next() {
		err := rows.Scan(&result)

		if err != nil {
			return -1, err
		}
	}

	return result, nil
}

func (nR *NotifyRepository) GetNotifySettingByUserId(userId uuid.UUID) (models.RespUserNotifySetting, error) {
	rows, err := nR.db.Query(getUserSettingQuery, userId)

	if err != nil {
		return models.RespUserNotifySetting{}, err
	}
	defer rows.Close()

	var result models.RespUserNotifySetting

	for rows.Next() {
		err := rows.Scan(&result.UserId, &result.EmailNotifications, &result.PushNotifications)

		if err != nil {
			return result, err
		}
	}

	return result, nil
}

func (nR *NotifyRepository) SaveDataNotify(req models.ReqSendNotify) error {
	_, err := nR.db.Exec(saveNotifyQuery, req.SenderId, req.ReceiverId, req.Messages, req.Type, req.Url)

	if err != nil {
		return err
	}

	return nil
}

func (nR *NotifyRepository) UpdateReadStatus(id uuid.UUID) error {
	result, err := nR.db.Exec(updateRead, id)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}

func (nR *NotifyRepository) SaveRedisNotify(ctx context.Context, userId string, notifyId string) error {
	_, err := nR.rdb.HSet(ctx, "notificationList", userId, notifyId).Result()

	if err != nil {
		return err
	}

	return nil
}
