package repositories

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/billzayy/social-media/back-end/user-service/internal/models"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type UserRepositoryInterface interface {
	GetUserById(userId string) (models.ResponseUser, error)
	UpdateInfoUser(req models.ReqUpdateUser) error
	GetPassword(id string) (string, error)
	UpdatePassword(id string, newPassword string) error
	GetUserRedis(ctx context.Context, userId string) (models.ResponseUser, error)
	AddUserRedis(ctx context.Context, userId string, data models.ResponseUser) error
}

type UserRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewUserRepository(db *sql.DB, rdb *redis.Client) *UserRepository {
	return &UserRepository{
		db:  db,
		rdb: rdb,
	}
}

func (ur *UserRepository) GetUserById(userId string) (models.ResponseUser, error) {
	rows, err := ur.db.Query(getUserByIdQuery, userId)

	if err != nil {
		return models.ResponseUser{}, err
	}

	var result models.ResponseUser
	var profilePicture, description, location, theme, language, country sql.NullString
	var website pq.StringArray
	var birthDate sql.NullTime

	for rows.Next() {
		err := rows.Scan(
			&result.ID, &result.UserName, &result.FullName, &result.Email,
			&result.UserSince, &result.LastLogin, &profilePicture, &result.Active,
			&description, &location, &website, &birthDate,
			&theme, &language, &country,
		)

		if err != nil {
			return models.ResponseUser{}, err
		}

		if profilePicture.Valid {
			result.ProfilePicture = profilePicture.String
		}
		if description.Valid {
			result.Description = description.String
		}

		if birthDate.Valid {
			result.BirthDate = birthDate.Time
		}
		if location.Valid {
			result.Location = location.String
		}
		if theme.Valid {
			result.Theme = theme.String
		}
		if language.Valid {
			result.Language = language.String
		}
		if country.Valid {
			result.Country = country.String
		}
		if len(website) > 0 {
			result.Website = website
		}
	}

	return result, nil
}

func (ur *UserRepository) UpdateInfoUser(req models.ReqUpdateUser) error {
	tx, err := ur.db.Begin() // Begin SQL Transactions
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // error case
		} else {
			err = tx.Commit() // success case
		}
	}()

	// Update Users table
	_, err = tx.Exec(updateUserTable, req.UserName, req.Email, req.FullName, req.ProfilePicture, req.ID)
	if err != nil {
		return err
	}

	// Update UserProfile table
	_, err = tx.Exec(updateUserProfile, req.Description, req.Location, req.Website, req.BirthDate, req.ID)
	if err != nil {
		return err
	}

	// Update UserSettings table
	_, err = tx.Exec(updateUserSetting, req.Theme, req.Language, req.Country, req.ID)
	if err != nil {
		return err
	}

	return err
}

func (ur *UserRepository) GetPassword(id string) (string, error) {
	rows, err := ur.db.Query(getPasswordQuery, id)

	if err != nil {
		return "", err
	}

	var result string

	for rows.Next() {
		err := rows.Scan(&result)

		if err != nil {
			return result, err
		}
	}

	return result, nil
}

func (ur *UserRepository) UpdatePassword(id string, newPassword string) error {
	_, err := ur.db.Exec(updatePassword, newPassword, id)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserRedis(ctx context.Context, userId string) (models.ResponseUser, error) {
	var user models.ResponseUser

	val, err := ur.rdb.HGet(ctx, "userList", userId).Result()

	if err != nil {
		if err == redis.Nil {
			return user, nil // or return an error if you want to handle "not found"
		}
		return user, err
	}

	err = json.Unmarshal([]byte(val), &user) // Convert json type into models.ResponseUser type
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) AddUserRedis(ctx context.Context, userId string, data models.ResponseUser) error {
	jsonData, err := json.Marshal(data) // Convert data into Json type

	if err != nil {
		return err
	}
	_, err = ur.rdb.HSet(ctx, "userList", userId, jsonData).Result()

	if err != nil {
		return err
	}

	return nil
}
