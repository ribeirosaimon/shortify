package entity

import (
	"github.com/google/uuid"
)

type Id struct {
	id string
}

func NewId() Id {
	return Id{
		id: uuid.NewString(),
	}
}
