package config

import (
	"os"
)

var (
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	JWT_SECRET string
)

func Load() {
	DBUser = getEnv("DB_USER", "")
	DBPassword = getEnv("DB_PASSWORD", "")
	DBName = getEnv("DB_NAME", "")
	DBHost = getEnv("DB_HOST", "")
	DBPort = getEnv("DB_PORT", "5432")
	JWT_SECRET = getEnv("JWT_SECRET", "")
}

func GetDSN() string {
	return "user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " host=" + DBHost + " port=" + DBPort + " sslmode=disable"
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
