package main

import (
	"log"
	"os"

	"github.com/ribeirosaimon/shortify/config/mediator"
	"github.com/ribeirosaimon/shortify/config/server"
	"github.com/ribeirosaimon/shortify/internal/cache"
	"github.com/ribeirosaimon/shortify/internal/controller"
	"github.com/ribeirosaimon/shortify/internal/repository"
	"github.com/ribeirosaimon/shortify/internal/usecase"

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

	urlRepository := repository.NewUrl()
	urlCache := cache.NewUrlRecord()
	urlMediator := mediator.NewPersistUrlMediator(urlRepository, urlCache)
	urlUseCase := usecase.NewUrlRecord(urlMediator)

	server.NewServices(
		server.WithUrlRepository(urlRepository),
		server.WithUrlCache(urlCache),
		server.WithUrlUseCase(urlUseCase),
		server.WithUrlPersistMediator(urlMediator),
	)

	controller.Start()
	config := tserver.GetEnvironment()
	if config.Env == "" {
		log.Fatal("Environment variable not set")
	}
	tserver.NewServer(config)
}
