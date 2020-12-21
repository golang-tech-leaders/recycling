package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

func PrepareConfig() *Config {
	var cfg Config
	configFile := getConfigFile()

	if err := cleanenv.ReadConfig(configFile, &cfg); err != nil {
		fmt.Printf("Unable to get app configuration due to: %s\n", err.Error())
	}
	return &cfg
}

func getConfigFile() string {
	configFile := flag.String("config", "config.yaml", "config file")
	flag.Parse()
	return *configFile
}
