package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func PrepareConfig() *Config {
	var cfg Config
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
	configFile := flag.String("config", "config.yaml", "config file")
	flag.Parse()
	return *configFile
}
