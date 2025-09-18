package repositories

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type Repositories struct {
	NotifyRepository *NotifyRepository
}

func NewRepositories(db *sql.DB, rdb *redis.Client) *Repositories {
	return &Repositories{
		NotifyRepository: NewNotifyRepository(db, rdb),
	}
}
