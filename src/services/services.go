package services

import (
	"database/sql"
	"fmt"
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

func (us *UrlService) saveUrl(url string) (int64, error) {
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

func (us *UrlService) urlExists(url string) bool {
	query := "SELECT id FROM urls WHERE url = ?"
	row := config.DB.QueryRow(query, url)

	var id int
	err := row.Scan(&id)
	return err != sql.ErrNoRows
}

func (us *UrlService) getUrlId(url string) (int64, error) {
	query := "SELECT id FROM urls WHERE url = ?"
	row := config.DB.QueryRow(query, url)

	var id int
	err := row.Scan(&id)
	return int64(id), err
}

func (us *UrlService) ShortenUrl(url string, baseUrl string) (string, error) {
	var id int64
	var err error

	urlExists := us.urlExists(url)
	if urlExists {
		id, err = us.getUrlId(url)
		if err != nil {
			return "", err
		}
	} else {
		id, err = us.saveUrl(url)
		if err != nil {
			return "", err
		}
	}

	return baseUrl + "?id=" + strconv.FormatInt(id, 10), nil
}
