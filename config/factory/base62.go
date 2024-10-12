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

type EntityBase62 struct {
	encodeType hash.Method
}

func NewBase62() *EntityBase62 {
	return &EntityBase62{
		encodeType: "base62",
	}
}

func (e *EntityBase62) Encode() vo.HashedUrl {

	rand.NewSource(time.Now().UnixNano())
	b := make([]rune, 5)

	// Preenche o slice com caracteres aleatórios
	for i := range b {
		b[i] = rune(base62Chars[rand.Intn(len(base62Chars))])
	}

	return vo.NewHashedUrl(string(b))
}

func (e *EntityBase62) GetEncodeType() hash.Method {
	return e.encodeType
}
