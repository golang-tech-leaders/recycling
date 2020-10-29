package main

import (
	"recycling/internal/platform/database"
	"recycling/internal/platform/server"
)

func main() {
	db := database.NewInmemoryDb()
	srv := server.NewServer(":12345", db)
	srv.Run()
}
