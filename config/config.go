package config

import (
	"os"
)

type Config struct {
	OllamaURL    string
	DefaultModel string
}

func New() *Config {
	return &Config{
		OllamaURL:    getEnvOrDefault("OLLAMA_URL", "http://localhost:11434"),
		DefaultModel: getEnvOrDefault("OLLAMA_MODEL", "llama2"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
