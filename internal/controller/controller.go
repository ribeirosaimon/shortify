package controller

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/ribeirosaimon/shortify/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Start() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/health", NewHealth)
	http.HandleFunc("/url-record", NewUrlRecord().NewUrlRecord)

	http.Handle("/metrics", promhttp.Handler())
}
