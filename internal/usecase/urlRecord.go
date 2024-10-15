package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/ribeirosaimon/shortify/config/di"
	"github.com/ribeirosaimon/shortify/config/factory/hash"
	"github.com/ribeirosaimon/shortify/internal/dto"
	"github.com/ribeirosaimon/shortify/internal/entity"
	"github.com/ribeirosaimon/shortify/internal/repository"
	"github.com/ribeirosaimon/shortify/internal/vo"
)

type UrlRecord interface {
	Create(ctx context.Context, urlRecord *dto.UrlRecord) (*entity.UrlRecord, error)
}

type urlRecordUseCase struct {
	urlRepository repository.UrlRecordRepository
}

func NewUrlRecord() *urlRecordUseCase {
	recordRepository := di.GetRegistry().Inject(di.UrlRecordRepository).(repository.UrlRecordRepository)
	return &urlRecordUseCase{
		urlRepository: recordRepository,
	}
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

	record, err := u.urlRepository.InsertUrlRecord(ctx, &urlRecord)
	if err != nil {
		return nil, err
	}

	return record, nil
}
