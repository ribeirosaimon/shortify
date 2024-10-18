package main

import (
	"log"
	"os"

	"github.com/ribeirosaimon/shortify/config/adapter"
	"github.com/ribeirosaimon/shortify/internal/controller"
	_ "github.com/ribeirosaimon/shortify/internal/controller"

	"github.com/ribeirosaimon/tooltip/tserver"
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
		tserver.StartEnv(tserver.Environment(myEnv))
	}

	adapter.NewExpress()

	controller.Start()
	config := tserver.GetEnvironment()
	if config.Env == "" {
		log.Fatal("Environment variable not set")
	}
	tserver.NewServer(config)
}
