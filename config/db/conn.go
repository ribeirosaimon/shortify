package db

import (
	"context"

	"github.com/ribeirosaimon/tooltip/storage/mongo"
	"github.com/ribeirosaimon/tooltip/storage/pgsql"
	"github.com/ribeirosaimon/tooltip/storage/redis"
	"github.com/ribeirosaimon/tooltip/tlog"
	"github.com/ribeirosaimon/tooltip/tserver"
)

func NewMongoConnection(ctx context.Context) mongo.MConnInterface {
	config := tserver.GetMongoConfig()
	tlog.Warn("NewMongoConnection", "Url: ", config.Host)
	return mongo.NewConnMongo(ctx,
		mongo.WithUrl(config.Host),
		mongo.WithDatabase(config.Database),
	)
}

func NewPgsqlConnection() pgsql.PConnInterface {
	config := tserver.GetPgsqlConfig()
	tlog.Warn("NewPgsqlConnection", "Url: ", config.Host)
	return pgsql.NewConnPgsql(
		pgsql.WithUrl(config.Host),
	)
}

func NewRedisConnection() redis.RConnInterface {
	config := tserver.GetRedisConfig()
	tlog.Warn("NewRedisConnection", "Url: ", config.Host)
	return redis.NewRedisConnection(
		redis.WithUrl(config.Host),
	)
}
