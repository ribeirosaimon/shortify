package promt

import (
	"github.com/prometheus/client_golang/prometheus"
)

var requestCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Número total de requisições HTTP",
})

func Register() {
	prometheus.MustRegister(requestCount)
}
