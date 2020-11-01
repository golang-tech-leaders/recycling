package main

import (
	"recycling/internal/platform/database"
	"recycling/internal/platform/server"
)

func main() {
	db := database.NewPostgresWasteStorage("postgres://postgres:postgres@localhost:15432/wastedb?sslmode=disable")
	db.PopulateWasteTypes()
	srv := server.NewServer(":12345", db)
	srv.Run()
}
