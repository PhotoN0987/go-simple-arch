package main

import (
	"go-simple-arch/cmd/logger"
	"go-simple-arch/cmd/server"
	"go-simple-arch/pkg/config"
	"go-simple-arch/pkg/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger.LoggingSetting()
	cfg := config.NewConfig()
	db := database.NewDB(cfg)

	server := server.NewServer(cfg, db)
	server.SetUpRouter()
	server.Run()
}
