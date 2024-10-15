package cache

import (
	"context"

	"github.com/ribeirosaimon/shortify/internal/dto"
	"github.com/ribeirosaimon/tooltip/storage/redis"
)

type UrlRecord interface {
	Create(ctx context.Context, urlRecord *dto.UrlRecord) error
}

type urlRecordCache struct {
	urlCache redis.Connection
}

func NewUrlRecordCache() *urlRecordCache {
	return &urlRecordCache{
		urlCache: redis.NewRedisConnection(),
	}
}

func (u *urlRecordCache) Create(ctx context.Context, urlRecord *dto.UrlRecord) error {
	u.urlCache.
}
