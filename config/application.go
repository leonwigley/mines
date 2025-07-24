package config

import (
	"os"
)

type Config struct {
	AppName     string
	Environment string
	Port        string
}

// GlobalConfig will hold the application-wide configuration
var GlobalConfig Config

// LoadConfig initializes the config from env vars or defaults
func LoadConfig() {
	GlobalConfig = Config{
		AppName:     getEnv("APP_NAME", "Mines"),
		Environment: getEnv("APP_ENV", "development"),
		Port:        getEnv("PORT", "4000"),
	}
}

// helper to read env var or default value
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
