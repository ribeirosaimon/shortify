package adapter

import (
	"github.com/ribeirosaimon/shortify/config/di"
	"github.com/ribeirosaimon/shortify/config/mediator"
	"github.com/ribeirosaimon/shortify/internal/cache"
	"github.com/ribeirosaimon/shortify/internal/repository"
	"github.com/ribeirosaimon/shortify/internal/usecase"
)

func NewExpress() {
	di.GetRegistry().Provide(di.UrlRecordRepository, func() any {
		return repository.NewUrl()
	})

	di.GetRegistry().Provide(di.UrlRecordUseCase, func() any {
		return usecase.NewUrlRecord()
	})

	di.GetRegistry().Provide(di.UrlRecordCache, func() any {
		return cache.NewUrlRecord()
	})

	mediator.PersistUrl()
}
