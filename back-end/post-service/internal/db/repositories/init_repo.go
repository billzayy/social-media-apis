package repositories

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type Repositories struct {
	PostRepository     *PostRepository
	InteractRepository *InteractRepository
}

func NewRepositories(db *sql.DB, rdb *redis.Client) *Repositories {
	return &Repositories{
		PostRepository:     NewPostRepository(db, rdb),
		InteractRepository: NewInteractRepository(db, rdb),
	}
}
