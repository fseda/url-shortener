package models

type Url struct {
	ID           int64  `json:"id"`
	ShortenedUrl string `json:"shortened_url"`
	OriginalUrl  string `json:"original_url"`
}
