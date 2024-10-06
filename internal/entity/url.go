package entity

import "time"

type UrlRecord struct {
	createdAt    time.Time
	id           Id
	originalUrl  Url
	shortenedUrl Url
}

func NewUrlRecord(originalUrl Url, shortenedUrl Url) UrlRecord {
	return UrlRecord{
		id:           NewId(),
		createdAt:    time.Now(),
		originalUrl:  originalUrl,
		shortenedUrl: shortenedUrl,
	}
}

type Url struct {
	url string
}

func NewUrl(url string) Url {
	return Url{
		url: url,
	}
}
