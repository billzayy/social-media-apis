package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load("./.env")

	if err != nil {
		if err := godotenv.Load("./internal/.env"); err != nil {
			log.Fatalf("Error loading .env on Postgres")
		}
	}

	// Connection string
	USERNAME := os.Getenv("POSTGRES_USERNAME")
	PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	HOST := os.Getenv("POSTGRES_HOST")
	PORT := os.Getenv("POSTGRES_PORT")
	DATABASE := os.Getenv("POSTGRES_DATABASE")

	connectStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", USERNAME, PASSWORD, HOST, PORT, DATABASE)

	// Connect to the database
	conn, err := sql.Open("postgres", connectStr)

	if err != nil {
		panic(err)
	}

	return conn, err
}
