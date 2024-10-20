package server

import (
	"github.com/ribeirosaimon/shortify/internal/cache"
	"github.com/ribeirosaimon/shortify/internal/repository"
	"github.com/ribeirosaimon/shortify/internal/usecase"
)

func WithUrlRepository(urlRepository repository.UrlRecordRepository) Option {
	return func(s *services) {
		s.urlRepository = urlRepository
	}
}

func WithUrlUseCase(urlUseCase usecase.UrlRecord) Option {
	return func(s *services) {
		s.urlUseCase = urlUseCase
	}
}

func WithUrlCache(urlCache cache.UrlRecord) Option {
	return func(s *services) {
		s.urlCache = urlCache
	}
}
