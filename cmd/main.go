package main

import (
	"log"

	restwebprj "github.com/SyberiaEmperor/rest_web_prj"
)

func main() {
	srv := new(restwebprj.Server)

	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}