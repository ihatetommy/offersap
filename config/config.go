package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser		 	string
	DBName		 	string
	DBPassword 	string
	DBHost		 	string
	DBPort		 	string
	JWTSecretKey string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading from environment variables")
	}

	return &Config{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBName:     getEnv("DB_NAME", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		JWTSecretKey: getEnv("JWT_SECRET_KEY", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
