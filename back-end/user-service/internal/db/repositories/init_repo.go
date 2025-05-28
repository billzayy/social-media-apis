package repositories

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type Repositories struct {
	UserRepository *UserRepository
}

func NewRepositories(db *sql.DB, rdb *redis.Client) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db, rdb),
	}
}
