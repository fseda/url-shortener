package config

import (
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB      *sql.DB
	DB_PATH = "src/db/db.sqlite"
)

func Config(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	values, err := godotenv.Read(".env")
	return values[key], err
}

func InitializeDB() error {
	var err error

	DB, err = sql.Open("sqlite3", DB_PATH)
	if err != nil {
		return err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT
	)
	`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}
