package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

func (iR *InteractRepository) GetIndexPostRedis(ctx context.Context, postId string) (int, models.PostResp, error) {
	key := "postList"

	rawJSON, err := iR.rdb.Do(ctx, "JSON.GET", key, "$").Result()
	if err != nil {
		log.Printf("Failed to get posts from Redis: %v", err)
		return -1, models.PostResp{}, err
	}

	var wrapped [][]models.PostResp
	jsonStr, ok := rawJSON.(string)
	if !ok {
		return -1, models.PostResp{}, fmt.Errorf("unexpected type for JSON.GET result: %T", rawJSON)
	}

	if err := json.Unmarshal([]byte(jsonStr), &wrapped); err != nil {
		log.Printf("Failed to unmarshal posts: %v", err)
		return -1, models.PostResp{}, err
	}

	if len(wrapped) == 0 {
		return -1, models.PostResp{}, fmt.Errorf("no posts found")
	}

	posts := wrapped[0]

	for i, post := range posts {
		if post.PostId == uuid.MustParse(postId) {
			return i, posts[i], nil
		}
	}

	return -1, models.PostResp{}, fmt.Errorf("postId %s not found", postId)
}

func (iR *InteractRepository) UpdateInteractRedis(ctx context.Context, index int, content models.PostResp) error {
	key := "postList"

	path := fmt.Sprintf("$[%d]", index)

	postJSON, err := json.Marshal(content)
	if err != nil {
		log.Printf("Failed to marshal new post: %v", err)
		return err
	}

	_, err = iR.rdb.Do(ctx, "JSON.SET", key, path, string(postJSON)).Result()
	if err != nil {
		log.Printf("Failed to update post at index %d: %v", index, err)
		return err
	}

	return nil
}

func (iR *InteractRepository) RemoveInteractRedis(ctx context.Context, index int) error {
	key := "posts"
	path := fmt.Sprintf("$[%d]", index) // Remove element at this index

	_, err := iR.rdb.Do(ctx, "JSON.DEL", key, path).Result()
	if err != nil {
		log.Printf("Failed to delete post at index %d: %v", index, err)
		return err
	}

	return nil
}
