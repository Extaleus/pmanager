package main

import (
	"log"

	"github.com/extaleus/pmanager/pkg/handler"
	server "github.com/extaleus/pmanager/server"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
