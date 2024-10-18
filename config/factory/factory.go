package factory

import (
	"time"

	"github.com/ribeirosaimon/shortify/config/factory/hash"
	"github.com/ribeirosaimon/shortify/internal/role"
	"github.com/ribeirosaimon/shortify/internal/vo"
)

var now = time.Now()

type EncodeUrlFactory interface {
	Encode() vo.HashedUrl
	GetEncodeType() hash.Method
}

type PlanFactory interface {
	TimeToExpired() time.Duration
}

func EncriptByHashMethod(method hash.Method) EncodeUrlFactory {
	return map[hash.Method]entityBase62{
		hash.Base62: NewBase62(),
	}[method]
}

func GetUserPlan(method role.User) PlanFactory {
	return map[role.User]plan{
		role.BasicUser: newPlan(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())),
	}[method]
}
