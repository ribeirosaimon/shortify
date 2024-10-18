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
// @Success 200 {object} dto.Health "Success"
func NewHealth(w http.ResponseWriter, r *http.Request) {
	response.Ok(w, dto.Health{
		Up:          true,
		Environment: string(tserver.GetEnvironment().Env),
		Time:        time.Now(),
	})
}
