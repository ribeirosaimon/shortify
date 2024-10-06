package controller

import (
	"net/http"
	"time"

	"github.com/ribeirosaimon/shortify/config/env"
	"github.com/ribeirosaimon/shortify/config/response"
)

// NewHealth
// @Summary Get Health
// @Description Check up app
// @Produce  json
// @Router /health [get]
// @Success 200 {HealthDto} HealthDto
func NewHealth(w http.ResponseWriter, r *http.Request) {
	response.Ok(w, HealthDto{
		Up:          true,
		Environment: env.GetConfig().Env,
		Time:        time.Now(),
	})
}

type HealthDto struct {
	Time        time.Time       `json:"time"`
	Environment env.Environment `json:"environment"`
	Up          bool            `json:"up"`
}
