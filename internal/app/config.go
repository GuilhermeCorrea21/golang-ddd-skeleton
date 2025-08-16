package app

import "os"

var (
	CFG *Config
)

type Config struct {
	DB DBConfig
}

func New() *Config {
	CFG = &Config{
		DB: DBConfig{
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}

	return CFG
}
