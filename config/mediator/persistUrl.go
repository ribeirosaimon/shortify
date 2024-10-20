package mediator

import (
	"context"

	"github.com/ribeirosaimon/shortify/internal/cache"
	"github.com/ribeirosaimon/shortify/internal/entity"
	"github.com/ribeirosaimon/shortify/internal/repository"
)

type PersistUrlMediator struct {
	urRepository repository.UrlRecordRepository
	urlCache     cache.UrlRecord
}

func NewPersistUrlMediator(recordRepository repository.UrlRecordRepository, record cache.UrlRecord) *PersistUrlMediator {
	return &PersistUrlMediator{
		urRepository: recordRepository,
		urlCache:     record,
	}
}

func (p *PersistUrlMediator) Notify(ctx context.Context, T any) error {
	if err := p.urlCache.Create(ctx, T.(*entity.UrlRecord)); err != nil {
		return err
	}
	_, err := p.urRepository.InsertUrlRecord(ctx, T.(*entity.UrlRecord))
	if err != nil {
		return err
	}
	return nil
}
