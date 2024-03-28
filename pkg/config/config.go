package config

import "os"

type Config struct {
	DatabaseURL string
}

func Load() *Config {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres password=pass dbname=postgres port=5432 sslmode=disable"
	}

	return &Config{
		DatabaseURL: databaseURL,
	}
}
