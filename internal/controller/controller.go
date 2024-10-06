package controller

import (
	"net/http"

	_ "github.com/ribeirosaimon/shortify/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/health", NewHealth)
}
