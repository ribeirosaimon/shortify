package mediator

import "context"

type Handler interface {
	Notify(ctx context.Context, T any) error
}
