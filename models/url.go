package models

type URL struct {
	OriginalURL string `form:"original_url"`
	ShortURL    string `form:"short_url"`
}
