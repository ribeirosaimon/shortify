package cache

import (
	"context"

	"github.com/ribeirosaimon/shortify/config/db"
	"github.com/ribeirosaimon/shortify/config/factory"
	"github.com/ribeirosaimon/shortify/internal/entity"
	"github.com/ribeirosaimon/shortify/internal/role"
	"github.com/ribeirosaimon/tooltip/storage/redis"
	"github.com/ribeirosaimon/tooltip/tlog"
)

type UrlRecord interface {
	Create(ctx context.Context, urlRecord *entity.UrlRecord) error
}

type urlRecordCache struct {
	urlCache redis.RConnInterface
}

func NewUrlRecord() *urlRecordCache {
	return &urlRecordCache{
		urlCache: db.NewRedisConnection(),
	}
}

func (u *urlRecordCache) Create(ctx context.Context, urlRecord *entity.UrlRecord) error {
	tlog.Debug("NewUrlRecord.Create", "Create Url in Redis")

	plan := factory.GetUserPlan(role.BasicUser)
	if err := u.urlCache.GetConnection().
		Set(
			ctx,
			urlRecord.GetShortenedUrl().GetValue(),
			urlRecord.GetOriginalUrl().GetValue(),
			plan.TimeToExpired()).
		Err(); err != nil {
		tlog.Warn("NewUrlRecord.Create", "Error when create Url in Redis")
	}
	return nil
}
