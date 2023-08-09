package config

import (
	"database/sql"
	// "fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func Config(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	values, err := godotenv.Read(".env")
	return values[key], err
}

func InitializeDB() {
	dbPath := os.Getenv("DB_PATH")
	// tableName := "urls"

	// _, err := os.Stat(dbPath)
	// if err == nil {
	// 	// File exists

	// 	// Open connection
	// 	db, err := sql.Open("sqlite3", "db.sqlite")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer db.Close()

	// 	createTableSQL := `
	// 	CREATE TABLE IF NOT EXISTS urls (
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		url TEXT
	// 	)
	// 	`

	// 	_, err = db.Exec(createTableSQL)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer db.Close()

	// } else if os.IsNotExist(err) {
	// 	fmt.Println("File does not exist")
	// } else {
	// 	fmt.Println("Error:", err)
	// }

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT
	)
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
