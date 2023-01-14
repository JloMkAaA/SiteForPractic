package main

import (
	"log"

	"github.com/JloMkAaA/SiteForPractic"
	"github.com/JloMkAaA/SiteForPractic/pkg/handler"
	"github.com/JloMkAaA/SiteForPractic/pkg/repository"
	"github.com/JloMkAaA/SiteForPractic/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(SiteForPractic.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
