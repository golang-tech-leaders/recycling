package config

type Config struct {
	AppPort       string `yaml:"app_port" env:"PORT"`
	DbURL         string `yaml:"db_address" env:"DATABASE_URL" env-default:""`
	DbPort        string `yaml:"db_port" env:"DBPORT" env-default:"5432"`
	DbHost        string `yaml:"db_host" env:"DBHOST" env-default:"localhost"`
	DbName        string `yaml:"db_name" env:"DBNAME" env-default:"postgres"`
	DbUser        string `yaml:"db_user" env:"DBUSER" env-default:"postgres"`
	DbPassword    string `yaml:"db_password" env:"DBPASSWORD"`
	ReqTimeoutSec int    `yaml:"req_timeout_sec" env:"REQTIMEOUTSEC" env-default:"10"`
}
