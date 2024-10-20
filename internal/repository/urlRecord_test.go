package repository

import (
	"context"
	"testing"

	"github.com/ribeirosaimon/shortify/config/factory/hash"
	"github.com/ribeirosaimon/shortify/internal/entity"
	"github.com/ribeirosaimon/shortify/internal/vo"
	"github.com/ribeirosaimon/tooltip/testutils/tcontainer"
	"github.com/ribeirosaimon/tooltip/tserver"
	"github.com/stretchr/testify/assert"
)

func TestUrlRecordRepository(t *testing.T) {
	ctx := context.Background()
	container := tcontainer.NewPgsqlTestContainer()
	err := container.Pgsql()
	assert.Nil(t, err)

	tserver.NewMockEnvironment(tserver.MockEnvironment{PgsqlHost: container.GetHost(), PgsqlDatabase: "test", PgsqlEntryPoint: "../../init.sql"})

	repository := NewUrl()
	for _, v := range []struct {
		testName        string
		entity          entity.UrlRecord
		wantOriginalUrl string
		hasError        bool
	}{
		{
			testName:        "have to persist url in postgres",
			entity:          entity.NewUrlRecord(vo.NewUrl("https://google.com"), hash.Base62),
			wantOriginalUrl: "https://google.com",
		},
	} {
		t.Run(v.testName, func(t *testing.T) {
			record, err := repository.InsertUrlRecord(ctx, &v.entity)
			assert.Nil(t, err)
			assert.Equal(t, record.GetOriginalUrl().GetValue(), v.wantOriginalUrl)

			urlRecord, err := repository.GetUrlRecord(ctx, record.GetId().GetValue())
			assert.Nil(t, err)

			assert.Equal(t, urlRecord.GetOriginalUrl().GetValue(), record.GetOriginalUrl().GetValue())
			assert.Equal(t, urlRecord.GetId().GetValue(), record.GetId().GetValue())
			assert.Equal(t, urlRecord.GetShortenedUrl().GetValue(), record.GetShortenedUrl().GetValue())

		})
	}

}
