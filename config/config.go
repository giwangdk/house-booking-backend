package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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



	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var Config = AppConfig{
	ENV: getENV("ENV", "development"),
	DBConfig: dbConfig{
		Host:     getENV("DB_HOST", ""),
		User:     getENV("DB_USER", ""),
		Password: getENV("DB_PASSWORD", ""),
		DBName:   getENV("DB_NAME", ""),
		Port:     getENV("DB_PORT", ""),
	},
	AuthConfig: authConfig{
		HmacSampleSecret: getENV("HMAC_SAMPLE_SECRET", ""),
		Duration:         getENV("DURATION", ""),
	},
}

func EnvCloudName() string {
	
	return getENV("CLOUDINARY_CLOUD_NAME","")
}

func EnvCloudAPIKey() string {

	return getENV("CLOUDINARY_API_KEY","")
}

func EnvCloudAPISecret() string {

	return getENV("CLOUDINARY_API_SECRET","")
}

func EnvCloudUploadFolder() string {

	return getENV("CLOUDINARY_UPLOAD_FOLDER","")
}