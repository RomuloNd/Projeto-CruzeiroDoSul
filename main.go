package main

import (
	"fmt"
	"log"

	"github.com/RomuloNd/Projeto-CruzeiroDoSul/bootstrap"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/config"
)

func main() {
	app := bootstrap.NewApplication()
	cfg := config.Get()
	addr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)

	if cfg.DevMode {
		log.Fatal(app.Listen(addr))
	} else {
		log.Fatal(app.ListenTLS(addr, cfg.CertFilePath, cfg.KeyFilePath))
	}
}
