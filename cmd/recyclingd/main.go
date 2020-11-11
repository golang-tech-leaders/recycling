package main

import (
	"flag"
	"fmt"
	"os"
	models "recycling/internal/model"
	"recycling/internal/platform/database"
	"recycling/internal/platform/server"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	cfg := prepareConfig()
	db := database.NewPostgresWasteStorage(cfg)
	srv := server.NewServer(cfg, db)
	srv.Run()
}

func prepareConfig() *models.Config {
	var cfg models.Config
	configFile := getConfigFile()

	if err := cleanenv.ReadConfig(configFile, &cfg); err != nil {
		fmt.Printf("Unable to get app configuration due to: %s\n", err.Error())
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		fmt.Printf("Unable to retrieve app configuration due to: %s\n", err.Error())
		os.Exit(1)
	}
	return &cfg
}

func getConfigFile() string {
	configFile := flag.String("config", "config.yml", "config file")
	return *configFile
}
