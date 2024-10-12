package vo

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

func TransformStringInId(v string) Id {
	return Id{
		id: v,
	}
}

func (i Id) GetValue() string {
	return i.id
}
