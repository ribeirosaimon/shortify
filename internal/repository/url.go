package repository

import (
	"sync"

	"github.com/ribeirosaimon/shortify/internal/entity"
)

var urlRepository sync.Once

type urlRepository interface {
	SaveUrl(url *entity.Url) error
}

type urlRepositoryImpl struct {
}

func NewUrl() {
	urlRepositoryImpl := &urlRepositoryImpl{}
}
