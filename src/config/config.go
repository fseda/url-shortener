package config

import (
	"database/sql"
	// "fmt"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Config(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	values, err := godotenv.Read(".env")
	return values[key], err
}

func InitializeDB() {
	var err error
	dbPath, _ := Config("DB_PATH")

	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT
	)
	`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	defer DB.Close()
}
