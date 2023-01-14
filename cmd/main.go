package main

import (
	"log"

	"github.com/JloMkAaA/SiteForPractic"
	"github.com/JloMkAaA/SiteForPractic/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(SiteForPractic.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
