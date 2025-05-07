package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/billzayy/social-media/back-end/post-service/internal/models"
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
