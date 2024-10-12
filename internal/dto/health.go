package dto

import (
	"time"

	"github.com/ribeirosaimon/tooltip/tserver"
)

type HealthResponse struct {
	Time        time.Time           `json:"time"`
	Environment tserver.Environment `json:"environment"`
	Up          bool                `json:"up"`
}
