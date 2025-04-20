package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/billzayy/social-media/back-end/authen-service/internal/middleware"
	"github.com/billzayy/social-media/back-end/authen-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type AuthRepositoryInterface interface {
	Register(request models.RegisterRequest) error
	Login(userName string, hashedPassword string, c *gin.Context) (models.UserToken, error)
	CheckUserWithNameAndPass(userName string, password string) (string, bool, error)
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

// Register Function
func (ur *AuthRepository) Register(request models.RegisterRequest) error {
	fullName := request.FirstName + request.Surname

	hashedPass, err := middleware.HashPassword(request.Password)

	if err != nil {
		return err
	}

	query := fmt.Sprintf(`INSERT INTO public."Users" 
	("UserName", "FullName", "Email", "Password", "LastLogin") 
	VALUES ('%s', '%s', '%s', '%s', '%s')`,
		request.UserName, fullName, request.Email, hashedPass, time.DateTime)

	checkedEmail, err := ur.CheckEmail(request.Email)

	if err != nil {
		return err
	}

	if checkedEmail {
		return fmt.Errorf("account existed!")
	}

	_, err = ur.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

// Login Function
func (ur *AuthRepository) Login(userName string, hashedPassword string) (models.UserToken, http.Cookie, error) {
	ctx := context.Background()

	query := fmt.Sprintf(`SELECT u."ID" FROM public."Users" u 
	WHERE u."UserName" = '%s' AND u."Password" = '%s'`, userName, hashedPassword)

	rows, err := ur.db.Query(query)

	if err != nil {
		return models.UserToken{}, http.Cookie{}, err
	}
	defer rows.Close()

	var userId string
	for rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return models.UserToken{}, http.Cookie{}, err
		}
	}

	accessToken, cookie, err := middleware.GenerateTokens(userId)

	if err != nil {
		return models.UserToken{}, http.Cookie{}, err
	}

	_, err = ur.rdb.HSet(ctx, "loginList", userId, accessToken).Result()

	if err != nil {
		panic(err)
	}

	return models.UserToken{UserId: uuid.MustParse(userId), Token: accessToken, Type: "Bearer"}, cookie, nil
}

// Check User Function
func (ur *AuthRepository) CheckUserWithNameAndPass(userName string, password string) (string, bool, error) {
	query := fmt.Sprintf(`SELECT "Password" FROM public."Users" 
	WHERE "UserName" = '%s' OR "Email" = '%s'`, userName, userName)

	rows, err := ur.db.Query(query)

	defer rows.Close()

	if err != nil {
		return "", false, err
	}

	var hashedPass string
	for rows.Next() {
		err := rows.Scan(&hashedPass)

		if err != nil {
			return "", false, err
		}
	}

	res, err := middleware.ValidatePassword(password, hashedPass)

	if err != nil {
		return "", res, err
	}

	return hashedPass, res, nil
}

func (ur *AuthRepository) CheckEmail(email string) (bool, error) {
	query := fmt.Sprintf(`SELECT exists(SELECT "Email" FROM public."Users" 
	WHERE "Email" = '%s' )`, email)

	rows, err := ur.db.Query(query)

	if err != nil {
		return false, err
	}

	var checked bool
	for rows.Next() {
		err := rows.Scan(&checked)

		if err != nil {
			return false, err
		}
	}

	return checked, nil
}

func (ur *AuthRepository) RefreshToken(cookieToken string) (string, string, http.Cookie, error) {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		err := godotenv.Load("../internal/.env")

		if err != nil {
			fmt.Println("Error loading file .env to refresh Token")
			return "", "", http.Cookie{}, err
		}
	}

	// Check the refresh token is valid or not
	userId, err := middleware.VerifyRefreshToken(cookieToken)

	if err != nil {
		return "", "", http.Cookie{}, err
	}

	// Create new access & refresh token
	newToken, cookie, err := middleware.GenerateTokens(userId)

	if err != nil {
		return "", "", http.Cookie{}, err
	}

	return userId, newToken, cookie, nil
}
