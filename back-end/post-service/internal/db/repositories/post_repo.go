package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/billzayy/social-media/back-end/post-service/internal/models"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type PostRepositoryInterface interface {
	AddPost(request models.AddPostRequest) error
	GetPost() ([]models.PostResp, error)
	DeletePost(postId string) error
}

type PostRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewPostRepository(db *sql.DB, rdb *redis.Client) *PostRepository {
	return &PostRepository{
		db:  db,
		rdb: rdb,
	}
}

func (pR *PostRepository) AddPost(request models.AddPostRequest) error {
	query := fmt.Sprintf(`INSERT INTO public."Posts" ("UserId", "Content") VALUES ('%s', '%s')`,
		request.UserId, request.Content)

	_, err := pR.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (pR *PostRepository) GetPost() ([]models.PostResp, error) {
	rows, err := pR.db.Query(GetPostQuery)

	defer rows.Close()

	if err != nil {
		log.Fatalf("Error on Get post : %v", err)
		return []models.PostResp{}, err
	}

	var result []models.PostResp
	for rows.Next() {
		var post models.PostResp
		var userJSON []byte
		var mediaJSON []byte

		err := rows.Scan(
			&post.PostId, &userJSON,
			&post.Content, &post.CreatedAt,
			&mediaJSON, &post.Likes, &post.Comments)

		post.Shares = 0

		if err != nil {
			log.Fatalf("Error on Scan rows : %v", err)
			return []models.PostResp{}, err
		}

		// Decode user JSON
		if err := json.Unmarshal(userJSON, &post.Author); err != nil {
			return []models.PostResp{}, fmt.Errorf("Error decoding user JSON: %v", err)
		}

		// Decode media JSON
		if err := json.Unmarshal(mediaJSON, &post.Media); err != nil {
			return []models.PostResp{}, fmt.Errorf("Error decoding media JSON: %v", err)
		}

		result = append(result, post)
	}

	return result, nil
}

func (pR *PostRepository) DeletePost(postId string) (int64, error) {
	query := fmt.Sprintf(`DELETE FROM public."Posts" WHERE "ID" = '%s';`, postId)

	rowResult, err := pR.db.Exec(query)

	if err != nil {
		return 0, err
	}

	resp, err := rowResult.RowsAffected()

	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (pR *PostRepository) GetPostRedis(ctx context.Context) ([]models.PostResp, error) {
	rawData, err := pR.rdb.JSONGet(ctx, "postList", "$").Result()
	if err != nil {
		return []models.PostResp{}, err
	}

	if rawData == "" {
		return []models.PostResp{}, nil
	}

	var outer []json.RawMessage
	if err := json.Unmarshal([]byte(rawData), &outer); err != nil {
		return []models.PostResp{}, fmt.Errorf("failed to unmarshal outer array: %w", err)
	}

	if len(outer) == 0 {
		return []models.PostResp{}, nil
	}

	var posts []models.PostResp
	if err := json.Unmarshal(outer[0], &posts); err != nil {
		return []models.PostResp{}, fmt.Errorf("failed to unmarshal posts: %w", err)
	}

	return posts, nil
}

func (pR *PostRepository) AddPostRedis(ctx context.Context, input models.PostResp) error {
	// Serialize newPost to JSON string
	postJSON, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("failed to marshal post: %w", err)
	}

	// Check if postList key exists
	exists, err := pR.rdb.Exists(ctx, "postList").Result()
	if err != nil {
		return fmt.Errorf("failed to check key existence: %w", err)
	}

	if exists == 0 {
		// Key does not exist, create it as an array with the new post
		var posts []models.PostResp
		posts = append(posts, input)
		fullJSON, err := json.Marshal(posts)
		if err != nil {
			return fmt.Errorf("failed to marshal post array: %w", err)
		}

		_, err = pR.rdb.JSONSet(ctx, "postList", "$", fullJSON).Result()
		if err != nil {
			return fmt.Errorf("failed to JSONSet: %w", err)
		}

		return nil
	}

	// Key exists, append new post to the array
	_, err = pR.rdb.JSONArrAppend(ctx, "postList", "$", string(postJSON)).Result()
	if err != nil {
		return fmt.Errorf("failed to JSONArrAppend: %w", err)
	}

	return nil
}

func (pR *PostRepository) DeletePostByIdRedis(ctx context.Context, postId string) error {
	parseId, err := uuid.Parse(postId)

	if err != nil {
		return err
	}

	getData, err := pR.GetPostRedis(ctx)

	if err != nil {
		return err
	}

	var newPost []models.PostResp
	for _, v := range getData {
		if v.PostId != parseId {
			newPost = append(newPost, v)
		}
	}

	// Marshal back to JSON
	updatedData, err := json.Marshal(newPost)
	if err != nil {
		return err
	}

	// Save back to Redis
	err = pR.rdb.JSONSet(ctx, "postList", "$", updatedData).Err()
	if err != nil {
		return err
	}

	return nil
}
