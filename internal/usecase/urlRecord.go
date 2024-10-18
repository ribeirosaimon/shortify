package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/ribeirosaimon/shortify/config/factory/hash"
	"github.com/ribeirosaimon/shortify/config/mediator"
	"github.com/ribeirosaimon/shortify/internal/dto"
	"github.com/ribeirosaimon/shortify/internal/entity"
	"github.com/ribeirosaimon/shortify/internal/vo"
)

type UrlRecord interface {
	Create(ctx context.Context, urlRecord *dto.UrlRecord) (*entity.UrlRecord, error)
}

type urlRecordUseCase struct {
}

func NewUrlRecord() *urlRecordUseCase {
	return &urlRecordUseCase{}
}

func (u *urlRecordUseCase) Create(ctx context.Context, url *dto.UrlRecord) (*entity.UrlRecord, error) {
	if url == nil {
		return nil, errors.New("urlRecord is nil")
	}

	if url.Url == "" {
		return nil, errors.New("urlRecord url is empty")
	}

	if !(strings.HasPrefix(url.Url, "http://") || strings.HasPrefix(url.Url, "https://")) {
		return nil, errors.New("is not a valid url")
	}
	urlRecord := entity.NewUrlRecord(vo.NewUrl(url.Url), hash.Base62)

	mediator.Get().Notify(mediator.PersistUrlRecord, &urlRecord)

	return &urlRecord, nil
}
