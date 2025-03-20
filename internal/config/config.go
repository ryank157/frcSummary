package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port             int
	LogLevel         string
	OpenRouterAPIKey string
}

func LoadConfig() (Config, error) {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 4000 // Default port if not specified
	}

	config := Config{
		Port:             port,
		LogLevel:         os.Getenv("LOG_LEVEL"),
		OpenRouterAPIKey: os.Getenv("OPENROUTER_API_KEY"),
	}
	return config, nil
}
