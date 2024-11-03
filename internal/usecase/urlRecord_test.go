package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mediatorMock "github.com/ribeirosaimon/shortify/config/mediator/mocks"
	"github.com/ribeirosaimon/shortify/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestUrlRecordUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mediator := mediatorMock.NewMockHandler(ctrl)
	recordUseCase := NewUrlRecord(mediator)

	ctx := context.Background()

	for _, v := range []struct {
		auxFunc      func()
		urlRecord    dto.UrlRecord
		testName     string
		errorMessage string
		hasError     bool
	}{
		{
			testName:  "Need pass",
			urlRecord: dto.UrlRecord{Url: "https://google.com"},
			auxFunc: func() {
				mediator.EXPECT().Notify(gomock.Any(), gomock.Any()).Return(nil)
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
