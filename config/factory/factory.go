package factory

import (
	"github.com/ribeirosaimon/shortify/config/factory/hash"
	"github.com/ribeirosaimon/shortify/internal/vo"
)

type EncodeUrlFactory interface {
	Encode() vo.HashedUrl
	GetEncodeType() hash.Method
}

func EncriptByHashMethod(method hash.Method) vo.HashedUrl {
	return map[hash.Method]vo.HashedUrl{
		hash.Base62: NewBase62().Encode(),
	}[method]
}
