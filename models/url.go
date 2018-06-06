package models

type URL struct {
	OriginalURL string `form:"original_url"`
	ShortURL    string `form:"short_url"`
}

type UrlManager struct {
	URLs []*URL
}

func (o *UrlManager) Save(url *URL) {
	o.URLs[len(o.URLs)] = url
	return
}
