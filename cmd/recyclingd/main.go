package main

import (
	"fmt"
	"recycling/internal/config"
	"recycling/internal/logger"
	"recycling/internal/platform/database"
	"recycling/internal/platform/server"
)

func main() {

	cfg := config.PrepareConfig()
	logger := logger.New(cfg.LogConf)
	fmt.Printf("%s\n", cfg)
	db := database.NewPostgresWasteStorage(cfg.DbConf, logger)
	db.Migrate()
	srv := server.NewServer(cfg.AppConf, db, logger)
	srv.Run()
}
