package main

import (
	"log"

	config "github.com/donnyirianto/go-clean/pkg/config"
	di "github.com/donnyirianto/go-clean/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Gagal Load Config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("Gagal Start Server: ", diErr)
	} else {
		server.Start()
	}
}
