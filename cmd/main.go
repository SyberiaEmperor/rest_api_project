package main

import (
	"log"

	restwebprj "github.com/SyberiaEmperor/rest_web_prj"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/handler"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/repository"
	"github.com/SyberiaEmperor/rest_web_prj/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(restwebprj.Server)

	if err := srv.Run("8000",handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}