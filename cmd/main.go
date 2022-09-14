package main

import (
	"log"

	restwebprj "github.com/SyberiaEmperor/rest_web_prj"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/handler"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/repository"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil{
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(restwebprj.Server)

	if err := srv.Run(viper.GetString("port"),handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}