package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ribeirosaimon/shortify/config/env"
	_ "github.com/ribeirosaimon/shortify/internal/controller"
)

// @title Shortify
// @swagger: "2.0"
// @description Shortener api
// @termsOfService  http://swagger.io/terms/
// @contact.url    http://www.swagger.io/support
// @contact.email  suporte@swagger.io
// @host      localhost:8080
// @BasePath  /
func main() {
	if myEnv := os.Getenv("ENVIRONMENT"); myEnv != "" {
		env.StartEnv(myEnv)
	} else {
		log.Fatal(errors.New("missing ENVIRONMENT variable"))
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.GetConfig().Port), nil))
}
