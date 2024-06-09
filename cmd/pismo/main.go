package main

import (
	"github.com/hameedhub/pismo/internal/config"
	"github.com/hameedhub/pismo/internal/repository"
	"github.com/hameedhub/pismo/internal/repository/database"
	"github.com/hameedhub/pismo/internal/server"
	"github.com/hameedhub/pismo/internal/service"
)

func main() {
	cfg := config.ReadConfig("dev")
	db := database.Run(cfg)
	repo := repository.Init(db)
	init := service.Init(repo)

	server.NewServer(cfg, init).Start()
}
