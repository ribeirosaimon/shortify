package controller

import (
	"net/http"

	_ "github.com/ribeirosaimon/shortify/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Start() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/health", NewHealth)
	http.HandleFunc("/url-record", NewUrlRecord().NewUrlRecord)
}
