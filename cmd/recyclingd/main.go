package main

import (
	"recycling/internal/config"
	"recycling/internal/platform/database"
	"recycling/internal/platform/server"
)

func main() {
	cfg := config.PrepareConfig()
	db := database.NewPostgresWasteStorage(cfg)
	db.Migrate()
	srv := server.NewServer(cfg, db)
	srv.Run()
}
