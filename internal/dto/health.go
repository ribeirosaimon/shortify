package dto

import (
	"time"
)

type Health struct {
	Time        time.Time `json:"time"`
	Environment string    `json:"environment"`
	Up          bool      `json:"up"`
}
