package main

import (
	"fmt"
	"recycling/internal/config"
	"recycling/internal/platform/database"
	"recycling/internal/platform/server"
)

func main() {
	cfg := config.PrepareConfig()
	fmt.Println("DATABASE ENV IS: " + cfg.DbAddress)
	db := database.NewPostgresWasteStorage(cfg)
	db.Migrate()
	srv := server.NewServer(cfg, db)
	srv.Run()
}
