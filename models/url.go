package models

import "fmt"

type URL struct {
	originalURL string
	clickCount  int64
	shortURL    string
}

var DefaultURLMap *URLManager

func NewURL(originalURL string) (*URL, error) {
	if originalURL == "" {
		return nil, fmt.Errorf("empty title")
	}
	return &URL{originalURL, 0}, nil
}

type URLManager struct {
	urlMap map[string]*URL
	Count  int64
}

func NewUrlManager() *URLManager {
	DefaultURLMap = &URLManager{
		Count:  10000,
		urlMap: new(map[string]*URL),
	}
	return DefaultURLMap
}

////////internal functions
func (m *URLManager) AddURL(url *URL) error {

	if m.Count == 0 {
		return nil
	}

	for k, v := range m.urlMap {

	}
	return fmt.Errorf("unknown task")
}

func cloneURL(t *URL) *URL {
	c := *t
	return &c
}
