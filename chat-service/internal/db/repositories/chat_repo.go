package repositories

import (
	"database/sql"
	"fmt"

	"github.com/billzayy/social-media/chat-service/internal/models"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ChatRepositoryInterface interface {
	CheckPrivateConversationMember()
	CreatePrivateConversation()
	CreateGroupConversation()
	CreateConversationMember()
	SaveChat()
	GetChat()
}

type ChatRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewChatRepository(db *sql.DB, rdb *redis.Client) *ChatRepository {
	return &ChatRepository{
		db:  db,
		rdb: rdb,
	}
}

func (cR *ChatRepository) CheckPrivateConversationMember(senderId, receiverId uuid.UUID) (uuid.UUID, error) {
	rows, err := cR.db.Query(getIdPrivateConversationQuery, senderId, receiverId)

	if err != nil {
		return uuid.UUID{}, err
	}

	defer rows.Close()

	var result uuid.UUID

	for rows.Next() {
		var temp string
		err := rows.Scan(&temp)

		if err != nil {
			return uuid.UUID{}, err
		}

		result, err = uuid.Parse(temp)

		if err != nil {
			return result, err
		}
	}

	if result == uuid.Nil {
		return result, fmt.Errorf("not found")
	}

	return result, nil
}

func (cR *ChatRepository) CreatePrivateConversation(id uuid.UUID) (uuid.UUID, error) {
	rows, err := cR.db.Query(createPrivateConversationQuery, id)

	if err != nil {
		return uuid.UUID{}, err
	}

	var result uuid.UUID

	for rows.Next() {
		err := rows.Scan(&result)

		if err != nil {
			return uuid.UUID{}, err
		}
	}

	return result, nil
}

func (cR *ChatRepository) CreateGroupConversation(id uuid.UUID) error {
	_, err := cR.db.Exec(createGroupConversationQuery, "New Group", true, id)

	if err != nil {
		return err
	}

	return nil
}

func (cR *ChatRepository) CreateConversationMember(userId, conversationId uuid.UUID) error {
	_, err := cR.db.Exec(addMemberQuery, userId, conversationId)

	if err != nil {
		return err
	}

	return nil
}

func (cR *ChatRepository) SaveChat(senderId, conversationId uuid.UUID, content, messageType string) error {
	_, err := cR.db.Exec(saveChatQuery, senderId, conversationId, content, messageType)

	if err != nil {
		return err
	}

	return nil
}

func (cR *ChatRepository) GetPrivateChat(conversationId uuid.UUID, limit, offset int) ([]models.RespGetMessage, error) {
	rows, err := cR.db.Query(getChatQuery, conversationId, limit, offset)

	if err != nil {
		return []models.RespGetMessage{}, err
	}
	defer rows.Close()

	var result []models.RespGetMessage

	for rows.Next() {
		var temp models.RespGetMessage
		var avatar sql.NullString

		err := rows.Scan(
			&temp.ID,
			&temp.Content, &temp.MessageType, &temp.CreatedAt,
			&temp.Sender.UserId, &temp.Sender.UserName, &avatar)

		if err != nil {
			return []models.RespGetMessage{}, err
		}

		if avatar.Valid {
			temp.Sender.Avatar = avatar.String
		} else {
			temp.Sender.Avatar = ""
		}

		result = append(result, temp)
	}

	return result, nil
}

func (cR *ChatRepository) DeleteChat(messageId uuid.UUID) error {
	_, err := cR.db.Exec(deleteChatQuery, messageId)

	if err != nil {
		return err
	}

	return nil
}
