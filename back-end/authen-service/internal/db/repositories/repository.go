package repositories

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type Repositories struct {
	AuthRepository *AuthRepository
}

func NewRepositories(db *sql.DB, rdb *redis.Client) *Repositories {
	return &Repositories{
		AuthRepository: NewAuthRepository(db, rdb),
	}
}
