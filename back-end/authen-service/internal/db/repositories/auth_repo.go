package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/billzayy/social-media/back-end/authen-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type AuthRepositoryInterface interface {
	AddUser(request models.RegisterRequest) error
	GetUserId(userName string, hashedPassword string) (string, error)
	GetHashedPassword(userName string) (string, error)
	CheckEmail(email string) (bool, error)
	RefreshToken(cookieToken string, c *gin.Context) (string, string, error)
}

type AuthRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewAuthRepository(db *sql.DB, rdb *redis.Client) *AuthRepository {
	return &AuthRepository{
		db:  db,
		rdb: rdb,
	}
}

// * Add User Function
func (ar *AuthRepository) AddUser(userName string, fullName string, email string, pass string) error {
	query := fmt.Sprintf(`INSERT INTO public."Users" 
	("UserName", "FullName", "Email", "Password", "LastLogin") 
	VALUES ('%s', '%s', '%s', '%s', '%s')`,
		userName, fullName, email, pass, time.DateTime)

	_, err := ar.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

// * GetUserId Function
func (ar *AuthRepository) GetUserId(userName string, hashedPassword string) (string, error) {
	query := fmt.Sprintf(`SELECT u."ID" FROM public."Users" u 
	WHERE u."UserName" = '%s' AND u."Password" = '%s'`, userName, hashedPassword)

	rows, err := ar.db.Query(query)

	if err != nil {
		return "", err
	}
	defer rows.Close()

	var userId string
	for rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return "", err
		}
	}

	return userId, nil
}

func (ar *AuthRepository) GetHashedPassword(userName string) (string, error) {
	query := fmt.Sprintf(`SELECT "Password" FROM public."Users" 
	WHERE "UserName" = '%s' OR "Email" = '%s'`, userName, userName)

	rows, err := ar.db.Query(query)

	defer rows.Close()

	if err != nil {
		return "", err
	}

	var hashedPass string
	for rows.Next() {
		err := rows.Scan(&hashedPass)

		if err != nil {
			return "", err
		}
	}

	return hashedPass, nil
}

func (ar *AuthRepository) CheckEmail(email string) (bool, error) {
	query := fmt.Sprintf(`SELECT exists(SELECT "Email" FROM public."Users" 
	WHERE "Email" = '%s' )`, email)

	rows, err := ar.db.Query(query)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	var checked bool
	for rows.Next() {
		err := rows.Scan(&checked)

		if err != nil {
			return false, err
		}
	}

	return checked, nil
}

func (ar *AuthRepository) SaveUserRedis(ctx context.Context, userId string, token string) error {
	_, err := ar.rdb.HSet(ctx, "loginList", userId, token).Result()

	if err != nil {
		return err
	}

	return nil
}
