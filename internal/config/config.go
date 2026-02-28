package config

import "os"

type Config struct {
	DBURL  string
	APIKey string
}

func Load() *Config {
	return &Config{
		DBURL:  os.Getenv("DATABASE_URL"),
		APIKey: os.Getenv("API_KEY"),
	}
}