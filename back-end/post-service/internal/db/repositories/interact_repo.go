package repositories

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type InteractRepositoryInterface interface {
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

func (iR *InteractRepository) AddLike() {}
