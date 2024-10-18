package factory

import (
	"math/rand"
	"time"

	"github.com/ribeirosaimon/shortify/config/factory/hash"
	"github.com/ribeirosaimon/shortify/internal/vo"
)

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type entityBase62 struct {
	encodeType hash.Method
}

func NewBase62() entityBase62 {
	return entityBase62{
		encodeType: "base62",
	}
}

func (e entityBase62) Encode() vo.HashedUrl {
	rand.NewSource(time.Now().UnixNano())
	b := make([]rune, 5)
	for i := range b {
		b[i] = rune(base62Chars[rand.Intn(len(base62Chars))])
	}

	return vo.NewHashedUrl(string(b))
}

func (e entityBase62) GetEncodeType() hash.Method {
	return e.encodeType
}
