package main

import (
	_ "github.com/hameedhub/pismo/docs"
	"github.com/hameedhub/pismo/internal/config"
	"github.com/hameedhub/pismo/internal/repository"
	"github.com/hameedhub/pismo/internal/repository/database"
	"github.com/hameedhub/pismo/internal/server"
	"github.com/hameedhub/pismo/internal/service"
	"path/filepath"
)

// @title           Swagger Pismo API
// @version         1.0
// @description     This is a pismo server.

// @host      localhost:8080
// @BasePath  /
func main() {
	path := filepath.Join("..", "..")
	cfg := config.ReadConfig(path)
	db := database.Run(cfg)
	repo := repository.Init(db)
	svc := service.Init(repo)

	server.NewServer(cfg, svc).Start()
}
