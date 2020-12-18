package config

type DBConfig struct {
	DbURL      string `yaml:"address" env:"DATABASE_URL" env-default:""`
	DbPort     string `yaml:"port" env:"DBPORT" env-default:"5432"`
	DbHost     string `yaml:"host" env:"DBHOST" env-default:"localhost"`
	DbName     string `yaml:"name" env:"DBNAME" env-default:"postgres"`
	DbUser     string `yaml:"user" env:"DBUSER" env-default:"postgres"`
	DbPassword string `yaml:"password" env:"DBPASSWORD"`
}

type AppConfig struct {
	AppPort       string `yaml:"port" env:"PORT"`
	ReqTimeoutSec int    `yaml:"timeout" env:"REQTIMEOUTSEC" env-default:"10"`
}

type LogConfig struct {
	LogLevel string `yaml:"level" env:"LOGLEVEL" env-default:"INFO"`
}

type Config struct {
	DbConf  *DBConfig  `yaml:"db"`
	AppConf *AppConfig `yaml:"app"`
	LogConf *LogConfig `yaml:"logging"`
}
