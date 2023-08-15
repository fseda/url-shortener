package services

import (
	"database/sql"
	"strconv"

	"github.com/fseda/url-shortener/src/config"
)

type UrlService struct{}

func NewUrlService() *UrlService {
	return &UrlService{}
}

func (us *UrlService) GetUrl(id int) (string, error) {
	query := "SELECT url FROM urls WHERE id = ?"
	row := config.DB.QueryRow(query, id)

	var url string
	err := row.Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
	}

	return url, nil
}

func (us *UrlService) SaveUrl(url string) (int64, error) {
	query := "INSERT INTO urls (url) VALUES (?)"
	result, err := config.DB.Exec(query, url)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (us *UrlService) UrlExists(url string) (bool, error) {
	query := "SELECT * FROM urls WHERE url = ?"
	result, err := config.DB.Exec(query, url)
	if err != nil {
		return false, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (us *UrlService) GetUrlId(url string) (int64, error) {
	query := "SELECT * FROM urls WHERE url = ?"
	result, err := config.DB.Exec(query, url)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (us *UrlService) ShortenUrl(url string, baseUrl string) (string, error) {
	var id int64
	var err error
	
	if urlExists, _ := us.UrlExists(url); urlExists {
		id, err = us.GetUrlId(url)
		if err != nil {
			return "", err
		}
	} else {
		id, err = us.SaveUrl(url)
		if err != nil {
			return "", err
		}
	}

	return baseUrl + "?id=" + strconv.FormatInt(id, 10), nil
}
