package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ribeirosaimon/shortify/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	NewHealth(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var actual dto.Health
	err = json.NewDecoder(rr.Body).Decode(&actual)
	assert.NoError(t, err)

	assert.True(t, actual.Up)
}
