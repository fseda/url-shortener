package config

import (
	"github.com/joho/godotenv"
)

func Config(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	values, err := godotenv.Read(".env")
	return values[key], err
}