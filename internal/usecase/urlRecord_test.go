package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ribeirosaimon/shortify/config/di"
	"github.com/ribeirosaimon/shortify/internal/dto"
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

	NewUrlRecord()

	ctx := context.Background()

	for _, v := range []struct {
		testName     string
		urlRecord    dto.UrlRecord
		errorMessage string
		hasError     bool
	}{
		{
			testName:  "Need pass",
			urlRecord: dto.UrlRecord{Url: "https://google.com"},
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
