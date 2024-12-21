package config

import (
	"os"
)

type Config struct {
	Host string
	Port string
}

func New() *Config {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "8080"
	}

	return &Config{
		Host: host,
		Port: port,
	}
}
