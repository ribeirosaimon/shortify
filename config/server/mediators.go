package server

import "github.com/ribeirosaimon/shortify/config/mediator"

func WithUrlPersistMediator(urlPersistMediator mediator.Handler) Option {
	return func(s *services) {
		s.urlPersistMediator = urlPersistMediator
	}
}
