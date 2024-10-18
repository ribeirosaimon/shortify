package mediator

import (
	"context"

	"github.com/ribeirosaimon/shortify/config/di"
	"github.com/ribeirosaimon/shortify/internal/cache"
	"github.com/ribeirosaimon/shortify/internal/entity"
	"github.com/ribeirosaimon/shortify/internal/repository"
)

func PersistUrl() {
	Get().Register(PersistUrlRecord, func(T any) error {
		record := di.GetRegistry().Inject(di.UrlRecordRepository).(repository.UrlRecordRepository)
		recordCache := di.GetRegistry().Inject(di.UrlRecordCache).(cache.UrlRecord)

		persistUrlRecordContext := context.Background()
		if err := recordCache.Create(persistUrlRecordContext, T.(*entity.UrlRecord)); err != nil {
			return err
		}
		_, err := record.InsertUrlRecord(persistUrlRecordContext, T.(*entity.UrlRecord))
		if err != nil {
			return err
		}
		return nil
	})
}
