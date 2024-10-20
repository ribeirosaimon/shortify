package server

import (
	"github.com/ribeirosaimon/shortify/config/mediator"
	"github.com/ribeirosaimon/shortify/internal/cache"
	"github.com/ribeirosaimon/shortify/internal/repository"
	"github.com/ribeirosaimon/shortify/internal/usecase"
)

// Services application
// is nice only apply interfaces
type services struct {
	urlRepository      repository.UrlRecordRepository
	urlUseCase         usecase.UrlRecord
	urlCache           cache.UrlRecord
	urlPersistMediator mediator.Handler
}

// appServerService is my app server
// This variable cannot be null, it is where all the application services are found
// It makes testing easier, where I just need to map out what is needed
var appServerService = &services{}

// Option is the function with handler my servicees
type Option func(*services)

// NewServices
// is the only place where I can initialize the variable appServerService
// because this my services was private
func NewServices(opts ...Option) *services {
	if appServerService == nil {
		appServerService = &services{}
	}

	for _, opt := range opts {
		opt(appServerService)
	}
	return appServerService
}

// functions to get a mapped service

func GetUrlRecordRepository() repository.UrlRecordRepository {
	return appServerService.urlRepository
}

func GetUrlRecordUsecase() usecase.UrlRecord {
	return appServerService.urlUseCase
}
func GetUrlRecordCache() cache.UrlRecord {
	return appServerService.urlCache
}
func GetPersistMediator() mediator.Handler {
	return appServerService.urlPersistMediator
}
