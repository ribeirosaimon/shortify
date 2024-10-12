package entity

import (
	"testing"

	"github.com/ribeirosaimon/shortify/config/factory/hash"
	"github.com/ribeirosaimon/shortify/internal/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewUrlRecord(t *testing.T) {
	for _, v := range []struct {
		testName string
		url      string
		hasError bool
	}{
		{
			testName: "have to create new urlRecord",
			url:      "https://example.com",
		},
		{
			testName: "have to create new realy big urlRecord",
			url:      "https://example.com/5b1efef9-6f45-417d-a65a-6a412c905437/5b1efef9-6f45-417d-a65a-6a412c905437/5b1efef9-6f45-417d-a65a-6a412c905437",
		},
		{
			testName: "error create url record",
			url:      "",
			hasError: true,
		},
		{
			testName: "error create url record because is not http",
			url:      "ht://example.com",
			hasError: true,
		},
	} {
		t.Run(v.testName, func(t *testing.T) {
			record := NewUrlRecord(vo.NewUrl(v.url), hash.Base62)
			url := record.GetOriginalUrl()
			if v.hasError {
				assert.True(t, url.GetValue() == "")
				return
			}
			assert.NotNil(t, record)
			assert.NotNil(t, record.id)
			assert.Equal(t, url.GetValue(), v.url)
		})
	}

}
