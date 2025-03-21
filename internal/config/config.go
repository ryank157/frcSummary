package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)
type Config struct {
    Port             int
    LogLevel         string
    OpenRouterAPIKey string
}
func LoadConfig() (Config, error) {
    // Load .env file
    // First get the current directory
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Error getting current working directory: %v", err)
        return Config{}, err
    }
    // Find the .env file
    envPath := filepath.Join(currentDir, ".env")
    err = godotenv.Load(envPath)
    if err != nil {
        log.Println("Error loading .env file:", err)
        // It is okay to continue without .env file, use production environment variables
    }
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
