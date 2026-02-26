package config

import (
	"os"
)

type Config struct {
	DBPath    string
	JWTSecret string
	Port      string
}

var AppConfig *Config

func Init() {
	AppConfig = &Config{
		DBPath:    getEnv("DB_PATH", "game.db"),
		JWTSecret: getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		Port:      getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
