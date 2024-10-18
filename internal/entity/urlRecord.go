package entity

import (
	"time"

	"github.com/ribeirosaimon/shortify/config/factory"
	"github.com/ribeirosaimon/shortify/config/factory/hash"
	"github.com/ribeirosaimon/shortify/internal/vo"
)

type UrlRecord struct {
	createdAt    time.Time
	id           vo.Id
	originalUrl  vo.Url
	shortenedUrl vo.HashedUrl
}

func NewUrlRecord(originalUrl vo.Url, hashMethod hash.Method) UrlRecord {
	return UrlRecord{
		id:           vo.NewId(),
		createdAt:    time.Now(),
		originalUrl:  originalUrl,
		shortenedUrl: factory.EncriptByHashMethod(hashMethod).Encode(),
	}
}

func TranformInUrlRecord(id, original, shortened string, timestamp time.Time) *UrlRecord {
	return &UrlRecord{
		id:           vo.TransformStringInId(id),
		createdAt:    timestamp,
		originalUrl:  vo.NewUrl(original),
		shortenedUrl: vo.NewHashedUrl(shortened),
	}
}

func (u *UrlRecord) GetShortenedUrl() vo.HashedUrl {
	return u.shortenedUrl
}

func (u *UrlRecord) GetOriginalUrl() vo.Url {
	return u.originalUrl
}

func (u *UrlRecord) GetId() vo.Id {
	return u.id
}

func (u *UrlRecord) GetCreatedAt() time.Time {
	return u.createdAt
}
