package main

import (
	"log"

	"github.com/extaleus/pmanager/pkg/handler"
	"github.com/extaleus/pmanager/pkg/repository"
	"github.com/extaleus/pmanager/pkg/service"
	server "github.com/extaleus/pmanager/server"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
