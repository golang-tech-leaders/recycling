package main

import (
	"recycling/internal/config"
	"recycling/internal/logger"
	"recycling/internal/platform/database"
	"recycling/internal/platform/server"
)

func main() {

	cfg := config.PrepareConfig()
	logger := logger.New(cfg.LogConf)
	db := database.NewPostgresWasteStorage(cfg.DbConf, logger)
	db.Migrate()
	srv := server.NewServer(cfg.AppConf, db, logger)
	srv.Run()
}
