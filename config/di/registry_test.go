package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type logService struct {
}
type LogService interface {
	GetValue() string
}

type LogService2 interface {
	GetValue2() string
}

func NewLogService() {
	GetRegistry().Provide("logService", func() any {
		return &logService{}
	})
}

func (l *logService) GetValue() string {
	return "log service called"
}

type secondLogService struct {
}

func NewSecondLogService() {
	GetRegistry().Provide("secondLogService", func() any {
		return &secondLogService{}
	})
}

func (l *secondLogService) GetValue2() string {
	return "log secondLogService called"
}

func TestRegistry(t *testing.T) {
	NewLogService()
	NewSecondLogService()

	t.Run("Have to provide services", func(t *testing.T) {

		myService := registry.Inject("logService").(LogService)
		value := myService.GetValue()
		assert.Equal(t, "log service called", value)

		mySecondService := registry.Inject("secondLogService").(LogService2)
		value2 := mySecondService.GetValue2()
		assert.Equal(t, "log secondLogService called", value2)

	})
}
