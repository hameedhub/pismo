package main

import (
	"github.com/hameedhub/pismo/internal/config"
	"github.com/hameedhub/pismo/internal/server"
)

func main() {
	cfg := config.ReadConfig("dev")
	//db := database.Run(cfg)

	server.NewServer(cfg).Start()
}
