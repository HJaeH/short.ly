package url_organizer

import (
	"net/url"
	"strings"
)

func GetUniqueURL(originalURL string) string {
	strArr1 := strings.Split(originalURL, "://")
	if len(strArr1) < 1 {
		return originalURL
	}
	if strArr1[0] != "http" && strArr1[0] != "https" {
		originalURL = "https://" + originalURL
	}

	urlObject, err := url.Parse(originalURL)
	if err != nil {
		return originalURL
	}

	host := urlObject.Hostname()

	strArr := strings.Split(urlObject.Hostname(), ".")
	if len(strArr) < 1 {
		return originalURL
	}

	if strArr[0] != "www" {
		host = "www." + host
	}
	if urlObject.RawQuery != "" {
		return "https://"+ host + "?" + urlObject.RawQuery
	}
	return "https://"+host
}
