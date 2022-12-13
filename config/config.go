package config

import (
	"os"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type authConfig struct {
	HmacSampleSecret string
	Duration         string
}

type AppConfig struct {
	ENV        string
	DBConfig   dbConfig
	AuthConfig authConfig
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var Config = AppConfig{
	ENV: getENV("ENV", "development"),
	DBConfig: dbConfig{
		Host:     getENV("DB_HOST", "localhost"),
		User:     getENV("DB_USER", "giwang"),
		Password: getENV("DB_PASSWORD", "admin"),
		DBName:   getENV("DB_NAME", "house_booking_giwang"),
		Port:     getENV("DB_PORT", "5432"),
	},
}
