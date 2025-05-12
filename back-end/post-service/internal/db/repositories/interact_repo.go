package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/billzayy/social-media/back-end/post-service/internal/models"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type InteractRepositoryInterface interface {
	AddLike(userId uuid.UUID, postId uuid.UUID) error
	DeleteLike(userId uuid.UUID, postId uuid.UUID) (int64, error)
	AddComment(reqComm models.CommentRequest) error
	DeleteComment(id uuid.UUID) (int64, error)
}

type InteractRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewInteractRepository(db *sql.DB, rdb *redis.Client) *InteractRepository {
	return &InteractRepository{
		db:  db,
		rdb: rdb,
	}
}

func (iR *InteractRepository) GetLikeOnPost(userId uuid.UUID, postId uuid.UUID) (int, error) {
	query := `SELECT COUNT(*) FROM public."PostLikes" WHERE "UserId" = $1 AND "PostId" = $2;`

	rows, err := iR.db.Query(query, userId, postId)

	if err != nil {
		return 0, err
	}

	var count int
	for rows.Next() {
		err := rows.Scan(&count)

		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (iR *InteractRepository) AddLike(userId uuid.UUID, postId uuid.UUID) error {
	query := `INSERT INTO public."PostLikes" ("UserId", "PostId", "DateLike") VALUES ($1, $2, $3)`
	_, err := iR.db.Exec(query, userId, postId, time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (iR *InteractRepository) DeleteLike(userId uuid.UUID, postId uuid.UUID) (int64, error) {
	query := fmt.Sprintf(`DELETE FROM public."PostLikes" WHERE "UserId" = '%s' AND "PostId" = '%s'`, userId, postId)

	affected, err := iR.db.Exec(query)

	if err != nil {
		return 0, err
	}

	resp, err := affected.RowsAffected()

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (iR *InteractRepository) AddComment(reqComm models.CommentRequest) error {
	query := `INSERT INTO public."PostComments" ("UserId", "PostId", "Comment", "SentDate") VALUES($1,$2,$3,$4)`

	_, err := iR.db.Exec(query, reqComm.UserId, reqComm.PostId, reqComm.Comment, time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (iR *InteractRepository) DeleteComment(id uuid.UUID) (int64, error) {
	query := fmt.Sprintf(`DELETE FROM public."PostComments" WHERE "ID" = '%s'`, id)

	affected, err := iR.db.Exec(query)

	if err != nil {
		return 0, err
	}

	resp, err := affected.RowsAffected()

	if err != nil {
		return 0, err
	}

	if resp == 0 {
		return resp, fmt.Errorf("comment not found")
	}

	return resp, nil
}
