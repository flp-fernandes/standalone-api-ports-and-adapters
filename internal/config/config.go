package config

import "os"

type Config struct {
	PostgresUrl string
}

func Load() Config {
	return Config{
		PostgresUrl: os.Getenv("DATABASE_URL"),
	}
}
