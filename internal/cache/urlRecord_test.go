package cache

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

func TestUrlCache(t *testing.T) {
	ctx := context.Background()
	container := tcontainer.NewRedisTestContainer()
	err := container.Start()
	assert.Nil(t, err)

	tserver.NewMockEnvironment(tserver.MockEnvironment{RedisHost: container.GetHost(), RedisDatabase: "test"})
	service := NewUrlRecord()

	for _, v := range []struct {
		testName string
		entity   entity.UrlRecord
		hasErr   bool
	}{
		{
			testName: "have to pass",
			entity:   entity.NewUrlRecord(vo.NewUrl("https://google.com"), hash.Base62),
		},
	} {
		t.Run(v.testName, func(t *testing.T) {
			err = service.Create(ctx, &v.entity)

			if v.hasErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
