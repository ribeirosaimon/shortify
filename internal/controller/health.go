package controller

import (
	"net/http"
	"time"

	"github.com/ribeirosaimon/shortify/internal/dto"
	"github.com/ribeirosaimon/tooltip/response"
	"github.com/ribeirosaimon/tooltip/tserver"
)

// NewHealth
// @Summary Get Health
// @Description Check up app
// @Produce  json
// @Router /health [get]
// @Success 200 {HealthDto} HealthDto
func NewHealth(w http.ResponseWriter, r *http.Request) {
	response.Ok(w, dto.HealthResponse{
		Up:          true,
		Environment: tserver.GetEnvironment().Env,
		Time:        time.Now(),
	})
}
