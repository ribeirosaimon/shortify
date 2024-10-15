package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ribeirosaimon/shortify/config/di"
	"github.com/ribeirosaimon/shortify/internal/dto"
	"github.com/ribeirosaimon/shortify/internal/entity"
	repoMock "github.com/ribeirosaimon/shortify/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUrlRecordUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	di.GetRegistry().Provide(di.UrlRecordRepository, func() any {
		return repoMock.NewMockUrlRecordRepository(ctrl)
	})
	di.GetRegistry().Provide(di.UrlRecordUseCase, func() any {
		return NewUrlRecord()
	})

	recordRepository := di.GetRegistry().Inject(di.UrlRecordRepository).(*repoMock.MockUrlRecordRepository)
	NewUrlRecord()

	ctx := context.Background()

	for _, v := range []struct {
		testName     string
		auxFunc      func()
		urlRecord    dto.UrlRecord
		errorMessage string
		hasError     bool
	}{
		{
			testName:  "Need pass",
			urlRecord: dto.UrlRecord{Url: "https://google.com"},
			auxFunc: func() {
				recordRepository.EXPECT().InsertUrlRecord(gomock.Any(), gomock.Any()).
					Return(entity.TranformInUrlRecord("", "https://google.com", "", time.Now()), nil)
			},
		},
		{
			testName:     "Can't pass because is a no valid url",
			urlRecord:    dto.UrlRecord{Url: "httttp://google.com"},
			hasError:     true,
			errorMessage: "is not a valid url",
		},
		{
			testName:     "Can't pass because empty url",
			urlRecord:    dto.UrlRecord{Url: ""},
			hasError:     true,
			errorMessage: "urlRecord url is empty",
		},
		{
			testName:     "Can't pass because empty url",
			hasError:     true,
			errorMessage: "urlRecord is nil",
		},
	} {
		t.Run(v.testName, func(t *testing.T) {
			if v.auxFunc != nil {
				v.auxFunc()
			}
			recordUseCase := di.GetRegistry().Inject(di.UrlRecordUseCase).(UrlRecord)
			recordDb, err := recordUseCase.Create(ctx, &v.urlRecord)
			if v.hasError {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.NotEmpty(t, recordDb)
			assert.Equal(t, recordDb.GetOriginalUrl().GetValue(), v.urlRecord.Url)
		})
	}
}
