package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/billzayy/social-media/authen-service/internal/models"
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
func (ar *AuthRepository) AddUser(userName string, fullName string, email string, pass string) (string, error) {
	var id string

	query := `INSERT INTO public."Users" ("UserName", "FullName", "Email", "Password", "LastLogin") 
	VALUES ($1,$2,$3,$4,$5) RETURNING "ID"`

	rows, err := ar.db.Query(query, userName, fullName, email, pass, time.Now())

	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)

		if err != nil {
			return "", err
		}
	}

	return id, nil
}

func (ar *AuthRepository) AddUserProfile(id string) error {
	query := `INSERT INTO public."UserProfile" ("UserId") VALUES ($1)`

	_, err := ar.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (ar *AuthRepository) AddUserSettings(id string) error {
	query := `INSERT INTO public."UserSettings" ("UserId") VALUES ($1)`

	_, err := ar.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (ar *AuthRepository) GetUser(userName string, hashedPassword string) (models.GetUserData, error) {
	query := `SELECT u."ID", u."FullName", u."Email", u."ProfilePicture" FROM public."Users" u 
	WHERE u."UserName" = $1 AND u."Password" = $2`

	rows, err := ar.db.Query(query, userName, hashedPassword)

	if err != nil {
		return models.GetUserData{}, err
	}
	defer rows.Close()

	var data models.GetUserData
	var profilePicture sql.NullString

	for rows.Next() {
		err = rows.Scan(&data.ID, &data.FullName, &data.Email, &profilePicture)
		if err != nil {
			return models.GetUserData{}, err
		}

		if profilePicture.Valid {
			data.ProfilePicture = profilePicture.String
		}
	}

	return data, nil
}

func (ar *AuthRepository) GetUserById(id string) (models.GetUserData, error) {
	query := `SELECT u."ID", u."FullName", u."Email", u."ProfilePicture" FROM public."Users" u 
	WHERE u."ID" = $1`

	rows, err := ar.db.Query(query, id)

	if err != nil {
		return models.GetUserData{}, err
	}
	defer rows.Close()

	var data models.GetUserData
	var profilePicture sql.NullString

	for rows.Next() {
		err = rows.Scan(&data.ID, &data.FullName, &data.Email, &profilePicture)
		if err != nil {
			return models.GetUserData{}, err
		}

		if profilePicture.Valid {
			data.ProfilePicture = profilePicture.String
		}
	}

	return data, nil
}

func (ar *AuthRepository) UpdateLoginTime(id string) error {
	query := `UPDATE public."Users" SET "LastLogin" = $1 WHERE "ID" = $2`

	timeNow := time.Now()

	_, err := ar.db.Exec(query, timeNow, id)

	if err != nil {
		return err
	}

	return nil
}

func (ar *AuthRepository) GetHashedPassword(userName string) (string, error) {
	query := `SELECT "Password" FROM public."Users" WHERE "UserName" = $1 OR "Email" = $2`

	rows, err := ar.db.Query(query, userName, userName)

	if err != nil {
		return "", err
	}
	defer rows.Close()

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
